package main

import (
	"context"
	"fmt"
	"time"

	"github.com/SuperJake03/gator/internal/database"
	"github.com/google/uuid"
)

func handlerAgg(s *state, cmd command) error {
	feedURL := "https://www.wagslane.dev/index.xml"

	rssFeed, err := fetchFeed(context.Background(), feedURL)
	if err != nil {
		return fmt.Errorf("couldn't fetch feed: %w", err)
	}

	fmt.Printf("%v+\n", *rssFeed)

	return nil
}

func handlerAddFeed(s *state, cmd command) error {
	if len(cmd.Args) != 2 {
		return fmt.Errorf("usage: %s <name> <url>", cmd.Name)
	}

	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("couldn't find user: %w", err)
	}

	feedName := cmd.Args[0]
	feedUrl := cmd.Args[1]

	params := database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      feedName,
		Url:       feedUrl,
		UserID:    user.ID,
	}

	feed, err := s.db.CreateFeed(context.Background(), params)
	if err != nil {
		return fmt.Errorf("couldn't create feed: %w", err)
	}

	fmt.Println("Feed was created!")
	fmt.Printf("Feed created: %v+\n", feed)
	return nil
}

func handlerListFeeds(s *state, cmd command) error {
	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("couldn't list feeds: %w", err)
	}

	for _, feed := range feeds {
		fmt.Printf("Feed Name: %s, URL: %s, Created By: %s\n", feed.FeedName, feed.Url, feed.CreatedBy)
	}

	return nil
}
