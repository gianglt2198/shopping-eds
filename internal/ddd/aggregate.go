package ddd

type Aggregate interface {
	Entity
	AddEvents(Event)
	GetEvents() []Event
}

type AggregateBase struct {
	ID     string
	events []Event
}

func (a AggregateBase) GetID() string {
	return a.ID
}

func (a *AggregateBase) AddEvents(event Event) {
	a.events = append(a.events, event)
}

func (a *AggregateBase) GetEvents() []Event {
	return a.events
}
