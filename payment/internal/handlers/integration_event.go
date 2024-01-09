package handlers

import (
	"shopping/internal/ddd"
	"shopping/payment/internal/domain"
)

func RegisterIntegrationEventHandlers(eventHandlers ddd.EventHandler[ddd.Event], domainSubscriber ddd.EventSubscriber[ddd.Event]) {
	domainSubscriber.Subscribe(eventHandlers,
		domain.InvoiceCreatedEvent,
		domain.InvoicePaidEvent,
	)
}
