package main

import (
	"fmt"
	"rabbitmq-crash-course/internal/rabbitmq"
)

type App struct {
	Rmq *rabbitmq.RabbitMq
}

func Run() error {
	fmt.Println("Go RabbitMQ Course!")
	rmq := rabbitmq.NewRabbitMqService()
	app := App{
		Rmq: rmq,
	}
	err := app.Rmq.Connect()
	if err != nil {
		return err
	}
	defer app.Rmq.Conn.Close()
	err = app.Rmq.Publish("Code with Bisshwajit")
	if err != nil {
		return err
	}
	app.Rmq.Consume()
	return nil
}

func main() {
	if err := Run(); err != nil {
		fmt.Println("Error Setting up RabbitMQ Application")
		fmt.Println(err)
	}
}
