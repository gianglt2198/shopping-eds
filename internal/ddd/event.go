package ddd

import (
	"time"

	"github.com/google/uuid"
)

type (
	EventPayload interface{}

	Event interface {
		IDer
		EventName() string
		Payload() EventPayload
		Metadata() Metadata
		OccurredAt() time.Time
	}

	event struct {
		Entity
		payload    EventPayload
		metadata   Metadata
		occurredAt time.Time
	}
)

var _ Event = (*event)(nil)

func NewEvent(name string, payload EventPayload, ops ...EventOption) event {
	return newEvent(name, payload, ops...)
}

func newEvent(name string, payload EventPayload, ops ...EventOption) event {
	e := event{
		Entity:     NewEntity(uuid.New().String(), name),
		payload:    payload,
		metadata:   make(Metadata),
		occurredAt: time.Now(),
	}

	for _, op := range ops {
		op.configureEvent(&e)
	}

	return e
}

func (e event) EventName() string     { return e.name }
func (e event) Payload() EventPayload { return e.payload }
func (e event) Metadata() Metadata    { return e.metadata }
func (e event) OccurredAt() time.Time { return e.occurredAt }
