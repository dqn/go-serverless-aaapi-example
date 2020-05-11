package main

import (
	"context"
	"encoding/json"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/dqn/go-serverless-aaapi-example/aaapi"
)

func Handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	crcToken := request.QueryStringParameters["crc_token"]
	consumerSecret := os.Getenv("CONSUMER_SECRET")

	resp := events.APIGatewayProxyResponse{
		StatusCode: http.StatusInternalServerError,
	}

	responseToken, err := aaapi.MakeResponseToken(crcToken, consumerSecret)
	if err != nil {
		return resp, err
	}

	b, err := json.Marshal(map[string]string{
		"response_token": responseToken,
	})
	if err != nil {
		return resp, err
	}

	resp = events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       string(b),
	}

	return resp, nil
}

func main() {
	lambda.Start(Handler)
}
