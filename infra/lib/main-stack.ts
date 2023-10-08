import { Stack, StackProps } from "aws-cdk-lib";
import { Construct } from "constructs";
import { TableConstruct } from "./constructs/table-construct";
import { SocketConstruct } from "./constructs/socket-construct";
import {QueueConstruct} from "./constructs/queue-construct";

export class MainStack extends Stack {
    constructor(scope: Construct, id: string, props: StackProps) {
        super(scope, id, props);

        const version = new Date().toISOString();
        const table = new TableConstruct(this, "TableConstruct");
        const queue = new QueueConstruct(this, "QueueConstruct");
        const socket = new SocketConstruct(this, "SocketConstruct", table.table, queue.queue);
    }
}
