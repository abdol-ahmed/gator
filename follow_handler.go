package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/abdol-ahmed/gator/internal/database"
	"github.com/google/uuid"
)

func CreateFollowHandler(s *state, cmd command, user database.User) error {
	if len(cmd.Args) == 0 {
		return errors.New("no arguments")
	}

	feedUrl := cmd.Args[0]
	feed, err := s.db.GetFeedByURL(context.Background(), feedUrl)
	if err != nil {
		return err
	}

	feedFollowParam := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	}

	feedFollow, err := s.db.CreateFeedFollow(context.Background(), feedFollowParam)
	if err != nil {
		return fmt.Errorf("couldn't follow the feed: %w", err)
	}

	println("feed has been followed by '%s'.", user.Name)
	printFeedFollow(feedFollow)
	return nil
}

func GetFeedFollowOfUserHandler(s *state, cmd command, user database.User) error {
	feedFollows, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("error while fetching feeds: %w", err)
	}
	printFeedFollows(feedFollows)

	return nil
}

func UnfollowFeedHandler(s *state, cmd command, user database.User) error {
	if len(cmd.Args) == 0 {
		return errors.New("no arguments")
	}

	feedUrl := cmd.Args[0]
	feed, err := s.db.GetFeedByURL(context.Background(), feedUrl)
	if err != nil {
		return err
	}

	feedFollowParams := database.DeleteFeedFollowParams{
		UserID: user.ID,
		FeedID: feed.ID,
	}

	err = s.db.DeleteFeedFollow(context.Background(), feedFollowParams)
	if err != nil {
		return fmt.Errorf("error while unfollowing a feed: %w", err)
	}

	return nil
}

func printFeedFollow(feedFollow database.CreateFeedFollowRow) {
	fmt.Printf(" * Feed Name:    %v\n", feedFollow.FeedName)
	fmt.Printf(" * User Name:  	%v\n", feedFollow.UserName)
}

func printFeedFollows(feedFollows []database.GetFeedFollowsForUserRow) {
	for _, feedFollow := range feedFollows {
		fmt.Printf(" * Feed Name:    %v\n", feedFollow.FeedName)
		fmt.Printf(" * User Name:  	%v\n", feedFollow.UserName)
	}
}
