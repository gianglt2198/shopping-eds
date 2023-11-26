package usecase

import (
	"context"
	"shopping/internal/ddd"
	"shopping/order/internal/domain"
	"shopping/order/internal/usecase/commands"
	"shopping/order/internal/usecase/queries"

	"github.com/google/wire"
)

type (
	ServiceUsecase interface {
		Commands
		Queries
	}
	Commands interface {
		CreateOrder(ctx context.Context, cmd commands.CreateOrder) error
		CancelOrder(ctx context.Context, cmd commands.CancelOrder) error
		CheckoutOrder(ctx context.Context, cmd commands.CheckoutOrder) error
		ReadyOrder(ctx context.Context, cmd commands.ReadyOrder) error
		CompleteOrder(ctx context.Context, cmd commands.CompleteOrder) error
		AddItem(ctx context.Context, cmd commands.AddItem) error
	}
	Queries interface {
		GetOrder(ctx context.Context, query queries.GetOrder) (*domain.Order, error)
	}

	serviceUsecase struct {
		usecaseCommands
		usecaseQueries
	}
	usecaseCommands struct {
		commands.CreateOrderHandler
		commands.CancelOrderHandler
		commands.ReadyOrderHandler
		commands.CompleteOrderHandler
		commands.AddItemHandler
		commands.CheckoutOrderHandler
	}
	usecaseQueries struct {
		queries.GetOrderHandler
	}
)

var _ ServiceUsecase = (*serviceUsecase)(nil)

var UseCaseSet = wire.NewSet(NewService)

func NewService(
	orders domain.OrderRepository,
	payments domain.PaymentRepository,
	customers domain.CustomerRepository,
	products domain.ProductRepository,
	domainPubliser ddd.EventPublisher) ServiceUsecase {
	return &serviceUsecase{
		usecaseCommands: usecaseCommands{
			CreateOrderHandler:   commands.NewCreateOrderHandler(orders, customers, domainPubliser),
			CancelOrderHandler:   commands.NewCancelOrderHandler(orders, payments, domainPubliser),
			ReadyOrderHandler:    commands.NewReadyOrderHandler(orders, payments, domainPubliser),
			CompleteOrderHandler: commands.NewCompleteOrderHandler(orders, domainPubliser),
			AddItemHandler:       commands.NewAddItemHandler(orders, products, domainPubliser),
			CheckoutOrderHandler: commands.NewCheckoutOrderHandler(orders, payments, domainPubliser),
		},
		usecaseQueries: usecaseQueries{
			GetOrderHandler: queries.NewGetOrderHandler(orders),
		},
	}
}
