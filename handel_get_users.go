package main

import (
	"context"
	"fmt"
)

func handlerGetUsers(s *state, cmd command) error {
	if cmd.name != "users" {
		return fmt.Errorf("Trying to to get users with wrong command: %v\n", cmd)
	}
	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return err
	}
	for _, u := range users {
		if u.Name == s.cfg.CurrentUserName {
			fmt.Printf("* %s (current)", u.Name)
			continue
		}
		fmt.Printf("* %s", u.Name)
	}

	return nil
}
