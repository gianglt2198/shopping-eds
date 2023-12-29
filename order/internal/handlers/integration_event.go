package handlers

import (
	"shopping/internal/ddd"
	"shopping/order/internal/domain"
)

func RegisterIntegrationEventHandlers(eventHandlers ddd.EventHandler[ddd.AggregateEvent], domainSubscriber ddd.EventSubscriber[ddd.AggregateEvent]) {
	domainSubscriber.Subscribe(eventHandlers,
		domain.OrderCreatedEvent,
		domain.OrderAddedItemEvent,
		domain.OrderCheckedOutEvent,
		domain.OrderReadiedEvent,
		domain.OrderCanceledEvent,
		domain.OrderCompletedEvent,
	)
}
