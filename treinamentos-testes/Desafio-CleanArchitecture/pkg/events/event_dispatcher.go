package events

import (
	"errors"
	"sync"
)

type EventDispatcher struct {
	handlers map[string][]EventHandlerInterface
}

var ErrHandlerAlreadyRegistered = errors.New("handler already registered for this event")

func NewEventDispatcher() *EventDispatcher {
	return &EventDispatcher{
		handlers: make(map[string][]EventHandlerInterface),
	}
}

// Registra um handler para um evento
func (ed *EventDispatcher) Register(eventName string, handler EventHandlerInterface) error {

	// Verifica se o evento ja possui handlers registrados
	if _, ok := ed.handlers[eventName]; ok {
		for _, h := range ed.handlers[eventName] {
			if h == handler {
				return ErrHandlerAlreadyRegistered
			}
		}
	}

	// Adiciona o handler ao evento
	ed.handlers[eventName] = append(ed.handlers[eventName], handler)

	return nil
}

// Limpa os handlers registrados para um evento
func (ed *EventDispatcher) Clear() error {
	ed.handlers = make(map[string][]EventHandlerInterface)
	return nil
}

// Verifica se um handler está registrado para um evento
func (ed *EventDispatcher) Has(eventName string, handler EventHandlerInterface) bool {
	// verifica se o evento existe
	if _, ok := ed.handlers[eventName]; ok {
		// verifica se o handler esta registrado para o evento
		for _, h := range ed.handlers[eventName] {
			if h == handler {
				return true
			}
		}
	}
	return false
}

// Dispara um evento de forma assincrona
func (ed *EventDispatcher) Dispatch(exchange string, event EventInterface) error {
	if handlers, ok := ed.handlers[event.GetName()]; ok {
		wg := &sync.WaitGroup{} // Garante que todos os handlers sejam executados
		for _, handler := range handlers {
			wg.Add(1)                              // Adiciona um goroutine ao WaitGroup
			go handler.Handle(exchange, event, wg) // Executa o handler de forma assincrona
		}
		wg.Wait() // Aguarda todos os goroutines finalizarem
	}
	return nil
}

// Remove um handler registrado para um evento
func (ed *EventDispatcher) Remove(eventName string, handler EventHandlerInterface) error {
	if _, ok := ed.handlers[eventName]; ok {
		for i, h := range ed.handlers[eventName] {
			if h == handler {
				ed.handlers[eventName] = append(ed.handlers[eventName][:i], ed.handlers[eventName][i+1:]...)
				return nil
			}
		}
	}
	return nil
}
