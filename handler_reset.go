package main

import (
	"context"
	"fmt"
)

func handlerReset(s *state, cmd command) error {
	if cmd.name != "reset" {
		return fmt.Errorf("Trying to reset users without using the reset commond: %v\n", cmd)
	}
	err := s.db.ResetFeeds(context.Background())
	if err != nil {
		return err
	}
	err = s.db.ResetUsers(context.Background())
	if err != nil {
		return err
	}
	return nil
}
