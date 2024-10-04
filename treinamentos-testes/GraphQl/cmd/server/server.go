package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/rcfacil/GraphQl-go/graph"
	"github.com/rcfacil/GraphQl-go/internal/database"

	_ "github.com/mattn/go-sqlite3"
)

const defaultPort = "8090"

func main() {
	// Conectando ao banco de dados
	db, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Instanciando o resolver
	categoryDB := database.NewCategory(db)
	courseDB := database.NewCourse(db)
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	// Criando o servidor GraphQL
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		CategoryDB: categoryDB,
		CourseDB:   courseDB,
	}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
