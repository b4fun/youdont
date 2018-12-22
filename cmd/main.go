package main

import (
	"fmt"
	"time"

	"github.com/turnage/graw/reddit"
)

func main() {
	s, err := reddit.NewScript("you-dont-have-to", time.Duration(3)*time.Second)
	if err != nil {
		panic(err)
	}

	h, err := s.Listing("/u/ogordained", "")
	if err != nil {
		panic(err)
	}
	for _, post := range h.Posts[:5] {
		fmt.Printf(
			"post: %s %s %s %s\n",
			post.Author, post.Title, post.Subreddit, post.SelfText,
		)
	}

}
