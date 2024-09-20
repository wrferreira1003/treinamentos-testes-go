package events

import (
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type TestEvent struct {
	Name    string
	Payload interface{}
}

func (e *TestEvent) GetName() string {
	return e.Name
}

func (e *TestEvent) GetPayload() interface{} {
	return e.Payload
}

func (e *TestEvent) GetDateTime() time.Time {
	return time.Now()
}

// TestHandler: Estrutura que implementa o EventHandlerInterface
type TestEventHandler struct {
	ID int
}

// Handle: Metodo que implementa o Handle da interface EventHandlerInterface
func (th *TestEventHandler) Handle(event EventInterface, wg *sync.WaitGroup) {
	fmt.Println("Handler called with event:", event.GetName())
}

// suite de teste ajuda a organizar os testes e facilitar a escrita de testes
type EventDispatcherTestSuite struct {
	suite.Suite     // Heranca da suite de teste
	event           TestEvent
	event2          TestEvent
	handler         TestEventHandler
	handler2        TestEventHandler
	handler3        TestEventHandler
	eventDispatcher *EventDispatcher
}

// SetupTest: Metodo que inicializa o teste
func (suite *EventDispatcherTestSuite) SetupTest() {
	suite.eventDispatcher = NewEventDispatcher()
	suite.handler = TestEventHandler{
		ID: 1,
	}
	suite.handler2 = TestEventHandler{
		ID: 2,
	}
	suite.handler3 = TestEventHandler{
		ID: 3,
	}
	suite.event = TestEvent{Name: "test", Payload: "test"}
	suite.event2 = TestEvent{Name: "test2", Payload: "test2"}
}

func (suite *EventDispatcherTestSuite) TestEventDispatcher_Register() {
	// Testa se o registro do evento e do handler ocorrem com sucesso
	err := suite.eventDispatcher.Register(suite.event.GetName(), &suite.handler) // Registra o evento e o handler
	suite.Nil(err)                                                               // Verifica se o erro e nulo
	suite.Equal(1, len(suite.eventDispatcher.handlers[suite.event.GetName()]))   // Verifica se o handler foi registrado

	err = suite.eventDispatcher.Register(suite.event.GetName(), &suite.handler2) // Tenta registrar o mesmo evento e handler novamente
	suite.Nil(err)
	suite.Equal(2, len(suite.eventDispatcher.handlers[suite.event.GetName()])) // Verifica se o handler foi registrado

	// Verifica se o handler foi registrado no evento
	suite.Equal(&suite.handler, suite.eventDispatcher.handlers[suite.event.GetName()][0])
	suite.Equal(&suite.handler2, suite.eventDispatcher.handlers[suite.event.GetName()][1])

}

// Registra o handler no evento 2
func (suite *EventDispatcherTestSuite) TestEventDispatcher_Register_WithSameHandler() {
	err := suite.eventDispatcher.Register(suite.event.GetName(), &suite.handler)
	suite.Nil(err)
	suite.Equal(1, len(suite.eventDispatcher.handlers[suite.event.GetName()]))

	err = suite.eventDispatcher.Register(suite.event.GetName(), &suite.handler)
	suite.Equal(errHandlerAlreadyRegistered, err)
	suite.Equal(1, len(suite.eventDispatcher.handlers[suite.event.GetName()]))

}

// Testa o Clear
func (suite *EventDispatcherTestSuite) TestEventDispatch_Clear() {
	// colocando o handler no evento 1
	err := suite.eventDispatcher.Register(suite.event.GetName(), &suite.handler)
	suite.Nil(err)
	suite.Equal(1, len(suite.eventDispatcher.handlers[suite.event.GetName()]))

	err = suite.eventDispatcher.Register(suite.event.GetName(), &suite.handler2)
	suite.Nil(err)
	suite.Equal(2, len(suite.eventDispatcher.handlers[suite.event.GetName()]))

	// Colocando o handler no evento 2
	err = suite.eventDispatcher.Register(suite.event2.GetName(), &suite.handler3)
	suite.Nil(err)
	suite.Equal(1, len(suite.eventDispatcher.handlers[suite.event2.GetName()]))

	// limpando o evento 1
	suite.eventDispatcher.Clear()
	suite.Equal(0, len(suite.eventDispatcher.handlers))

}

// SetupTest: Metodo que inicializa o teste
func TestSuite(t *testing.T) {
	suite.Run(t, new(EventDispatcherTestSuite))
}

// Testa o Has
func (suite *EventDispatcherTestSuite) TestEventDispatcher_Has() {
	// colocando o handler no evento 1
	err := suite.eventDispatcher.Register(suite.event.GetName(), &suite.handler)
	suite.Nil(err)
	suite.Equal(1, len(suite.eventDispatcher.handlers[suite.event.GetName()]))

	// colocando o handler no evento 2
	err = suite.eventDispatcher.Register(suite.event.GetName(), &suite.handler2)
	suite.Nil(err)
	suite.Equal(2, len(suite.eventDispatcher.handlers[suite.event.GetName()]))

	// Metodos has
	assert.True(suite.T(), suite.eventDispatcher.Has(suite.event.GetName(), &suite.handler))
	assert.True(suite.T(), suite.eventDispatcher.Has(suite.event.GetName(), &suite.handler2))
	assert.False(suite.T(), suite.eventDispatcher.Has(suite.event.GetName(), &suite.handler3))
}

type MockHandler struct {
	mock.Mock
}

func (m *MockHandler) Handle(event EventInterface, wg *sync.WaitGroup) {
	m.Called(event)
	wg.Done()
}

// Testa o Dispatch
func (suite *EventDispatcherTestSuite) TestEventDispatcher_Dispatch() {
	eh := &MockHandler{}
	eh.On("Handle", &suite.event)

	eh2 := &MockHandler{}
	eh2.On("Handle", &suite.event)

	suite.eventDispatcher.Register(suite.event.GetName(), eh)
	suite.eventDispatcher.Register(suite.event.GetName(), eh2)

	suite.eventDispatcher.Dispatch(&suite.event)
	eh.AssertExpectations(suite.T())  // Verifica se o mockHandler foi chamado
	eh2.AssertExpectations(suite.T()) // Verifica se o mockHandler foi chamado

	eh.AssertNumberOfCalls(suite.T(), "Handle", 1)  // Verifica se o mockHandler foi chamado 1 vez
	eh2.AssertNumberOfCalls(suite.T(), "Handle", 1) // Verifica se o mockHandler foi chamado 1 vez

}

// Remove um handler de um evento
func (suite *EventDispatcherTestSuite) TestEventDispatcher_Remove() {
	// colocando o handler no evento 1
	err := suite.eventDispatcher.Register(suite.event.GetName(), &suite.handler)
	suite.Nil(err)
	suite.Equal(1, len(suite.eventDispatcher.handlers[suite.event.GetName()]))

	// // colocando o handler no evento 2
	err = suite.eventDispatcher.Register(suite.event.GetName(), &suite.handler2)
	suite.Nil(err)
	suite.Equal(2, len(suite.eventDispatcher.handlers[suite.event.GetName()]))

	// // colocando o handler no evento 3
	err = suite.eventDispatcher.Register(suite.event2.GetName(), &suite.handler3)
	suite.Nil(err)
	suite.Equal(1, len(suite.eventDispatcher.handlers[suite.event2.GetName()]))

	// Removendo o handler do evento 1
	suite.eventDispatcher.Remove(suite.event.GetName(), &suite.handler)
	suite.Equal(1, len(suite.eventDispatcher.handlers[suite.event.GetName()]))
	assert.Equal(suite.T(), &suite.handler2, suite.eventDispatcher.handlers[suite.event.GetName()][0])

	// Removendo o handler do evento 2
	suite.eventDispatcher.Remove(suite.event.GetName(), &suite.handler2)
	suite.Equal(0, len(suite.eventDispatcher.handlers[suite.event.GetName()]))

	// Removendo o handler do evento 3
	suite.eventDispatcher.Remove(suite.event2.GetName(), &suite.handler3)
	suite.Equal(0, len(suite.eventDispatcher.handlers[suite.event2.GetName()]))

}
