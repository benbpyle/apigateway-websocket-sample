package main

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

// NewDynamoDBClient inits a DynamoDB session to be used throughout the services
func NewDynamoDBClient(cfg aws.Config) *dynamodb.Client {
	return dynamodb.NewFromConfig(cfg)
}

func DeleteConnection(ctx context.Context, client *dynamodb.Client, connectionId string) error {
	_, err := client.DeleteItem(ctx, &dynamodb.DeleteItemInput{
		TableName: aws.String("SocketRoster"),
		Key: map[string]types.AttributeValue{
			"PK": &types.AttributeValueMemberS{Value: "CONN#" + connectionId},
			"SK": &types.AttributeValueMemberS{Value: "CONN#" + connectionId},
		},
	})

	return err
}
