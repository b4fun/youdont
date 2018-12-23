package lambdaapi

import (
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

func GetHeaderFromRequest(req *events.APIGatewayProxyRequest) http.Header {
	h := make(http.Header)
	for k, vs := range req.MultiValueHeaders {
		for _, v := range vs {
			h.Add(k, v)
		}
	}

	return h
}
