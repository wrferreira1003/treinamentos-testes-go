package main

import (
	"database/sql"
	"log"
	"net"

	"github.com/rcfacil/gRPC-go/internal/database"
	"github.com/rcfacil/gRPC-go/internal/pb"
	"github.com/rcfacil/gRPC-go/internal/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// Cria uma conexão com o banco de dados
	db, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		log.Fatalf("Erro ao abrir o banco de dados: %v", err)
	}
	defer db.Close()

	// Cria uma instância do banco de dados
	categoryDB := database.NewCategory(db)

	// Cria uma instância do serviço de categoria
	categoryService := service.NewCategoryService(categoryDB)

	// Cria um novo servidor gRPC
	grpcServer := grpc.NewServer()

	// Registra o serviço de categoria no servidor
	pb.RegisterCourseCategoryServer(grpcServer, categoryService)

	// Habilita a reflexão para o servidor
	reflection.Register(grpcServer)
	// Cria um listener para o servidor
	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("Erro ao criar o listener: %v", err)
	}

	// Inicia o servidor
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Erro ao iniciar o servidor gRPC: %v", err)
	}
}
