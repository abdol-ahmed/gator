package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/abdol-ahmed/gator/internal/database"
	"github.com/google/uuid"
)

func GetPostsHandler(s *state, cmd command, user database.User) error {
	limit := 2
	if len(cmd.Args) == 1 {
		if specifiedLimit, err := strconv.Atoi(cmd.Args[0]); err == nil {
			limit = specifiedLimit
		} else {
			return fmt.Errorf("invalid limit: %w", err)
		}
	}

	postsOfUser, err := s.db.GetPostsOfUser(context.Background(),
		database.GetPostsOfUserParams{
			UserID: user.ID,
			Limit:  int32(limit),
		})
	if err != nil {
		return fmt.Errorf("couldn't get posts: %w", err)
	}

	if len(postsOfUser) == 0 {
		fmt.Println("No posts found")
		return nil
	}

	for _, post := range postsOfUser {
		fmt.Printf("%s from %s\n", post.PublishedAt.Time.Format("Mon Jan 2"), post.FeedName)
		fmt.Printf("--- %v ---\n", post.Title)
		fmt.Printf("    %v\n", post.Description.String)
		fmt.Printf("Link: %s\n", post.Url)
		fmt.Println("=====================================")
	}
	return nil
}

func savePosts(s *state, feedId uuid.UUID, rssItems []RSSItem) {
	for _, rssItem := range rssItems {
		_, err := s.db.CreatePost(context.Background(), database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now().UTC(),
			UpdatedAt:   time.Now().UTC(),
			Title:       sql.NullString{String: rssItem.Title, Valid: true},
			Url:         rssItem.Link,
			Description: sql.NullString{String: rssItem.Description, Valid: true},
			PublishedAt: stringToNullTime(rssItem.PubDate),
			FeedID:      feedId,
		})
		if err != nil && !strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			log.Printf("Couldn't create post: %v", err)
		}
	}
}

func stringToNullTime(s string) sql.NullTime {
	if s == "" {
		return sql.NullTime{Valid: false}
	}

	t, err := time.Parse(time.RFC1123Z, s) // Or use a different layout based on your string format
	if err != nil {
		fmt.Printf("Warning: Failed to parse time string '%s': %v\n", s, err)
		return sql.NullTime{Valid: false}
	}

	return sql.NullTime{Time: t, Valid: true}
}
