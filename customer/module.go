package customer

import (
	"context"
	"shopping/customer/internal/application"
	rest "shopping/customer/internal/application/router/rest"
	"shopping/internal/container"
)

type Module struct{}

func (m Module) Startup(ctx context.Context, container container.Container) error {

	application.InitApp("customer", container.DB(), container.RPC(), container.Logger())

	if err := rest.RegisterGateway(ctx, container.Mux(), container.Config().Rpc.Address()); err != nil {
		return err
	}
	if err := rest.RegisterSwagger(container.Mux()); err != nil {
		return err
	}

	return nil
}
