package model

import (
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// Post represents a post.
type Post struct {
	ID        string                 `json:"post_id"`
	SiteID    string                 `json:"site_id"`
	Content   string                 `json:"content"`
	Meta      map[string]interface{} `json:"meta"`
	CreatedAt time.Time              `json:"created_at"`
}

type PostRepository interface {
	QueryLatest(n int) ([]*Post, error)
	Add(*Post) error
}

type postRepository struct {
	db        *dynamodb.DynamoDB
	tableName string
}

func MakePostRepository(db *dynamodb.DynamoDB) *postRepository {
	return &postRepository{
		db:        db,
		tableName: "youdont_post",
	}
}

func (p *postRepository) QueryLatest(n int) ([]*Post, error) {
	q := &dynamodb.ScanInput{
		TableName: aws.String(p.tableName),
	}

	return nil, nil
}

func (p *postRepository) Add(*Post) error {
	return nil
}
