package events

import (
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

// Retorna o nome do evento
func (e *TestEvent) GetName() string {
	return e.Name
}

// Retorna o payload do evento
func (e *TestEvent) GetPayload() interface{} {
	return e.Payload
}

// Retorna a data e hora do evento
func (e *TestEvent) GetDateTime() time.Time {
	return time.Now()
}

type TestEventHandler struct {
	ID int
}

func (h *TestEventHandler) Handle(exchange string, event EventInterface, wg *sync.WaitGroup) error {
	return nil
}

// Adicionamos aqui tudo que precisamos para testar o EventDispatcher
type EventDispatcherTestSuite struct {
	suite.Suite
	event           TestEvent
	event2          TestEvent
	handler         TestEventHandler
	handler2        TestEventHandler
	handler3        TestEventHandler
	eventDispatcher *EventDispatcher
}

func (suite *EventDispatcherTestSuite) SetupTest() {
	suite.eventDispatcher = NewEventDispatcher()
	suite.event = TestEvent{Name: "test", Payload: "test"}
	suite.event2 = TestEvent{Name: "test2", Payload: "test2"}
	suite.handler = TestEventHandler{ID: 1}
	suite.handler2 = TestEventHandler{ID: 2}
	suite.handler3 = TestEventHandler{ID: 3}
	suite.event = TestEvent{Name: "test", Payload: "test"}
	suite.event2 = TestEvent{Name: "test3", Payload: "test3"}
}

func (suite *EventDispatcherTestSuite) TestEventDispatcher_Register() {
	// Registra um handler para o evento
	err := suite.eventDispatcher.Register(suite.event.GetName(), &suite.handler)
	suite.Nil(err)                                                                        // Verifica se o erro é nil
	suite.Equal(1, len(suite.eventDispatcher.handlers[suite.event.GetName()]))            // Verifica se o número de handlers registrados é 1
	suite.Equal(&suite.handler, suite.eventDispatcher.handlers[suite.event.GetName()][0]) // Verifica se o handler registrado é o mesmo que o esperado

	// Registra um segundo handler para o mesmo evento
	err = suite.eventDispatcher.Register(suite.event.GetName(), &suite.handler2)
	suite.Nil(err)                                                                         // Verifica se o erro é nil
	suite.Equal(2, len(suite.eventDispatcher.handlers[suite.event.GetName()]))             // Verifica se o número de handlers registrados é 2
	suite.Equal(&suite.handler, suite.eventDispatcher.handlers[suite.event.GetName()][0])  // Verifica se o handler registrado é o mesmo que o esperado
	suite.Equal(&suite.handler2, suite.eventDispatcher.handlers[suite.event.GetName()][1]) // Verifica se o handler registrado é o mesmo que o esperado

	// Verifica se os handlers registrados são os mesmos que os esperados
	assert.Equal(suite.T(), &suite.handler, suite.eventDispatcher.handlers[suite.event.GetName()][0])
	assert.Equal(suite.T(), &suite.handler2, suite.eventDispatcher.handlers[suite.event.GetName()][1])
}

// Testa se o evento pode ter mais de um handler registrado
func (suite *EventDispatcherTestSuite) TestEventDispatcher_WithSameHandler() {
	err := suite.eventDispatcher.Register(suite.event.GetName(), &suite.handler)
	suite.Nil(err)
	suite.Equal(1, len(suite.eventDispatcher.handlers[suite.event.GetName()]))

	//Tentar registra o mesmo handler novamente
	err = suite.eventDispatcher.Register(suite.event.GetName(), &suite.handler)
	suite.Equal(ErrHandlerAlreadyRegistered, err)
	suite.Equal(1, len(suite.eventDispatcher.handlers[suite.event.GetName()]))
}

// Testa se o Clear limpa os handlers registrados
func (suite *EventDispatcherTestSuite) TestEventDispatcher_Clear() {
	// Event one
	err := suite.eventDispatcher.Register(suite.event.GetName(), &suite.handler)
	suite.Nil(err)
	suite.Equal(1, len(suite.eventDispatcher.handlers[suite.event.GetName()]))

	err = suite.eventDispatcher.Register(suite.event.GetName(), &suite.handler2)
	suite.Nil(err)
	suite.Equal(2, len(suite.eventDispatcher.handlers[suite.event.GetName()]))

	// Event two
	err = suite.eventDispatcher.Register(suite.event2.GetName(), &suite.handler3)
	suite.Nil(err)
	suite.Equal(1, len(suite.eventDispatcher.handlers[suite.event2.GetName()]))

	suite.eventDispatcher.Clear()
	suite.Equal(0, len(suite.eventDispatcher.handlers))
}

// Testa se o Has funciona corretamente
func (suite *EventDispatcherTestSuite) TestEventDispatcher_Has() {
	// Event one
	err := suite.eventDispatcher.Register(suite.event.GetName(), &suite.handler)
	suite.Nil(err)
	suite.Equal(1, len(suite.eventDispatcher.handlers[suite.event.GetName()]))

	err = suite.eventDispatcher.Register(suite.event.GetName(), &suite.handler2)
	suite.Nil(err)
	suite.Equal(2, len(suite.eventDispatcher.handlers[suite.event.GetName()]))

	assert.True(suite.T(), suite.eventDispatcher.Has(suite.event.GetName(), &suite.handler))    // Verifica se o handler esta registrado para o evento
	assert.True(suite.T(), suite.eventDispatcher.Has(suite.event.GetName(), &suite.handler2))   // Verifica se o handler esta registrado para o evento
	assert.False(suite.T(), suite.eventDispatcher.Has(suite.event2.GetName(), &suite.handler3)) // Verifica se o handler nao esta registrado para o evento
}

type MockHandler struct {
	mock.Mock
}

// Handle é um metodo do MockHandler que implementa o metodo Handle da interface EventHandlerInterface
func (m *MockHandler) Handle(exchange string, event EventInterface, wg *sync.WaitGroup) error {
	m.Called(exchange, event)
	wg.Done()
	return nil
}

// Testa se o Dispatch funciona corretamente
func (suite *EventDispatcherTestSuite) TestEventDispatcher_Dispatch() {
	eh := &MockHandler{}
	eh.On("Handle", "test", &suite.event)

	eh2 := &MockHandler{}
	eh2.On("Handle", "test", &suite.event)

	// Registra o handler
	suite.eventDispatcher.Register(suite.event.GetName(), eh)
	suite.eventDispatcher.Register(suite.event.GetName(), eh2)

	// Dispara o evento
	suite.eventDispatcher.Dispatch("test", &suite.event)
	eh.AssertExpectations(suite.T())
	eh2.AssertExpectations(suite.T())

	eh.AssertNumberOfCalls(suite.T(), "Handle", 1)
	eh2.AssertNumberOfCalls(suite.T(), "Handle", 1)
}

func (suite *EventDispatcherTestSuite) TestEventDispatcher_Remove() {
	//Event 1
	err := suite.eventDispatcher.Register(suite.event.GetName(), &suite.handler)
	suite.Nil(err)
	suite.Equal(1, len(suite.eventDispatcher.handlers[suite.event.GetName()]))

	err = suite.eventDispatcher.Register(suite.event.GetName(), &suite.handler2)
	suite.Nil(err)
	suite.Equal(2, len(suite.eventDispatcher.handlers[suite.event.GetName()]))

	//Event 2
	err = suite.eventDispatcher.Register(suite.event2.GetName(), &suite.handler3)
	suite.Nil(err)
	suite.Equal(1, len(suite.eventDispatcher.handlers[suite.event2.GetName()]))

	//Remove o handler do evento 1
	err = suite.eventDispatcher.Remove(suite.event.GetName(), &suite.handler)
	suite.Nil(err)
	suite.Equal(1, len(suite.eventDispatcher.handlers[suite.event.GetName()]))
	assert.Equal(suite.T(), &suite.handler2, suite.eventDispatcher.handlers[suite.event.GetName()][0])

	//Remove o handler do evento 2
	err = suite.eventDispatcher.Remove(suite.event.GetName(), &suite.handler2)
	suite.Nil(err)
	suite.Equal(0, len(suite.eventDispatcher.handlers[suite.event.GetName()]))

	//Remove o handler do evento 2
	err = suite.eventDispatcher.Remove(suite.event2.GetName(), &suite.handler3)
	suite.Nil(err)
	suite.Equal(0, len(suite.eventDispatcher.handlers[suite.event2.GetName()]))

}

func TestSuite(t *testing.T) {
	suite.Run(t, new(EventDispatcherTestSuite))
}
