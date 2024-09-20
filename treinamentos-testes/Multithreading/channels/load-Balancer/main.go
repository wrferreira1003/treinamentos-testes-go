package main

import (
	"fmt"
	"time"
)

// Worker e uma funcao que recebe um workerId e um canal de dados
func Worker(workerId int, data chan int) {
	// forever loop para sempre receber dados do canal
	for x := range data {
		fmt.Printf("Worker %d received %d\n", workerId, x)
		time.Sleep(time.Second)
	}
}

func main() {
	// Criando um canal de dados
	data := make(chan int)
	numberOfWorkers := 10

	// Criando os workers
	for i := 0; i < numberOfWorkers; i++ {
		go Worker(i, data)
	}

	// Enviando dados para o canal
	for i := 0; i < 100; i++ { // Enviando 100 dados para o canal
		data <- i // Enviando dados para o canal
	}

}

// select: e uma estrutura que nos permite esperar uma goroutine terminar.4
