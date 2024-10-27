package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {

	db, err := sql.Open("sqlite3", "./db.sqlite")
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}

	// Criando uma instancia do usecase com o wire
	productUsecase := NewUsecase(db)

	// Buscando um produto pelo ID
	product, err := productUsecase.GetProduct(1)
	if err != nil {
		log.Fatalf("Error getting product: %v", err)
	}

	fmt.Println(product.Name)
}
