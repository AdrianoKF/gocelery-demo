package celery

import (
	"github.com/gocelery/gocelery"
)

func CreateClient() *gocelery.CeleryClient {
	broker := gocelery.NewAMQPCeleryBroker("amqp://guest:guest@localhost")
	backend := gocelery.NewAMQPCeleryBackend("amqp://guest:guest@localhost")

	client, err := gocelery.NewCeleryClient(broker, backend, 2)
	if err != nil {
		panic(err)
	}
	return client
}
