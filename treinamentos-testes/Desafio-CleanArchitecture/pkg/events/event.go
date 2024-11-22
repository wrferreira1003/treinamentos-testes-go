package events

import "time"

type OrderCreated struct {
	Name     string
	DateTime time.Time
	Payload  interface{}
}

func NewOrderCreatedEvent() EventInterface {
	return &OrderCreated{
		Name: "OrderCreated",
	}
}

func (e *OrderCreated) GetName() string {
	return e.Name
}

func (e *OrderCreated) GetDateTime() time.Time {
	return e.DateTime
}

func (e *OrderCreated) GetPayload() interface{} {
	return e.Payload
}
