package main

import (
	"context"
	"errors"
	"fmt"
)

func LoginHandler(s *state, cmd command) error {
	if len(cmd.Args) == 0 {
		return errors.New("no arguments")
	}

	name := cmd.Args[0]
	_, err := s.db.GetUserByName(context.Background(), name)

	if err != nil {
		return fmt.Errorf("couldn't find user: %w", err)
	}

	if err := s.config.SetUser(cmd.Args[0]); err != nil {
		return err
	}

	println("user has been set")
	return nil
}
