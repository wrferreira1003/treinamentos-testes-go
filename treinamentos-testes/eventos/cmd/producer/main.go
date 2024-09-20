package main

import (
	"fmt"

	"github.com/rcfacil/eventos/pkg/rabbitmq"
)

func main() {
	ch, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	//Vamos publicar mensagem na exchange "amq.direct" e ela esta vinculada a fila "minhafila"
	//rabbitmq.Publish(ch, "Hello World", "amq.direct") // envia a mensagem para a fila "minhafila"

	// Defina a quantidade de mensagens a serem enviadas
	quantidadeMensagens := 10

	// Vamos publicar várias mensagens na exchange "amq.direct" e ela está vinculada à fila "minhafila"
	for i := 0; i < quantidadeMensagens; i++ {
		mensagem := fmt.Sprintf("Hello World %d", i+1)
		rabbitmq.Publish(ch, mensagem, "amq.direct") // envia a mensagem para a fila "minhafila"
	}
}
