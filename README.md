# Gocelery example app

Demonstrates simple task invocation using the [https://github.com/gocelery/gocelery](https://github.com/gocelery/gocelery)
library. 

## RabbitMQ

The included [docker-compose.yaml](docker-compose.yaml) spins up a local RabbitMQ server for easy development.

## Running

Run [app/worker.go](app/worker.go) first, then [app/client.go](app/client.go).