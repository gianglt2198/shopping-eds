package customer

import (
	"context"
	"shopping/customer/customerspb"
	router_grpc "shopping/customer/internal/application/router/grpc"
	router_rest "shopping/customer/internal/application/router/rest"
	handlers "shopping/customer/internal/handler"
	"shopping/customer/internal/infra/repo"
	"shopping/customer/internal/logging"
	"shopping/customer/internal/usecase"
	"shopping/internal/am"
	"shopping/internal/container"
	"shopping/internal/ddd"
	"shopping/internal/jetstream"
	"shopping/internal/registry"
)

type Module struct{}

func (m Module) Startup(ctx context.Context, container container.Container) (err error) {
	// setup Driven Adapters
	reg := registry.New()
	if err = customerspb.Registrations(reg); err != nil {
		return err
	}
	eventStream := am.NewEventStream(reg, jetstream.NewStream(container.Config().Nats.Stream, container.JS()))
	domainDispatcher := ddd.NewEventDispatcher[ddd.AggregateEvent]()
	customers := repo.NewCustomerRepository("customers.customer", container.DB())

	// setup application
	app := logging.LogApplicationAccess(
		usecase.NewService(customers, domainDispatcher), container.Logger(),
	)
	integrationEventHandlers := logging.LogEventHandlerAccess[ddd.AggregateEvent](
		usecase.NewIntegrationEventHandlers(eventStream),
		"IntegrationEvents", container.Logger(),
	)
	if err = router_grpc.RegisterServer(app, container.RPC()); err != nil {
		return err
	}
	if err = router_rest.RegisterGateway(ctx, container.Mux(), container.Config().Rpc.Address()); err != nil {
		return err
	}
	if err = router_rest.RegisterSwagger(container.Mux()); err != nil {
		return err
	}
	handlers.RegisterIntegrationEventHandlers(integrationEventHandlers, domainDispatcher)

	return nil
}
