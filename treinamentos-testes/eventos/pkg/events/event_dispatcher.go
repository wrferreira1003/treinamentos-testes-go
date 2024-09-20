package events

import (
	"errors"
	"sync"
)

// Erros
var errHandlerAlreadyRegistered = errors.New("handler already registered")

// EventDispatcher: Estrutura que implementa o EventDispatcherInterface
type EventDispatcher struct {
	handlers map[string][]EventHandlerInterface // Mapa de eventos e seus manipuladores
}

// NewEventDispatcher: Funcao que cria uma nova instancia de EventDispatcher
func NewEventDispatcher() *EventDispatcher {
	return &EventDispatcher{
		handlers: make(map[string][]EventHandlerInterface),
	}
}

// Register: Metodo que registra um manipulador de eventos para um evento específico
func (ed *EventDispatcher) Register(eventName string, handler EventHandlerInterface) error {

	// Verifica se o evento ja existe e se o handler ja esta registrado
	if _, ok := ed.handlers[eventName]; ok {
		for _, h := range ed.handlers[eventName] {
			if h == handler {
				return errHandlerAlreadyRegistered
			}
		}
	}

	// Adiciona o handler ao evento
	ed.handlers[eventName] = append(ed.handlers[eventName], handler)

	return nil
}

// Metodo Clear: Limpa os handlers de um evento específico
func (ed *EventDispatcher) Clear() {
	ed.handlers = make(map[string][]EventHandlerInterface)

}

// Metodo has para verificar se um handler esta registrado para um evento
func (ed *EventDispatcher) Has(eventName string, handler EventHandlerInterface) bool {
	// Verifica se o evento existe
	if _, ok := ed.handlers[eventName]; ok {
		// Verifica se o handler esta registrado para o evento
		for _, h := range ed.handlers[eventName] {
			if h == handler {
				return true
			}
		}
	}
	return false
}

// Dispatch: Metodo que dispara um evento
func (ed *EventDispatcher) Dispatch(event EventInterface) error {
	// Verifica se o evento existe
	if handlers, ok := ed.handlers[event.GetName()]; ok {
		// Dispara o evento para todos os handlers registrados
		wg := sync.WaitGroup{} // Cria uma instancia de WaitGroup
		for _, handler := range handlers {
			wg.Add(1) // Incrementa o contador do WaitGroup
			go handler.Handle(event, &wg)
		}
		wg.Wait() // Aguarda todos os handlers terminarem
	}
	return nil
}

// Remove: Metodo que remove um handler de um evento
func (ed *EventDispatcher) Remove(eventName string, handler EventHandlerInterface) error {
	// Verifica se o evento existe
	if _, ok := ed.handlers[eventName]; ok {
		// Verifica se o handler esta registrado para o evento
		for i, h := range ed.handlers[eventName] {
			if h == handler {
				// Remove o handler do evento
				ed.handlers[eventName] = append(ed.handlers[eventName][:i], ed.handlers[eventName][i+1:]...)
				return nil
			}
		}
	}
	return nil
}
