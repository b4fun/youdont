package awsutil

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
)

func MustCreateSession() *session.Session {
	s, err := session.NewSession(&aws.Config{
		Region: &aws.String(os.Getenv("YOUDONT_AWS_REGION")),
	})
	if err != nil {
		panic(fmt.Errorf("create session: %v", err))
	}

	return s
}
