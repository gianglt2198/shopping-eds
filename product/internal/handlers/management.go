package handlers

import (
	"shopping/internal/ddd"
	"shopping/product/internal/domain"
)

func RegisterManagementHandler(managementHandlers ddd.EventHandler[ddd.AggregateEvent], domainSubscriber ddd.EventSubscriber[ddd.AggregateEvent]) {
	domainSubscriber.Subscribe(domain.ProductCreatedEvent, managementHandlers)
	domainSubscriber.Subscribe(domain.ProductDeletedEvent, managementHandlers)
	domainSubscriber.Subscribe(domain.ProductInscreasedPriceEvent, managementHandlers)
	domainSubscriber.Subscribe(domain.ProductDescreasedPriceEvent, managementHandlers)
}
