package lambdaapi

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
)

func AuthRequired() (*events.APIGatewayProxyResponse, error) {
	return &events.APIGatewayProxyResponse{
		StatusCode: 401,
	}, nil
}

func BadRequest() (*events.APIGatewayProxyResponse, error) {
	return &events.APIGatewayProxyResponse{
		StatusCode: 400,
	}, nil
}

func ErrorResponse(err error) (*events.APIGatewayProxyResponse, error) {
	resp, _ := JSONResponse(map[string]string{
		"error": err.Error(),
	})
	resp.StatusCode = 500
	return resp, nil
}

func OKResponse() (*events.APIGatewayProxyResponse, error) {
	return JSONResponse(map[string]bool{
		"ok": true,
	})
}

func JSONResponse(thing interface{}) (*events.APIGatewayProxyResponse, error) {
	b, err := json.Marshal(thing)
	if err != nil {
		return nil, err
	}

	return &events.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body: string(b),
	}, nil
}
