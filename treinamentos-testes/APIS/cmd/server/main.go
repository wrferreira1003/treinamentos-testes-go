package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/wrferreira1003/api-server-go/configs"
	"github.com/wrferreira1003/api-server-go/internal/entity"
	"github.com/wrferreira1003/api-server-go/internal/infra/database"
	"github.com/wrferreira1003/api-server-go/internal/infra/webserver/handlers"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth"
	httpSwagger "github.com/swaggo/http-swagger"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	// Importando o pacote do Swagger
	_ "github.com/wrferreira1003/api-server-go/docs"
)

// @title API Server Go
// @version 1.0
// @description API Server Go com Go e Chi
// @termsOfService http://swagger.io/terms/

// @contact.name Wellington Ferreira
// @contact.url https://github.com/wrferreira1003
// @contact.email wrferreira1003@gmail.com

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8085
// @BasePath /
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	// Estou carregando as configurações do arquivo de configuração
	config, err := configs.LoadConfig(".")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(config)

	// Estou criando a conexão com o banco de dados
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Estou criando as tabelas no banco de dados
	db.AutoMigrate(&entity.Product{}, &entity.User{})
	productDB := database.NewProduct(db)

	productHandler := handlers.NewProductHandler(*productDB)

	// criando um novo handler para usuários
	userDB := database.NewUser(db)
	userHandler := handlers.NewUserHandler(userDB, config.TokenAuth, config.JWTExpiresIn)

	// criando um novo roteador
	r := chi.NewRouter()        // Criando um novo roteador
	r.Use(middleware.Logger)    // Interceptando as requisições e injetando o logger
	r.Use(middleware.Recoverer) // Interceptando as requisições e injetando o recoverer para recuperar erros

	// criando as rotas para o produto
	r.Route("/products", func(r chi.Router) {
		r.Use(jwtauth.Verifier(config.TokenAuth)) // Verificando o token atraves do middleware do chi
		r.Use(jwtauth.Authenticator)              // Autenticando o token com o middleware do chi

		r.Post("/", productHandler.CreateProduct) // Criando uma rota para o produto
		r.Get("/{id}", productHandler.GetProduct)
		r.Put("/{id}", productHandler.UpdateProduct)
		r.Delete("/{id}", productHandler.DeleteProduct)
		r.Get("/", productHandler.GetProducts)
	})

	// criando as rotas para o usuário
	r.Post("/users", userHandler.CreateUser)
	r.Post("/users/generate_token", userHandler.GetJWT)

	// criando a rota para a documentação da API
	r.Get("/docs/*", httpSwagger.Handler(httpSwagger.URL("http://localhost:8085/docs/doc.json")))

	// iniciando o servidor
	http.ListenAndServe(":8085", r) // Iniciando o servidor na porta 8080

}
