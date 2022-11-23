package rabbitmq

import (
	"fmt"
	"github.com/streadway/amqp"
)

type Service interface {
	Connect() error
	Publish(message string) error
	Consume()
}

type RabbitMq struct {
	Conn    *amqp.Connection
	Channel *amqp.Channel
}

func (r *RabbitMq) Connect() error {
	fmt.Println("Connecting to RabbitMQ")
	var err error
	r.Conn, err = amqp.Dial("amqps://btvlhpap:BSx2P8HFocyZWECF4k6EfJvqTt82-sMd@possum.lmq.cloudamqp.com/btvlhpap")
	if err != nil {
		return err
	}
	fmt.Println("Successfully Connected to RabbitMQ Cloud !!")

	r.Channel, err = r.Conn.Channel()
	if err != nil {
		return err
	}
	_, err = r.Channel.QueueDeclare("TestQueue", false, false, false, false, nil)

	return nil
}

func (r *RabbitMq) Publish(message string) error {
	err := r.Channel.Publish("NewKey", "TestQueue", false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(message),
	})

	if err != nil {
		return err
	}
	fmt.Println("Successfully Message Published to Queue !!")
	return nil
}

func (r *RabbitMq) Consume() {
	msgs, err := r.Channel.Consume("TestQueue", "", true, false, false, false, nil)
	if err != nil {
		fmt.Println(err)
	}
	for msg := range msgs {
		fmt.Printf("Message Received: %s\n", msg.Body)
	}
}

func NewRabbitMqService() *RabbitMq {
	return &RabbitMq{}
}
