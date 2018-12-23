package lambdaapi

import (
	"github.com/aws/aws-lambda-go/events"
)

type CheckAuthFunc func(req *events.APIGatewayProxyRequest) bool

const authHeaderKey = "X-YOUDONT-AUTH"

func RequireRequestWithToken(token string) CheckAuthFunc {
	return func(req *events.APIGatewayProxyRequest) bool {
		var (
			authToken string
			exists    bool
		)

		header := GetHeaderFromRequest(req)
		authToken = header.Get(authHeaderKey)
		exists = authToken != ""
		if !exists {
			authToken, exists = req.QueryStringParameters[authHeaderKey]
		}
		if !exists {
			return false
		}

		return authToken == token
	}
}
