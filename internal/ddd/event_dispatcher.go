package ddd

import (
	"context"
	"sync"
)

type EventSubscriber interface {
	Subscribe(Event, EventHandler)
}

type EventPublisher interface {
	Publish(context.Context, ...Event) error
}

type EventDispatcher struct {
	handlers map[string][]EventHandler
	mu       sync.Mutex
}

func NewEventDispatcher() *EventDispatcher {
	return &EventDispatcher{
		handlers: make(map[string][]EventHandler),
	}
}

func (h *EventDispatcher) Subscribe(event Event, handler EventHandler) {
	h.mu.Lock()
	defer h.mu.Unlock()

	h.handlers[event.EventName()] = append(h.handlers[event.EventName()], handler)
}

func (h *EventDispatcher) Publish(ctx context.Context, events ...Event) error {
	for _, e := range events {
		for _, handler := range h.handlers[e.EventName()] {
			err := handler(ctx, e)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
