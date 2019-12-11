package amqp

import (
	"github.com/streadway/amqp"
	"os"
)

type IMQ interface {
	OpenMQ()
}

type MQ struct {
	conn *amqp.Connection
}

func (m *MQ) OpenMQ() {

	var err error
	url := os.Getenv("AMQP_URL")

	if url == "" {
		url = "amqp://rabbitmq:rabbitmq@localhost:5672"
	}

	m.conn, err = amqp.Dial(url)

	if err != nil {
		panic("could not establish connection with RabbitMQ:" + err.Error())
	}
}
