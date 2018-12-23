package main

import (
	"context"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/b4fun/youdont/lambdaapi"
	"github.com/b4fun/youdont/site/reddit"
	graw "github.com/turnage/graw/reddit"
)

var authFunc lambdaapi.CheckAuthFunc

func main() {
	authFunc = lambdaapi.RequireRequestWithToken(os.Getenv("YOUDONT_API_AUTH_TOKEN"))

	lambda.Start(query)
}

type QueryResponse struct {
	Posts []*graw.Post          `json:"posts"`
	Q     *reddit.UserPostQuery `json:"__q,omitempty"`
	QRaw  map[string][]string   `json:"__q_raw,omitempty"`
}

func query(
	ctx context.Context,
	req *events.APIGatewayProxyRequest,
) (*events.APIGatewayProxyResponse, error) {
	if !authFunc(req) {
		return lambdaapi.AuthRequired()
	}

	q := &reddit.UserPostQuery{
		Limit: 10,
	}
	if username, exists := req.QueryStringParameters["username"]; exists {
		q.Username = username
	} else {
		return lambdaapi.BadRequest()
	}
	q.QueryParams, q.Must = reddit.ParseQueryFromURLValues(req.MultiValueQueryStringParameters)

	posts, err := reddit.QueryUserPost(q)
	if err != nil {
		return lambdaapi.ErrorResponse(err)
	}

	resp := QueryResponse{
		Posts: posts,
	}
	if _, exists := req.QueryStringParameters["debug"]; exists {
		resp.Q = q
		resp.QRaw = req.MultiValueQueryStringParameters
	}

	return lambdaapi.JSONResponse(resp)
}
