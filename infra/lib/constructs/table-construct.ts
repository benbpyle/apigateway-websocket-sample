import { RemovalPolicy } from "aws-cdk-lib";
import {AttributeType, BillingMode, Table } from "aws-cdk-lib/aws-dynamodb";
import { Construct } from "constructs";

export class TableConstruct extends Construct {
    private readonly _table: Table;

    get table(): Table {
        return this._table;
    }

    constructor(scope: Construct, id: string) {
        super(scope, id);

        this._table = new Table(scope, "SocketTable", {
            billingMode: BillingMode.PAY_PER_REQUEST,
            removalPolicy: RemovalPolicy.DESTROY,
            partitionKey: { name: "PK", type: AttributeType.STRING },
            sortKey: { name: "SK", type: AttributeType.STRING },
            tableName: `SocketRoster`,
        });
    }
}