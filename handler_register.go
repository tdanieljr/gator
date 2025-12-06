package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/tdanieljr/gator/internal/database"
)

func handlerRegister(s *state, cmd command) error {
	args := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.args[0],
	}
	_, err := s.db.GetUserByName(context.Background(), args.Name)
	if err == nil {
		return fmt.Errorf("User already registered")
	}
	user, err := s.db.CreateUser(context.Background(), args)
	if err != nil {
		return err
	}
	fmt.Printf("User created: %v\n", user)
	s.cfg.SetUser(user.Name)
	return nil
}
