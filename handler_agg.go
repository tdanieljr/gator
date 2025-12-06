package main

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"time"

	"github.com/tdanieljr/gator/internal/database"
)

func handlerAgg(s *state, cmd command) error {
	timeBetweenRequests, err := time.ParseDuration(cmd.args[0])
	if err != nil {
		return err
	}
	ticker := time.NewTicker(timeBetweenRequests)
	for ; ; <-ticker.C {
		err = scrapeFeeds(s)
		if err != nil {
			return err
		}
	}

}
func scrapeFeeds(s *state) error {
	feed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return err
	}
	s.db.MarkFeedFetched(context.Background(), feed.ID)
	rss, err := fetchFeed(context.Background(), feed.Url)
	if err != nil {
		return err
	}
	for _, it := range rss.Channel.Item {
		//fmt.Printf("%v\n", it.Title)
		params := database.CreatePostParams{
			Title: it.Title,
			Url:   it.Link,
			Description: sql.NullString{
				String: it.Description,
				Valid:  true,
			},
			PublishedAt: sql.NullTime{
				Time:  time.Time{},
				Valid: false,
			},
			FeedID: feed.ID,
		}
		err = s.db.CreatePost(context.Background(), params)
		if err != nil {
			continue
		}
	}
	return nil
}
func handlerBrowse(s *state, cmd command) error {
	var limit int32
	if len(cmd.args) == 0 {
		limit = 2

	} else {
		l, err := strconv.Atoi(cmd.args[0])
		if err != nil {
			return err
		}
		limit = int32(l)
	}
	posts, err := s.db.GetPosts(context.Background(), limit)
	if err != nil {
		return err
	}
	for _, p := range posts {
		fmt.Printf("%v\n", p)
	}
	return nil
}
