package events

import (
	"sync"
	"time"
)

// Executa uma acao quando um evento ocorre
type EventHandlerInterface interface {
	Handle(exchange string, event EventInterface, wg *sync.WaitGroup) error // Método para lidar com o evento
}

// Representa um evento
type EventInterface interface {
	GetName() string         // Retorna o nome do evento
	GetDateTime() time.Time  // Retorna a data e hora do evento
	GetPayload() interface{} // Retorna o payload do evento
}

// Gerencia os eventos
type EventDispatcherInterface interface {
	Register(eventName string, handler EventHandlerInterface) error // Registra um handler para um evento
	Dispatch(exchange string, event EventInterface) error           // Dispara o evento
	Remove(eventName string, handler EventHandlerInterface) error   // Remove um handler registrado para um evento
	Has(eventName string, handler EventHandlerInterface) bool       // Verifica se o evento ja possui um handler registrado
	Clear() error                                                   // Limpa todos os handlers registrados para todos os eventos
}
