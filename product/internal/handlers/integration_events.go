package handlers

import (
	"shopping/internal/ddd"
	"shopping/product/internal/domain"
)

func RegisterIntegrationHandlers(eventHandlers ddd.EventHandler[ddd.AggregateEvent], domainSubscriber ddd.EventSubscriber[ddd.AggregateEvent]) {
	domainSubscriber.Subscribe(eventHandlers,
		domain.ProductCreatedEvent,
		domain.ProductDeletedEvent,
		domain.ProductInscreasedPriceEvent,
		domain.ProductDescreasedPriceEvent,
	)
}
