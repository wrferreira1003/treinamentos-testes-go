package main

import (
	"fmt"
	"net/http"
)

func main() {

	// Criando um handler para a rota raiz
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})

	// Iniciando o servidor na porta 8080
	http.ListenAndServe(":8080", nil)
}
