package webserver

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type WebServer struct {
	Router        chi.Router
	Handlers      map[string]http.HandlerFunc
	WebServerPort string
}

func NewWebServer(webServerPort string) *WebServer {
	return &WebServer{
		Router:        chi.NewRouter(),
		Handlers:      make(map[string]http.HandlerFunc),
		WebServerPort: webServerPort,
	}
}

func (r *WebServer) AddHandler(path string, handler http.HandlerFunc) {
	r.Handlers[path] = handler
	fmt.Println("Adicionado handler para o path:", path)
}

func (r *WebServer) Start() {
	fmt.Println("Iniciando servidor web na porta:", r.WebServerPort)
	r.Router.Use(middleware.Logger)    // Middleware para logar as requisições
	r.Router.Use(middleware.Recoverer) // Middleware para recuperar de panico no servidor

	for path, handler := range r.Handlers {
		r.Router.Handle(path, handler)
	}

	fmt.Println("Servidor web iniciado na porta:", r.WebServerPort)
	http.ListenAndServe(r.WebServerPort, r.Router)
}
