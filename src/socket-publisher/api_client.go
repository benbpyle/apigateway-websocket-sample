package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/apigatewaymanagementapi"
	"os"
)

func NewAPIGatewaySession() *apigatewaymanagementapi.ApiGatewayManagementApi {
	sess, _ := session.NewSession(&aws.Config{
		Region:   aws.String(os.Getenv("REGION")),
		Endpoint: aws.String(os.Getenv("API_ENDPOINT")),
	})

	return apigatewaymanagementapi.New(sess)
}
