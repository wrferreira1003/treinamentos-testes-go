package event

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/wrferreira1003/Desafio-Clean-Architecture/internal/infra/event/handler"
	"github.com/wrferreira1003/Desafio-Clean-Architecture/pkg/events"
)

// HandlerRegistry encapsula o registro de todos os handlers.
type HandlerRegistry struct {
	EventDispatcher *events.EventDispatcher
	RabbitMQChannel *amqp.Channel
}

// NewHandlerRegistry cria uma instância do registro de handlers.
func NewHandlerRegistry(dispatcher *events.EventDispatcher, rabbitChannel *amqp.Channel) *HandlerRegistry {
	return &HandlerRegistry{
		EventDispatcher: dispatcher,
		RabbitMQChannel: rabbitChannel,
	}
}

// RegisterHandlers registra todos os handlers.
func (h *HandlerRegistry) RegisterHandlers() error {
	orderCreatedHandler := handler.NewOrderCreatedHandler(h.RabbitMQChannel)
	h.EventDispatcher.Register("OrderCreated", orderCreatedHandler)
	return nil
}
