//go:build wireinject
// +build wireinject

package application

import (
	"database/sql"
	"shopping/internal/ddd"
	routerGRPC "shopping/order/internal/application/router/grpc"
	"shopping/order/internal/infra/repo"
	"shopping/order/internal/logging"
	"shopping/order/internal/usecase"

	"github.com/google/wire"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"
)

func InitApp(
	tableName string,
	db *sql.DB,
	rpc *grpc.Server,
	logger zerolog.Logger,
	conn *grpc.ClientConn,
	publisher ddd.EventPublisher,
) error {
	panic(wire.Build(
		logging.LoggingSet,
		routerGRPC.OrderGRPCServerSet,
		repo.RepositorySet,
		usecase.NewService,
		routerGRPC.CustomerClientSet,
		routerGRPC.ProductClientSet,
		routerGRPC.PaymentClientSet,
	))
}
