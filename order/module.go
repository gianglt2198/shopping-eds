package customer

import (
	"context"
	"shopping/internal/container"
	"shopping/internal/ddd"
	router "shopping/order/internal/application/router/grpc"
	rest "shopping/order/internal/application/router/rest"
	"shopping/order/internal/handlers"
	"shopping/order/internal/infra/repo"
	"shopping/order/internal/logging"
	"shopping/order/internal/usecase"
)

type Module struct{}

func (m Module) Startup(ctx context.Context, container container.Container) error {
	// setup Driven applications
	conn, err := router.Dial(ctx, container.Config().Rpc.Address())
	if err != nil {
		return err
	}
	domainDispatcher := ddd.NewEventDispatcher()
	orders := repo.NewOrderRepository("ordering", container.DB())
	products := router.NewProductRepository(conn)
	payments := router.NewPaymentRepository(conn)
	cutsomers := router.NewCustomerRepository(conn)

	// setup Applications
	app := logging.LogApplicationAccess(
		usecase.NewService(orders, payments, cutsomers, products, domainDispatcher),
		container.Logger(),
	)
	paymentHandlers := logging.LogDomainEventHandlerAccess(
		usecase.NewPaymentHandlers(payments),
		container.Logger(),
	)

	// setup Driver applications
	if err := router.RegisterServer(app, container.RPC()); err != nil {
		return err
	}
	if err := rest.RegisterGateway(ctx, container.Mux(), container.Config().Rpc.Address()); err != nil {
		return err
	}
	if err := rest.RegisterSwagger(container.Mux()); err != nil {
		return err
	}
	handlers.RegisterPaymentHandlers(paymentHandlers, domainDispatcher)
	return nil
}
