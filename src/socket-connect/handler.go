package main

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	awscfg "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/sirupsen/logrus"
	awstrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/aws/aws-sdk-go-v2/aws"
)

var dbClient *dynamodb.Client

func main() {
	lambda.Start(handler)
}

func handler(ctx context.Context, event events.APIGatewayWebsocketProxyRequest) (*events.APIGatewayProxyResponse, error) {

	logrus.WithFields(
		logrus.Fields{
			"connectionId":   event.RequestContext.ConnectionID,
			"requestContext": event.RequestContext}).
		Debug("Handling the connection")

	err := WriteConnection(ctx, dbClient, event.RequestContext.ConnectionID)

	if err != nil {
		logrus.WithFields(
			logrus.Fields{"connectionId": event.RequestContext.ConnectionID}).
			Error("Error writing the connection")
		return &events.APIGatewayProxyResponse{
			StatusCode:        500,
			MultiValueHeaders: nil,
			Body:              "{ \"body\": \"bad\" }",
		}, err
	}

	return &events.APIGatewayProxyResponse{
		StatusCode:        200,
		MultiValueHeaders: nil,
		Body:              "{ \"body\": \"good\" }",
	}, nil
}

func init() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetLevel(logrus.DebugLevel)

	awsCfg, _ := awscfg.LoadDefaultConfig(context.Background())
	awstrace.AppendMiddleware(&awsCfg)
	dbClient = NewDynamoDBClient(awsCfg)
}
