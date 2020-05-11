package aaapi

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
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
