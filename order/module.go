package customer

import (
	"context"
	"shopping/customer/customerspb"
	"shopping/internal/am"
	"shopping/internal/container"
	"shopping/internal/db"
	"shopping/internal/ddd"
	"shopping/internal/es"
	"shopping/internal/jetstream"
	"shopping/internal/registry"
	"shopping/internal/registry/serdes"
	"shopping/product/productspb"

	grpc_router "shopping/order/internal/application/router/grpc"
	rest_router "shopping/order/internal/application/router/rest"
	"shopping/order/internal/domain"
	"shopping/order/internal/handlers"
	"shopping/order/internal/infra/repo"
	"shopping/order/internal/logging"
	"shopping/order/internal/usecase"
)

type Module struct{}

func (m Module) Startup(ctx context.Context, container container.Container) (err error) {
	// setup Driven applications
	reg := registry.New()
	if err = registration(reg); err != nil {
		return err
	}
	if err = productspb.Registration(reg); err != nil {
		return err
	}
	if err = customerspb.Registrations(reg); err != nil {
		return err
	}
	eventStream := am.NewEventStream(reg, jetstream.NewStream(container.Config().Nats.Stream, container.JS()))
	domainDispatcher := ddd.NewEventDispatcher[ddd.AggregateEvent]()
	aggregateStore := es.AggreagteStoreWithMiddleware(
		db.NewEventStore("ordering.events", container.DB(), reg),
		es.NewEventPublisher(domainDispatcher),
		db.NewSnapshotStore("ordering.snapshots", container.DB(), reg),
	)

	orders := es.NewAggregateRepository[*domain.Order](domain.OrderAggregate, reg, aggregateStore)
	conn, err := grpc_router.Dial(ctx, container.Config().Rpc.Address())
	if err != nil {
		return err
	}
	payments := grpc_router.NewPaymentRepository(conn)
	customers := repo.NewCustomerCacheRepository("ordering.customers_cache", container.DB(), grpc_router.NewCustomerRepository(conn))
	products := repo.NewProductCacheRepository(ctx, "ordering.products_cache", container.DB(), grpc_router.NewProductRepository(conn))
	searching := repo.NewSearchingRepository("ordering.order", container.DB())
	// setup Applications
	app := logging.LogApplicationAccess(usecase.NewService(orders, payments, customers, products, searching), container.Logger())

	paymentHandlers := logging.LogEventHandlerAccess[ddd.AggregateEvent](
		usecase.NewPaymentHandlers(payments),
		"Payments",
		container.Logger(),
	)
	productHandlers := logging.LogEventHandlerAccess[ddd.Event](
		usecase.NewProductHandlers(products),
		"Product", container.Logger(),
	)
	customerHandlers := logging.LogEventHandlerAccess[ddd.Event](
		usecase.NewCustomerHandlers(customers),
		"Customer", container.Logger(),
	)
	searchingHandlers := logging.LogEventHandlerAccess[ddd.AggregateEvent](
		usecase.NewSearchingOrdertHandlers(searching, customers, products),
		"Searching", container.Logger(),
	)

	// setup Driver applications
	if err := grpc_router.RegisterServer(app, container.RPC()); err != nil {
		return err
	}
	if err := rest_router.RegisterGateway(ctx, container.Mux(), container.Config().Rpc.Address()); err != nil {
		return err
	}
	if err := rest_router.RegisterSwagger(container.Mux()); err != nil {
		return err
	}
	handlers.RegisterSearchingHandler(searchingHandlers, domainDispatcher)
	handlers.RegisterPaymentHandlers(paymentHandlers, domainDispatcher)
	if err = handlers.RegisterProductHandlers(productHandlers, eventStream); err != nil {
		return err
	}
	if err = handlers.RegisterCustomerHandlers(customerHandlers, eventStream); err != nil {
		return err
	}
	return nil
}

func registration(reg registry.Registry) error {
	serde := serdes.NewJsonSerde(reg)

	if err := serde.Register(domain.Order{}, func(v any) error {
		order := v.(*domain.Order)
		order.Aggregate = es.NewAggregate("", domain.OrderAggregate)
		return nil
	}); err != nil {
		return err
	}

	if err := serde.Register(domain.OrderCreated{}); err != nil {
		return err
	}
	if err := serde.Register(domain.OrderAddedItem{}); err != nil {
		return err
	}
	if err := serde.Register(domain.OrderCheckedout{}); err != nil {
		return err
	}
	if err := serde.Register(domain.OrderReadied{}); err != nil {
		return err
	}
	if err := serde.Register(domain.OrderCompleted{}); err != nil {
		return err
	}
	if err := serde.Register(domain.OrderCancelled{}); err != nil {
		return err
	}

	if err := serde.RegisterKey(domain.OrderV1{}.SnapshotName(), domain.OrderV1{}); err != nil {
		return err
	}

	return nil
}
