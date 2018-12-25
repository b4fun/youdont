package main

import (
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/b4fun/youdont/lambdaapi"
	"github.com/b4fun/youdont/site/telegram"
	"github.com/b4fun/youdont/awsutil"
)

var (
	bot       *telegram.Bot
	channelID string
)

func main() {
	bot = telegram.NewBot(
		os.Getenv("YOUDONT_TELEGRAM_BOT_ID"),
		os.Getenv("YOUDONT_TELEGRAM_BOT_TOKEN"),
	)
	channelID = os.Getenv("YOUDONT_TELEGRAM_CHANNEL_ID")

	awsutil.MustCreateSession()

	lambda.Start(run)
}

func sendMessage(text string) error {
	return bot.SendMessage(telegram.ChatMessage{
		ChatID: channelID,
		Text:   text,
	})
}

func run() (*events.APIGatewayProxyResponse, error) {
	if err := sendMessage("hello"); err != nil {
		return lambdaapi.ErrorResponse(err)
	}

	return lambdaapi.OKResponse()
}
