package main

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/rcfacil/eventos/pkg/rabbitmq"
)

func main() {
	// abrir um canal de comunicacao com o broker
	ch, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}
	defer ch.Close() // fechar o canal de comunicacao com o broker

	// criar um canal de comunicacao com o consumidor
	msgs := make(chan amqp.Delivery)

	// consumir mensagens da fila
	go rabbitmq.Consume(ch, msgs, "minhafila")

	// loop infinito para consumir mensagens da fila
	for msg := range msgs {
		log.Printf("Message received: %s", msg.Body) // Processa a mensagem, poderia ser um envio de email, um webhook, um agendamento de tarefa, etc
		msg.Ack(false)                               // Envia a mensagem de volta para o broker dizendo que a mensagem foi processada com sucesso
	}
}
