package handlers

import (
	"shopping/internal/ddd"
	"shopping/order/internal/domain"
)

func RegisterSearchingHandler(searchingHandlers ddd.EventHandler[ddd.AggregateEvent], domainSubscriber ddd.EventSubscriber[ddd.AggregateEvent]) {
	domainSubscriber.Subscribe(searchingHandlers,
		domain.OrderCreatedEvent,
		domain.OrderAddedItemEvent,
		domain.OrderCheckedOutEvent,
		domain.OrderReadiedEvent,
		domain.OrderCanceledEvent,
		domain.OrderCompletedEvent,
	)
}
