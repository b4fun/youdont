package main

import (
	"context"
	"net/url"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/b4fun/youdont/lambdaapi"
	"github.com/b4fun/youdont/site/reddit"
	graw "github.com/turnage/graw/reddit"
)

func main() {
	lambda.Start(query)
}

type QueryResponse struct {
	Posts []*graw.Post `json:"posts"`
}

func query(
	ctx context.Context,
	req *events.APIGatewayProxyRequest,
) (*events.APIGatewayProxyResponse, error) {
	q := &reddit.UserPostQuery{
		Limit: 10,
	}
	if username, exists := req.QueryStringParameters["username"]; exists {
		q.Username = username
	} else {
		return lambdaapi.BadRequest()
	}
	q.QueryParams, q.Must = reddit.ParseQueryFromURLValues(url.Values(req.MultiValueQueryStringParameters))

	posts, err := reddit.QueryUserPost(q)
	if err != nil {
		return lambdaapi.ErrorResponse(err)
	}

	return lambdaapi.JSONResponse(QueryResponse{
		Posts: posts,
	})
}
