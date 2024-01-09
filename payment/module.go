package customer

import (
	"context"
	"shopping/internal/am"
	"shopping/internal/container"
	"shopping/internal/ddd"
	"shopping/internal/jetstream"
	"shopping/internal/registry"
	"shopping/order/orderspb"
	grpc "shopping/payment/internal/application/router/grpc"
	rest "shopping/payment/internal/application/router/rest"
	"shopping/payment/internal/handlers"
	"shopping/payment/internal/infra/repo"
	"shopping/payment/internal/logging"
	"shopping/payment/internal/usecase"
	"shopping/payment/paymentspb"
)

type Module struct{}

func (m Module) Startup(ctx context.Context, container container.Container) error {
	// setup Driven applications
	reg := registry.New()
	if err := orderspb.Registrations(reg); err != nil {
		return err
	}
	if err := paymentspb.Registration(reg); err != nil {
		return err
	}
	eventStream := am.NewEventStream(reg, jetstream.NewStream(container.Config().Nats.Stream, container.JS()))
	domainDispatcher := ddd.NewEventDispatcher[ddd.Event]()
	payment := repo.NewPaymentRepository("payments.payment", container.DB())

	// setup Applications
	paymentUsecase := logging.LogApplicationAccess(
		usecase.NewService(payment, domainDispatcher),
		container.Logger(),
	)
	orderHandlers := logging.LogEventHandlerAccess[ddd.Event](
		usecase.NewOrderHandlers(paymentUsecase),
		"Orders", container.Logger(),
	)
	integrationEventHandlers := logging.LogEventHandlerAccess[ddd.Event](
		usecase.NewIntegrationEventHandlers(eventStream),
		"IntegrationEvents", container.Logger(),
	)

	// setup Driver adapters
	if err := grpc.RegisterServer(paymentUsecase, container.RPC()); err != nil {
		return err
	}
	if err := rest.RegisterGateway(ctx, container.Mux(), container.Config().Rpc.Address()); err != nil {
		return err
	}
	if err := rest.RegisterSwagger(container.Mux()); err != nil {
		return err
	}
	if err := handlers.RegisterOrderHandlers(orderHandlers, eventStream); err != nil {
		return err
	}
	handlers.RegisterIntegrationEventHandlers(integrationEventHandlers, domainDispatcher)

	return nil
}
