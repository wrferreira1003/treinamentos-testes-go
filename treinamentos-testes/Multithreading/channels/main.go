package main

import "fmt"

// Thread 1
func main() {
	canal := make(chan string) // Criando um canal vazio

	// Thread 2
	go func() {
		canal <- "Olá Mundo" // Enviando dados para o canal, canal esta cheio
	}()

	// Thread 1
	msg := <-canal // Recebendo dados do canal, canal esta vazio
	fmt.Println(msg)

}

// Load Balancer: e uma estrutura que nos permite distribuir o trabalho entre multiplas goroutines.
// Producer: e uma goroutine que manda dados para o canal.
// Consumer: e uma goroutine que recebe dados do canal.
// Select: e uma estrutura que nos permite esperar uma goroutine terminar.
