package main

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"time"
)

type Connection struct {
	PK           string    `dynamodbav:"PK"`
	SK           string    `dynamodbav:"SK"`
	ConnectionId string    `dynamodbav:"ConnectionId"`
	Established  time.Time `dynamodbav:"EstablishedTime"`
}

// NewDynamoDBClient inits a DynamoDB session to be used throughout the services
func NewDynamoDBClient(cfg aws.Config) *dynamodb.Client {
	return dynamodb.NewFromConfig(cfg)
}

func FindConnections(ctx context.Context, client *dynamodb.Client) ([]Connection, error) {
	r, err := client.Scan(ctx, &dynamodb.ScanInput{
		TableName: aws.String("SocketRoster"),
	})

	if err != nil {
		return nil, err
	}
	var connections []Connection

	err = attributevalue.UnmarshalListOfMaps(r.Items, &connections)

	if err != nil {
		return nil, err
	}

	return connections, nil
}
