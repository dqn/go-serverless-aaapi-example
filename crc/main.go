package main

import (
	"context"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/dqn/go-serverless-aaapi-example/aaapi"
)

func Handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	crcToken := request.QueryStringParameters["crc_token"]
	consumerSecret := os.Getenv("CONSUMER_SECRET")

	resp, err := aaapi.MakeResponse(crcToken, consumerSecret)
	if err != nil {
		resp = &events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
		}
		return *resp, err
	}

	return *resp, nil
}

func main() {
	lambda.Start(Handler)
}
