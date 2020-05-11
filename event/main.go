package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/dqn/go-serverless-aaapi-example/aaapi"
)

func Handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	config := oauth1.NewConfig(os.Getenv("CONSUMER_KEY"), os.Getenv("CONSUMER_SECRET"))
	token := oauth1.NewToken(os.Getenv("ACCESS_TOKEN"), os.Getenv("ACCESS_TOKEN_SECRET"))
	httpClient := config.Client(oauth1.NoContext, token)
	client := twitter.NewClient(httpClient)

	resp := events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
	}

	var event aaapi.Event
	err := json.Unmarshal([]byte(request.Body), &event)
	if err != nil {
		fmt.Println(err)
		return resp, err
	}

	for _, v := range event.DirectMessageEvents {
		if v.MessageCreate.SenderID == v.MessageCreate.Target.RecipientID {
			continue
		}

		params := &twitter.DirectMessageEventsNewParams{
			Event: &twitter.DirectMessageEvent{
				Type: "message_create",
				Message: &twitter.DirectMessageEventMessage{
					Target: &twitter.DirectMessageTarget{
						RecipientID: v.MessageCreate.SenderID,
					},
					Data: &twitter.DirectMessageData{
						Text: v.MessageCreate.MessageData.Text,
					},
				},
			},
		}

		_, _, err = client.DirectMessages.EventsNew(params)

		if err != nil {
			fmt.Println(err)
		}
	}

	return resp, nil
}

func main() {
	lambda.Start(Handler)
}
