package main

import (
	"context"

	"github.com/abdol-ahmed/gator/internal/database"
)

func LoggedInMiddleware(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {
	return func(s *state, cmd command) error {
		userName := s.config.GetUser()
		user, err := s.db.GetUserByName(context.Background(), userName)
		if err != nil {
			return err
		}

		return handler(s, cmd, user)
	}
}
