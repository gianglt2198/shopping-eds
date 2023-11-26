package handlers

import (
	"shopping/internal/ddd"
	"shopping/order/internal/domain"
	"shopping/order/internal/usecase"
)

func RegisterPaymentHandlers(paymentHandlers usecase.DomainEventHandlers, domainSubscriber ddd.EventSubscriber) {
	domainSubscriber.Subscribe(domain.OrderCheckedout{}, paymentHandlers.OnOrderCheckedout)
	domainSubscriber.Subscribe(domain.OrderCancelled{}, paymentHandlers.OnOrderCancelled)
}
