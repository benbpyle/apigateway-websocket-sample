# API Gateway WebSocket Sample

Purpose: Working example of using AWS API Gateway, DynamoDB and Lambda to implement a WebSocket

## Getting Started

### Deploying

1. Install [Node.js](https://nodejs.org/en)
2. Install [Golang](https://go.dev/doc/install) 

```bash
# install AWS CDK
npm install -g aws-cdk
# clone the repository
cd apigateway-websocket-sample
npm install
```

Once dependencies have been installed, you are ready to run CDK

```bash
cdk deploy
```

## Destroying

Simply run:

```bash
cdk destroy
```

## Implementation

For a further and in-depth review of how to use this repository and what it supports, head on over the [Blog Article](https://binaryheap.com) <Article not complete but under construction>

## State Machine

![Diagram](./socket_diagram.png)