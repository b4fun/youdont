package model

import "time"

// Post represents a post.
type Post struct {
	ID        string                 `json:"post_id"`
	SiteID    string                 `json:"site_id"`
	Content   string                 `json:"content"`
	Meta      map[string]interface{} `json:"meta"`
	CreatedAt time.Time              `json:"created_at"`
}
