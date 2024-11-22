package handler

import (
	"encoding/json"
	"fmt"
	"sync"

	"github.com/streadway/amqp"
	events "github.com/wrferreira1003/Desafio-Clean-Architecture/pkg/events"
)

type OrderCreatedHandler struct {
	RabbitMQChannel *amqp.Channel
}

func NewOrderCreatedHandler(rabbitMQChannel *amqp.Channel) *OrderCreatedHandler {
	return &OrderCreatedHandler{
		RabbitMQChannel: rabbitMQChannel,
	}
}

func (h *OrderCreatedHandler) Handle(event events.EventInterface, wg *sync.WaitGroup) error {
	defer wg.Done()

	fmt.Printf("Order created: %v", event.GetPayload())
	jsonOutput, err := json.Marshal(event.GetPayload())
	if err != nil {
		return err
	}

	// Cria a mensagem para o RabbitMQ
	messageRabbitmq := amqp.Publishing{
		ContentType: "application/json",
		Body:        jsonOutput,
	}

	err = h.RabbitMQChannel.Publish(
		"amq.direct",
		"",
		false,
		false,
		messageRabbitmq,
	)
	if err != nil {
		return err
	}

	return nil
}
