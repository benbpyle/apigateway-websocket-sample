package main

import (
	"context"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/aws"
	awscfg "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/apigatewaymanagementapi"
	"github.com/sirupsen/logrus"
	awstrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/aws/aws-sdk-go-v2/aws"
)

var dbClient *dynamodb.Client
var apigateway *apigatewaymanagementapi.ApiGatewayManagementApi

func main() {
	lambda.Start(handler)
}

func handler(ctx context.Context, event events.SQSEvent) error {
	connections, err := FindConnections(ctx, dbClient)

	if err != nil {
		return err
	}

	logrus.WithFields(logrus.Fields{"event": event}).Debug("The body")
	for _, e := range event.Records {
		b, _ := json.Marshal(e.Body)

		for _, c := range connections {
			connectionInput := &apigatewaymanagementapi.PostToConnectionInput{
				ConnectionId: aws.String(c.ConnectionId),
				Data:         b,
			}

			output, err := apigateway.PostToConnection(connectionInput)

			if err != nil {
				logrus.Errorf("error posting=%s", err)
				return nil
			}

			logrus.Infof("(output)=%v", output)
		}
	}

	return nil
}

func init() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetLevel(logrus.DebugLevel)

	awsCfg, _ := awscfg.LoadDefaultConfig(context.Background())
	awstrace.AppendMiddleware(&awsCfg)
	dbClient = NewDynamoDBClient(awsCfg)
	apigateway = NewAPIGatewaySession()
}
