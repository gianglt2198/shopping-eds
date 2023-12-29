package usecase

import (
	"context"
	"shopping/internal/ddd"
	"shopping/order/internal/domain"
	"shopping/order/internal/models"
)

type SearchingOrdertHandlers[T ddd.AggregateEvent] struct {
	searching domain.SearchingRepository
	cusomters domain.CustomerCacheRepository
	products  domain.ProductCacheRepository
}

var _ ddd.EventHandler[ddd.AggregateEvent] = (*SearchingOrdertHandlers[ddd.AggregateEvent])(nil)

func NewSearchingOrdertHandlers(searching domain.SearchingRepository, customers domain.CustomerCacheRepository, products domain.ProductCacheRepository) *SearchingOrdertHandlers[ddd.AggregateEvent] {
	return &SearchingOrdertHandlers[ddd.AggregateEvent]{
		searching: searching,
		cusomters: customers,
		products:  products,
	}
}

func (h SearchingOrdertHandlers[T]) HandleEvent(ctx context.Context, event T) error {
	switch event.EventName() {
	case domain.OrderCreatedEvent:
		return h.onOrderCreated(ctx, event)
	case domain.OrderAddedItemEvent:
		return h.onOrderAddedItem(ctx, event)
	case domain.OrderCheckedOutEvent:
		return h.onOrderCheckedOut(ctx, event)
	case domain.OrderReadiedEvent:
		return h.onOrderReadied(ctx, event)
	case domain.OrderCanceledEvent:
		return h.onOrderCanceled(ctx, event)
	case domain.OrderCompletedEvent:
		return h.onOrderCompleted(ctx, event)
	}

	return nil
}

func (h SearchingOrdertHandlers[T]) onOrderCreated(ctx context.Context, event T) error {
	payload := event.Payload().(*domain.OrderCreated)

	customer, err := h.cusomters.GetCustomer(ctx, payload.CustomerID)
	if err != nil {
		return err
	}
	return h.searching.Add(ctx, &models.Order{
		OrderID:      event.AggregateID(),
		CustomerID:   payload.CustomerID,
		CustomerName: customer.Name,
		Status:       domain.OrderStatusPending.String(),
	})
}

func (h SearchingOrdertHandlers[T]) onOrderAddedItem(ctx context.Context, event T) error {
	payload := event.Payload().(*domain.OrderAddedItem)
	product, err := h.products.GetProduct(ctx, payload.Item.ProductID)
	if err != nil {
		return err
	}

	return h.searching.UpdateItem(ctx, event.AggregateID(), payload.Item.ProductID, int(payload.Item.Quantity), product.Price)
}

func (h SearchingOrdertHandlers[T]) onOrderCheckedOut(ctx context.Context, event T) error {
	return h.searching.UpdateStatus(ctx, event.AggregateID(), domain.OrderStatusCheckedout.String())
}

func (h SearchingOrdertHandlers[T]) onOrderReadied(ctx context.Context, event T) error {
	return h.searching.UpdateStatus(ctx, event.AggregateID(), domain.OrderStatusReady.String())
}

func (h SearchingOrdertHandlers[T]) onOrderCanceled(ctx context.Context, event T) error {
	return h.searching.UpdateStatus(ctx, event.AggregateID(), domain.OrderStatusCancelled.String())
}

func (h SearchingOrdertHandlers[T]) onOrderCompleted(ctx context.Context, event T) error {
	return h.searching.UpdateStatus(ctx, event.AggregateID(), domain.OrderStatusCompleted.String())
}
