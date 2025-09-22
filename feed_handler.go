package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/abdol-ahmed/gator/internal/database"
	"github.com/google/uuid"
)

func RSSFeedAggregatorHandler(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %v <time_between_reqs>", cmd.Name)
	}

	timeBetweenRequests, err := time.ParseDuration(cmd.Args[0])
	if err != nil {
		return fmt.Errorf("invalid duration: %w", err)
	}

	log.Printf("Collecting feeds every %s...", timeBetweenRequests)

	ticker := time.NewTicker(timeBetweenRequests)
	for ; ; <-ticker.C {
		scrapeFeeds(s)
	}

	return nil
}

func scrapeFeeds(s *state) {
	feed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		log.Println("Couldn't get next feeds to fetch", err)
	}

	feed, err = s.db.MarkFeedFetched(context.Background(), feed.ID)
	if err != nil {
		log.Printf("Couldn't mark feed %s fetched: %v", feed.Name, err)
		return
	}

	rssFeed, err := fetchFeed(context.Background(), feed.Url)
	if err != nil {
		log.Printf("Couldn't collect feed %s: %v", feed.Name, err)
		return
	}

	printRSSItems(rssFeed.Channel.Item)
	log.Printf("Feed %s collected, %v posts found", feed.Name, len(rssFeed.Channel.Item))
}

func CreateFeedHandler(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 2 {
		return fmt.Errorf("usage: %s <name> <url>", cmd.Name)
	}

	name := cmd.Args[0]
	url := cmd.Args[1]

	feedParam := database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      name,
		Url:       url,
		UserID:    user.ID,
	}

	feed, err := s.db.CreateFeed(context.Background(), feedParam)
	if err != nil {
		return fmt.Errorf("couldn't create feed: %w", err)
	}

	println("feed has been created successfully")

	feedFollowParam := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	}

	_, err = s.db.CreateFeedFollow(context.Background(), feedFollowParam)
	if err != nil {
		return fmt.Errorf("couldn't follow the feed: %w", err)
	}
	println("feed has been followed by '%s'.", user.Name)

	printFeed(feed)
	return nil
}

func GetFeedsHandler(s *state, cmd command) error {
	feedsWithUsers, err := s.db.GetFeedsWithUsers(context.Background())
	if err != nil {
		return fmt.Errorf("couldn't get feeds: %w", err)
	}

	if len(feedsWithUsers) == 0 {
		fmt.Println("No feeds found")
		return nil
	}

	printFeeds(feedsWithUsers)
	return nil
}

func printFeed(feed database.Feed) {
	fmt.Printf(" * ID:      %v\n", feed.ID)
	fmt.Printf(" * Name:    %v\n", feed.Name)
	fmt.Printf(" * URL:     %v\n", feed.Url)
	fmt.Printf(" * UserID:  %v\n", feed.UserID)
}

func printFeeds(feeds []database.GetFeedsWithUsersRow) {
	fmt.Printf("Found %d feeds:\n", len(feeds))
	for _, row := range feeds {
		fmt.Printf(" * Name:    %v\n", row.Feed.Name)
		fmt.Printf(" * URL:     %v\n", row.Feed.Url)
		fmt.Printf(" * User Name:  %v\n", row.User.Name)
	}
}

func printRSSItems(rssItems []RSSItem) {
	for _, rssItem := range rssItems {
		fmt.Printf("Found post: %s\n", rssItem.Title)
	}
}
