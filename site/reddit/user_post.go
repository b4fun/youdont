package reddit

// get-user-post?q:subreddit=investing,wallstreet&username=foo&q:title_contain=you

import (
	"fmt"

	graw "github.com/turnage/graw/reddit"
)

// UserPostQuery define how to query a user's posts.
type UserPostQuery struct {
	Username string
	Limit    int

	QueryParams QueryParams

	// Must defines a `AND` group.
	Must []PostQueryClause
}

func (q *UserPostQuery) ListingPath() string {
	return fmt.Sprintf("/u/%s", q.Username)
}

func (q *UserPostQuery) CheckPost(post *graw.Post) bool {
	for _, pred := range q.Must {
		if !pred.CheckPost(q.QueryParams, post) {
			return false
		}
	}
	return true
}

// QueryUserPost queries user latest posts.
func QueryUserPost(q *UserPostQuery) ([]*graw.Post, error) {
	s, err := graw.NewScript(botUserAgent, botScriptRateLimit)
	if err != nil {
		return nil, err
	}

	h, err := s.Listing(q.ListingPath(), "")
	if err != nil {
		return nil, err
	}

	var rv []*graw.Post
	for _, post := range h.Posts {
		if len(rv) >= q.Limit {
			break
		}

		if q.CheckPost(post) {
			rv = append(rv, post)
		}
	}

	return rv, nil
}
