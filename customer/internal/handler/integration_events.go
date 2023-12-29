package handlers

import (
	"shopping/customer/internal/domain"
	"shopping/internal/ddd"
)

func RegisterIntegrationEventHandlers(eventHandlers ddd.EventHandler[ddd.AggregateEvent], domainSubscriber ddd.EventSubscriber[ddd.AggregateEvent]) {
	domainSubscriber.Subscribe(eventHandlers,
		domain.CustomerRegisteredEvent,
	)
}
