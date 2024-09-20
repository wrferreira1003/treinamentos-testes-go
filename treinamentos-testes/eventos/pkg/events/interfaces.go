package events

import (
	"sync"
	"time"
)

type EventInterface interface {
	GetName() string         // Retorna o nome do evento
	GetDateTime() time.Time  // Retorna a data e hora do evento
	GetPayload() interface{} // Retorna o payload do evento
}

// Toda vez que eu criar um evento, ele tem que implementar essa interface
// Operacoes que serao executadas quando o evento ocorrer.

// EventHandlerInterface: Interface que define o contrato para um manipulador de eventos.
type EventHandlerInterface interface {
	Handle(event EventInterface, wg *sync.WaitGroup) // Metodo Handle que recebe um evento e trata-lo
}

// EventDispatcherInterface: Interface que define o contrato para um gerenciador de eventos.
type EventDispatcherInterface interface {
	Register(eventName string, handler EventHandlerInterface) error // Metodo Register que registra um manipulador de eventos para um evento específico
	Dispatch(event EventInterface) error                            // Metodo Dispatch que dispara um evento
	Remove(eventName string, handler EventHandlerInterface) error   // Metodo Remove que remove um manipulador de eventos para um evento específico
	Has(eventName string, handler EventHandlerInterface) bool       // Metodo Has que verifica se um manipulador de eventos está registrado para um evento específico
	Clear() error                                                   // Metodo Clear que limpa o gerenciador de eventos
}

// Registrando eventos
