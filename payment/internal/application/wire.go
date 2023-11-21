//go:build wireinject
// +build wireinject

package application

import (
	"database/sql"
	routerGRPC "shopping/payment/internal/application/router/grpc"
	"shopping/payment/internal/infra/repo"
	"shopping/payment/internal/logging"
	"shopping/payment/internal/usecase"

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
) error {
	panic(wire.Build(
		logging.LoggingSet,
		routerGRPC.PaymentGRPCServerSet,
		repo.RepositorySet,
		usecase.NewService,
		routerGRPC.OrderClientSet,
	))
}
