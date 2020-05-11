package aaapi

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

func MakeResponseToken(crcToken, consumerSecret string) (string, error) {
	mac := hmac.New(sha256.New, []byte(consumerSecret))
	_, err := mac.Write([]byte(crcToken))
	if err != nil {
		return "", err
	}
	responseToken := "sha256=" + base64.StdEncoding.EncodeToString(mac.Sum(nil))
	return responseToken, nil
}

func MakeResponse(crcToken, consumerSecret string) (*events.APIGatewayProxyResponse, error) {
	responseToken, err := MakeResponseToken(crcToken, consumerSecret)
	if err != nil {
		return nil, err
	}

	b, err := json.Marshal(map[string]string{
		"responseToken": responseToken,
	})
	if err != nil {
		return nil, err
	}

	resp := &events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       string(b),
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}

	return resp, nil
}
