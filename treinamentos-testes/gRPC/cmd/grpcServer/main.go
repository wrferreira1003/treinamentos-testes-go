package main

import (
	"database/sql"
	"log"
	"net"

	_ "github.com/mattn/go-sqlite3"
	"github.com/rcfacil/gRPC-go/internal/database"
	"github.com/rcfacil/gRPC-go/internal/pb"
	"github.com/rcfacil/gRPC-go/internal/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	// conectar no banco de dados
	db, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		log.Fatalf("Erro ao conectar no banco de dados: %v", err)
	}
	defer db.Close()

	// criar o categoryDB
	categoryDB := database.NewCategory(db)

	// criar o categoryService
	categoryService := service.NewCategoryService(*categoryDB)

	// criar o gRPC server
	grpcServer := grpc.NewServer()

	// reflection para o gRPC, serve
	reflection.Register(grpcServer)
	// registrar o categoryService no gRPC server
	pb.RegisterCategoryServiceServer(grpcServer, categoryService)

	// iniciar o servidor
	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("Erro ao iniciar o servidor gRPC: %v", err)
	}

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Erro ao iniciar o servidor gRPC: %v", err)
	}
}
