package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/rc5091119-pixel/Blog_Aggregator/internal/database"
)

func handlerBrowse(s *state, cmd command, user database.User) error {

	limit := 2

	if len(cmd.Args) > 0 {
		l, err := strconv.Atoi(cmd.Args[0])
		if err == nil {
			limit = l
		}
	}

	posts, err := s.db.GetPostsForUser(
		context.Background(),
		database.GetPostsForUserParams{
			UserID: user.ID,
			Limit:  int32(limit),
		},
	)

	if err != nil {
		return err
	}

	for _, post := range posts {
		fmt.Println(post.Title)
		fmt.Println(post.Url)
		fmt.Println()
	}

	return nil
}
