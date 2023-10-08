import {Construct} from "constructs";
import {Queue} from "aws-cdk-lib/aws-sqs";

export class QueueConstruct extends Construct {
    private readonly _queue: Queue

    get queue(): Queue {
        return this._queue;
    }

    constructor(scope: Construct, id: string) {
        super(scope, id);

        const dlq = new Queue(scope, "PublishDLQ", {
            queueName: "socket-dlq"
        })

        this._queue = new Queue(scope, "PublishQueue", {
           queueName: "socket-queue",
            deadLetterQueue: {
               queue: dlq,
                maxReceiveCount: 1
            }
        });

    }
}