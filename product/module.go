package customer

import (
	"context"
	"shopping/internal/container"
	"shopping/internal/ddd"
	"shopping/product/internal/application"
	rest "shopping/product/internal/application/router/rest"
)

type Module struct{}

func (m Module) Startup(ctx context.Context, container container.Container) error {

	// setup Driven adapters
	domainDispatcher := ddd.NewEventDispatcher()

	// setup Application
	application.InitApp("product", container.DB(), container.RPC(), container.Logger(), domainDispatcher)

	// setup Driver adapters
	if err := rest.RegisterGateway(ctx, container.Mux(), container.Config().Rpc.Address()); err != nil {
		return err
	}
	if err := rest.RegisterSwagger(container.Mux()); err != nil {
		return err
	}

	return nil
}
