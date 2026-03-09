package main

import (
	"context"
	"fmt"

	"github.com/rc5091119-pixel/Blog_Aggregator/internal/database"
)

func handlerUnfollow(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("could not unfollow")
	}
	url := cmd.Args[0]
	feed, err := s.db.GetFeedByURL(context.Background(), url)
	if err != nil {
		return err
	}
	errr := s.db.DeleteFeedFollows(context.Background(), database.DeleteFeedFollowsParams{
		UserID: user.ID,
		FeedID: feed.ID,
	})
	if errr != nil {
		return fmt.Errorf("could not unfollow %w", errr)
	}
	fmt.Printf("User %s has unfollowed feed %s\n", user.Name, feed.Name)
	return nil
}
