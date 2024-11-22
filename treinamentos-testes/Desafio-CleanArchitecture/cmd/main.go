package main

import (
	"fmt"
	"net"

	"github.com/wrferreira1003/Desafio-Clean-Architecture/cmd/graph"
	"github.com/wrferreira1003/Desafio-Clean-Architecture/cmd/webserver"
	"github.com/wrferreira1003/Desafio-Clean-Architecture/cmd/wire"
	"github.com/wrferreira1003/Desafio-Clean-Architecture/config"
	"github.com/wrferreira1003/Desafio-Clean-Architecture/internal/infra/event"
	"github.com/wrferreira1003/Desafio-Clean-Architecture/internal/infra/grpc/pb"
	"github.com/wrferreira1003/Desafio-Clean-Architecture/internal/infra/grpc/services"
	"github.com/wrferreira1003/Desafio-Clean-Architecture/internal/infra/web"
	"github.com/wrferreira1003/Desafio-Clean-Architecture/pkg/db/postgress"
	"github.com/wrferreira1003/Desafio-Clean-Architecture/pkg/events"
	"github.com/wrferreira1003/Desafio-Clean-Architecture/pkg/rabbitmq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	//Carregar configurações
	configPath := "./cmd"
	config, err := config.LoadConfig(configPath)
	if err != nil {
		panic(err)
	}
	fmt.Println("Configurações carregadas")
	fmt.Println(config)

	//Conectar com o banco de dados
	db, err := postgress.NewDBConnection(&postgress.PostgresDB{Config: config})
	if err != nil {
		panic(err)
	}
	fmt.Println("Conectado ao banco de dados")
	defer db.Close()

	//Conectar com o RabbitMQ
	channel, err := rabbitmq.OpenChannelConnection("clean_architecture", "input", &rabbitmq.RabbitMQ{Config: config})
	if err != nil {
		panic(err)
	}
	defer channel.Close()
	fmt.Println("Conectado ao RabbitMQ")

	// Inicializar e registrar os handlers
	eventDispatcher := events.NewEventDispatcher()
	handlerRegistry := event.NewHandlerRegistry(eventDispatcher, channel)
	handlerRegistry.RegisterHandlers() // Registra todos os handlers

	//Inicializar os arquivos do wire
	createOrderUseCase, err := wire.InitializeOrderUseCase(db)
	if err != nil {
		panic(err)
	}

	//Inicializar o handler do web
	webOrderHandler := web.NewWebOrderHandler(createOrderUseCase)

	//Inicializar o servidor web
	webserver := webserver.NewWebServer(config.WebServerPort)
	webserver.AddHandler("/order", webOrderHandler.Create)
	fmt.Println("Servidor web iniciado na porta", config.WebServerPort)
	go webserver.Start() // Iniciar o servidor web em uma goroutine

	//Inicializar o servidor grpc
	grpcServer := grpc.NewServer()
	createOrderService := services.NewOrderService(createOrderUseCase)
	pb.RegisterOrderServiceServer(grpcServer, createOrderService)
	reflection.Register(grpcServer)

	fmt.Println("Servidor gRPC iniciado na porta", config.GRPCPort)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", config.GRPCPort))
	if err != nil {
		panic(err)
	}
	go grpcServer.Serve(lis)

	//Servidor GraphQL.
	graph.NewServer(createOrderUseCase)

}
