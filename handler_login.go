package main

import (
	"context"
	"fmt"

	"github.com/tdanieljr/gator/internal/database"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("Expected 1 argument")
	}

	user, err := s.db.GetUserByName(context.Background(), cmd.args[0])
	if err != nil {
		return err
	}
	fmt.Printf("User has been set to: %s\n", user.Name)
	err = s.cfg.SetUser(user.Name)
	if err != nil {
		return err
	}

	return nil

}
func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {
	return func(s *state, cmd command) error {
		user, err := s.db.GetUserByName(context.Background(), s.cfg.CurrentUserName)
		if err != nil {
			return err
		}
		return handler(s, cmd, user)

	}
}
