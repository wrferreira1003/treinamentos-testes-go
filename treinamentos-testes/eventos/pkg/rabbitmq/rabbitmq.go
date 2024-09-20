package rabbitmq

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

func OpenChannel() (*amqp.Channel, error) {
	conn, err := amqp.Dial("amqp://admin:admin@localhost:5672/")
	if err != nil {
		log.Println("Error connecting to RabbitMQ:", err)
		panic(err)

	}
	// abrir um canal de comunicacao com o broker
	ch, err := conn.Channel()
	if err != nil {
		log.Println("Error opening channel:", err)
		panic(err)
	}

	return ch, nil
}

// Consumer: Metodo que consome mensagens de uma fila
func Consume(ch *amqp.Channel, out chan<- amqp.Delivery, queue string) error {
	msgs, err := ch.Consume(
		queue,         //nome da fila no exchange
		"go-consumer", //nome da tag do consumidor
		false,         //autoAck - se o consumidor deve ou nao enviar a mensagem de volta para o broker
		false,         //exclusive - se o consumidor tem permissao para consumir a mensagem de forma exclusiva
		false,         //noLocal - se o consumidor nao deve consumir a mensagem se ele mesmo ja a consumiu
		false,         //noWait - se o consumidor nao deve esperar pela confirmacao da mensagem
		nil,           //args
	)
	if err != nil {
		log.Println("Error consuming message:", err)
		return err
	}

	for msg := range msgs {
		out <- msg
	}

	return nil
}

// Producer: Metodo que envia mensagens para uma fila
func Publish(ch *amqp.Channel, body string, exchange string) error {
	err := ch.Publish(
		exchange, //exchange, se for vazio, o rabbitmq vai usar o default exchange, qunado bate nessa exchange, ele vai rotear para a fila conforme as configuracoes da fila
		"",       //routing key, se for vazio, o rabbitmq vai usar o default exchange
		false,    //mandatory, se a mensagem nao for roteada para nenhuma fila, ela sera descartada
		false,    //immediate
		amqp.Publishing{
			//DeliveryMode: amqp.Persistent, //mensagem persistente, sera salva no banco de dados do rabbitmq
			ContentType: "text/plain", //tipo de conteudo da mensagem
			Body:        []byte(body), //corpo da mensagem
		},
	)
	if err != nil {
		log.Println("Error publishing message:", err)
		return err
	}
	return nil
}
