# Example Microservice

[![Staging](https://github.com/kyani-inc/kms-example/workflows/Staging/badge.svg)](https://github.com/kyani-inc/kms-example/actions?query=workflow%3AStaging)
[![Production](https://github.com/kyani-inc/kms-example/workflows/Production/badge.svg)](https://github.com/kyani-inc/kms-example/actions?query=workflow%3AProduction)

This is an example second-generation [KMS](https://github.com/kyani-inc/kms/tree/v2.x/v2) microservice running in both staging and production. It showcases all available KMSv2 service layers (gRPC, REST, and Worker).

## Creating New Services

To create a new service, click the `Use this Template` button to copy this skeleton to a new repository. From there, clone the new repository and run `./setup.sh` to setup the basics for your new service:

![setup demo](./.github/images/setup.png)

**⚠️ NOTE:** You must have `AWS_ACCESS_KEY_ID` and `AWS_SECRET_ACCESS_KEY` variables in your environment as well as appropriate AWS permissions to run this script correctly.

--------------------------

## Service Diagram

![kms-example Architecture](./.github/images/architecture.png)

## Getting Started

1. Copy `env-sample` to `env` and add your credentials
2. Ensure you have installed `protoc` (see [here](https://github.com/kyani-inc/proto#compiling-protocol-buffer-definitions)).
3. Ensure you are connected to the Kyäni VPC

Run:

```bash
➜ make proto
➜ make api
```

## gRPC Service

The example microservice uses the following service definition: [service.proto](https://github.com/kyani-inc/proto/blob/master/example/service.proto)


## REST Endpoints


### `GET /hello/:name`

Simple example endpoint that takes a `name` and says hello :)

> 200 OK
```json
{
  "reply": "Hello Ross, nice to meet you!"
}
```

This handler publishes a `kyani.example.EventPersonSaidHello` event with the name of the person and the reply.

## Events Subscriptions

Definitions for the events that are published by `kms-example` can be found in the [Kyäni Proto](https://github.com/kyani-inc/proto/blob/master/example/events.proto) repository. `kms-example` also _subscribes_ to the following events:

### `kyani.example.EventPersonSaidHello`

Example event for saying hello and demonstrating KMS Events functionality
