package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/abdol-ahmed/gator/internal/database"
	"github.com/google/uuid"
)

func CreateUserHandler(s *state, cmd command) error {
	if len(cmd.Args) == 0 {
		return errors.New("no arguments")
	}

	name := cmd.Args[0]
	userParam := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      name,
	}

	user, err := s.db.CreateUser(context.Background(), userParam)
	if err != nil {
		return fmt.Errorf("couldn't create user: %w", err)
	}

	if err := s.config.SetUser(user.Name); err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}

	println("user has been created successfully")
	printUser(user)
	return nil
}

func DeleteUsersHandler(s *state, cmd command) error {
	err := s.db.DeleteAllUsers(context.Background())
	if err != nil {
		return fmt.Errorf("error while deleting all users: %w", err)
	}

	println("users have been deleted successfully")
	return nil
}

func GetUsersHandler(s *state, cmd command) error {
	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("couldn't fetch all users: %w", err)
	}
	loggedInUser := s.config.GetUser()
	printUsers(users, loggedInUser)
	return nil
}

func printUser(user database.User) {
	fmt.Printf(" * ID:      %v\n", user.ID)
	fmt.Printf(" * Name:    %v\n", user.Name)
}

func printUsers(users []database.User, loggedInUser string) {
	for _, user := range users {
		fmt.Printf(" * %s", user.Name)
		if user.Name == loggedInUser {
			fmt.Printf(" (current)\n")
		} else {
			fmt.Printf("\n")
		}
	}
}
