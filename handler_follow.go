package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/rc5091119-pixel/Blog_Aggregator/internal/database"
)

func handlerFollow(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("we have not any url so we can't follow")
	}
	url := cmd.Args[0]
	feed, err := s.db.GetFeedByURL(context.Background(), url)
	if err != nil {
		return fmt.Errorf("could not feed %w", err)
	}

	feed_follow, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	})
	if err != nil {
		return fmt.Errorf("could not follow feeds %w", err)
	}
	fmt.Printf("Feed: %s\nFollowed by: %s\n", feed_follow.FeedName, feed_follow.UserName)
	return nil
}
func handlerFollowing(s *state, cmd command, user database.User) error {
	following, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("could not found any feed for current user %w", err)
	}
	for _, feed := range following {
		fmt.Println(feed.FeedName)
	}
	return nil
}
