import {Construct} from "constructs";
import {WebSocketApi, WebSocketStage} from "@aws-cdk/aws-apigatewayv2-alpha";
import {GoFunction} from "@aws-cdk/aws-lambda-go-alpha";
import {Duration} from "aws-cdk-lib";
import { WebSocketLambdaIntegration } from '@aws-cdk/aws-apigatewayv2-integrations-alpha';
import {Table} from "aws-cdk-lib/aws-dynamodb";
import {Queue} from "aws-cdk-lib/aws-sqs";
import {SqsEventSource} from "aws-cdk-lib/aws-lambda-event-sources";


export class SocketConstruct extends Construct {
    private readonly _api: WebSocketApi;

    constructor(scope: Construct, id: string, table: Table, queue: Queue) {
        super(scope, id);

        const f = new GoFunction(scope, "SocketConnectFunction", {
            entry: "src/socket-connect",
            functionName: `socket-connect`,
            timeout: Duration.seconds(15),
            environment: {
                IS_LOCAL: "false",
                LOG_LEVEL: "DEBUG",
            },
        });


        const f2 = new GoFunction(scope, "SocketDisConnectFunction", {
            entry: "src/socket-disconnect",
            functionName: `socket-disconnect`,
            timeout: Duration.seconds(15),
            environment: {
                IS_LOCAL: "false",
                LOG_LEVEL: "DEBUG",
            },
        });

        const f3 = new GoFunction(scope, "SocketPublisher", {
            entry: "src/socket-publisher",
            functionName: `socket-stream-publisher`,
            timeout: Duration.seconds(15),
            environment: {
                IS_LOCAL: "false",
                LOG_LEVEL: "DEBUG",
                API_ENDPOINT: "<insert endpoint>",
                REGION: "<insert region>"
            },
        });

        table.grantReadWriteData(f);
        table.grantReadWriteData(f2);
        table.grantReadWriteData(f3);

        f3.addEventSource(new SqsEventSource(queue, {
            batchSize: 10
        }));

        this._api = new WebSocketApi(this, "RestApi", {
            description: "Sockets API",
            apiName: "sockets-api",
            connectRouteOptions: {
                integration: new WebSocketLambdaIntegration('ConnectIntegration', f)
            },
            disconnectRouteOptions: {
                integration: new WebSocketLambdaIntegration('DisConnectIntegration', f2)
            },
        });

        this._api.grantManageConnections(f3);
        new WebSocketStage(this, 'SocketStage', {
            webSocketApi: this._api,
            stageName: 'main',
            autoDeploy: true,
        });
    }


    get api(): WebSocketApi {
        return this._api;
    }
}
