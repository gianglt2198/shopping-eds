//go:build wireinject
// +build wireinject

package application

import (
	"database/sql"
	routerGRPC "shopping/product/internal/application/router/grpc"
	"shopping/product/internal/infra/repo"
	"shopping/product/internal/logging"
	"shopping/product/internal/usecase"

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
		routerGRPC.ProductGRPCServerSet,
		repo.RepositorySet,
		usecase.NewService,
	))
}
