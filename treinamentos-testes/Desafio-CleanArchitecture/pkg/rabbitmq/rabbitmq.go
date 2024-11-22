package rabbitmq

import (
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/wrferreira1003/Desafio-Clean-Architecture/config"
)

type RabbitMQ struct {
	Config *config.Config
}

// Abre uma conexão com o RabbitMQ e retorna um canal de comunicação
func OpenChannelConnection(exchange string, queue string, r *RabbitMQ) (*amqp.Channel, error) {
	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%s/", r.Config.RabbitMQUser, r.Config.RabbitMQPassword, r.Config.RabbitMQHost, r.Config.RabbitMQPort))
	if err != nil {
		panic(err)
	}

	// Cria um canal de comunicação com o RabbitMQ
	channel, err := conn.Channel()
	if err != nil {
		panic(err)
	}

	// Declara uma exchange personalizada
	err = channel.ExchangeDeclare(
		exchange, // Nome da exchange
		"direct", // Tipo de exchange
		true,     // Durable
		false,    // Auto-delete
		false,    // Internal
		false,    // NoWait
		nil,      // Arguments
	)
	if err != nil {
		return nil, err
	}

	// Declara a fila "input"
	_, err = channel.QueueDeclare(
		queue, // Nome da fila
		true,  // Durable
		false, // Delete when unused
		false, // Exclusive
		false, // NoWait
		nil,   // Arguments
	)
	if err != nil {
		return nil, err
	}

	// Vincula a fila à exchange personalizada
	err = channel.QueueBind(
		queue,         // Nome da fila
		"routing_key", // Routing key
		exchange,      // Nome da exchange
		false,         // NoWait
		nil,           // Arguments
	)
	if err != nil {
		return nil, err
	}

	return channel, nil
}

// Consome uma mensagem do RabbitMQ
func Consume(channel *amqp.Channel, output chan amqp.Delivery, queue string) error {
	msgs, err := channel.Consume(
		queue,         // Nome da fila
		"go-consumer", // Nome do consumidor
		false,         // AutoAck
		false,         // Exclusive
		false,         // NoLocal
		false,         // NoWait
		nil,           // Arguments
	)
	if err != nil {
		return err
	}

	go func() {
		for msg := range msgs {
			output <- msg
		}
	}()

	return nil
}

// Publica uma mensagem na exchange personalizada
func Publish(channel *amqp.Channel, body string, exchange string) error {
	err := channel.Publish(
		exchange,      // Nome da exchange
		"routing_key", // Routing key
		false,         // Mandatory
		false,         // Immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	if err != nil {
		return err
	}
	return nil
}
