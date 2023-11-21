package customer

import (
	"context"
	"shopping/internal/container"
	"shopping/payment/internal/application"
	router "shopping/payment/internal/application/router/grpc"
	rest "shopping/payment/internal/application/router/rest"
)

type Module struct{}

func (m Module) Startup(ctx context.Context, container container.Container) error {
	conn, err := router.Dial(ctx, container.Config().Rpc.Address())
	if err != nil {
		return err
	}

	application.InitApp("payment", container.DB(), container.RPC(), container.Logger(), conn)

	if err := rest.RegisterGateway(ctx, container.Mux(), container.Config().Rpc.Address()); err != nil {
		return err
	}
	if err := rest.RegisterSwagger(container.Mux()); err != nil {
		return err
	}

	return nil
}
