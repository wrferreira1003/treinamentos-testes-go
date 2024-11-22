package main

import (
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/wrferreira1003/Desafio-Clean-Architecture/cmd/graph"
	"github.com/wrferreira1003/Desafio-Clean-Architecture/pkg/rabbitmq"
)

func main() {
	channel, err := rabbitmq.OpenChannelConnection("clean_architecture", "input")
	if err != nil {
		panic(err)
	}
	defer channel.Close()

	msgs := make(chan amqp.Delivery)

	// Consumir mensagens da fila input
	go rabbitmq.Consume(channel, msgs, "input")
	for msg := range msgs {
		fmt.Println(string(msg.Body))
		msg.Ack(false) // Confirma a recepção da mensagem
	}

	//Servidor GraphQL.
	graph.NewServer()

}
