package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(query)
}

func query() (*events.APIGatewayProxyResponse, error) {
	resp := &events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       "hello",
	}
	return resp, nil
}
