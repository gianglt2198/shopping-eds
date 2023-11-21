//go:build wireinject
// +build wireinject

package application

import (
	"database/sql"
	routerGRPC "shopping/customer/internal/application/router/grpc"
	"shopping/customer/internal/infra/repo"
	"shopping/customer/internal/logging"
	"shopping/customer/internal/usecase"

	"github.com/google/wire"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"
)

func InitApp(
	tableName string,
	db *sql.DB,
	rpc *grpc.Server,
	logger zerolog.Logger,
) error {
	panic(wire.Build(
		logging.LoggingSet,
		routerGRPC.CustomerGRPCServerSet,
		repo.RepositorySet,
		usecase.NewService,
	))
}
