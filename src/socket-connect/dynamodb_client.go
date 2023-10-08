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

func WriteConnection(ctx context.Context, client *dynamodb.Client, connectionId string) error {

	c := &Connection{
		PK:           "CONN#" + connectionId,
		SK:           "CONN#" + connectionId,
		ConnectionId: connectionId,
		Established:  time.Now(),
	}
	m, _ := attributevalue.MarshalMap(c)

	_, err := client.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: aws.String("SocketRoster"),
		Item:      m,
	})

	return err
}
