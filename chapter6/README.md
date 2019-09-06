# Chapter 6 - OSS @ Beat

## [Patron](https://github.com/beatlabs/patron)

Patron is a microservice framework.

### Setting up

There is a docker compose file present in order to spin up everything needed.

```bash
docker-compose up -d
```

### First example

The first example is the entry point of the demonstration:

- Receives a HTTP POST with a JSON payload
- Sends the same payload in protobuf format to the second service

### Second example

- Receives a HTTP POST with a protobuf payload
- Sends the same payload in JSON format to a Kafka topic

### Third example

- Receives a JSON message from Kafka
- Sends the same payload in JSON format to a RabbitMQ

### Fourth example

- Receives a JSON message from RabbitMQ
- Sends the same payload in JSON format to AWS SNS which is connected to a SQS

### Fifth example

- Receives a JSON message from AWS SQS
- Logs the payload to console

### Try out

Send a request to the first service:

```bash
curl -d '{"Firstname":"John", "Lastname": "Doe"}' -H "Content-Type: application/json" -X POST http://localhost:50000
```

Watch the logging of each service and after that head over to [jaeger](http://localhost:16686/search) and [prometheus](http://localhost:9090/graph).

### CLI

The framework supplies a cli in order to simplify repository generation with the following features:

- git repository creation
- cmd folder and main.go creation with build version support

```bash
go build -ldflags '-X main.version=1.0.0' main.go)
```

- go module support and vendoring
- Dockerfile with version support (docker build --build-arg version=1.0.0)
- The latest version can be installed with

```bash
go get github.com/beatlabs/patron/cmd/patron
```

The below is an example of a service created with the cli that has a module name github.com/beatlabs/test and will be created in the test folder in the current directory.

```bash
patron -m "github.com/beatlabs/test" -p "test"
```

## [Harvester](https://github.com/beatlabs/harvester)

`Harvester` is a configuration library which helps setting up and monitoring configuration values in order to dynamically reconfigure your application.

`Harvester` is part of a bigger ecosystem which consists of:

- A backend for handling the configurations
- A portal for giving access to user to change configuration
- Consul, as a distributed KV store
- A service that monitors Kubernetes for new deployments in order to update the available configuration from the newly deployed service.

Any service that wants to participate in the above ecosystem has to provide a endpoint which exposes the available configuration options.

`Harvester` is used in the final services to help managing and updating configuration.

Configuration can be obtained from the following sources:

- Seed values, are hard-coded values into your configuration struct
- Environment values, are obtained from the environment
- Flag values, are obtained from CLI flags with the form -flag=value
- Consul, which is used to get initial values and to monitor them for changes
