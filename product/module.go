package customer

import (
	"context"
	"shopping/internal/am"
	"shopping/internal/container"
	"shopping/internal/db"
	"shopping/internal/ddd"
	"shopping/internal/es"
	"shopping/internal/jetstream"
	"shopping/internal/registry"
	"shopping/internal/registry/serdes"
	grpc_router "shopping/product/internal/application/router/grpc"
	rest_router "shopping/product/internal/application/router/rest"
	"shopping/product/internal/domain"
	"shopping/product/internal/handlers"
	"shopping/product/internal/infra/repo"
	"shopping/product/internal/logging"
	"shopping/product/internal/usecase"
	"shopping/product/productspb"
)

type Module struct{}

func (m Module) Startup(ctx context.Context, container container.Container) error {
	reg := registry.New()
	err := registrations(reg)
	if err != nil {
		return err
	}
	if err = productspb.Registration(reg); err != nil {
		return err
	}
	eventStream := am.NewEventStream(reg, jetstream.NewStream(container.Config().Nats.Stream, container.JS()))
	domainDispatcher := ddd.NewEventDispatcher[ddd.AggregateEvent]()
	aggregateProduct := es.AggreagteStoreWithMiddleware(
		db.NewEventStore("products.events", container.DB(), reg),
		es.NewEventPublisher(domainDispatcher),
		db.NewSnapshotStore("products.snapshots", container.DB(), reg),
	)
	products := es.NewAggregateRepository[*domain.Product](domain.ProductAggregate, reg, aggregateProduct)
	management := repo.NewManagementRepository("products.product", container.DB())

	// setup application
	app := logging.LogApplicationAccess(
		usecase.NewService(products, management),
		container.Logger(),
	)
	managementHandlers := logging.LogEventHandlerAccess[ddd.AggregateEvent](
		usecase.NewManagementHandlers(management),
		"Management", container.Logger(),
	)
	integrationEventHandlers := logging.LogEventHandlerAccess[ddd.AggregateEvent](
		usecase.NewIntegrationEventHandlers(eventStream),
		"IntegrationEvents", container.Logger(),
	)

	// setup Driver adapters
	if err := grpc_router.RegisterServer(app, container.RPC()); err != nil {
		return err
	}
	if err := rest_router.RegisterGateway(ctx, container.Mux(), container.Config().Rpc.Address()); err != nil {
		return err
	}
	if err := rest_router.RegisterSwagger(container.Mux()); err != nil {
		return err
	}
	handlers.RegisterManagementHandler(managementHandlers, domainDispatcher)
	handlers.RegisterIntegrationHandlers(integrationEventHandlers, domainDispatcher)
	return nil
}

func registrations(reg registry.Registry) (err error) {
	serde := serdes.NewJsonSerde(reg)

	err = serde.Register(domain.Product{}, func(v interface{}) error {
		product := v.(*domain.Product)
		product.Aggregate = es.NewAggregate("", domain.ProductAggregate)
		return nil
	})
	if err != nil {
		return err
	}

	if err = serde.Register(domain.ProductCreated{}); err != nil {
		return err
	}
	if err = serde.Register(domain.ProductDeleted{}); err != nil {
		return err
	}
	if err = serde.RegisterKey(domain.ProductInscreasedPriceEvent, domain.ProductPriceChanged{}); err != nil {
		return err
	}
	if err = serde.RegisterKey(domain.ProductDescreasedPriceEvent, domain.ProductPriceChanged{}); err != nil {
		return err
	}
	if err = serde.RegisterKey(domain.ProductV1{}.SnapshotName(), domain.ProductV1{}); err != nil {
		return err
	}
	return nil
}
