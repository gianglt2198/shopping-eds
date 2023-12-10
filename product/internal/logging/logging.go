package logging

import (
	"context"
	"shopping/product/internal/domain"
	"shopping/product/internal/usecase"
	"shopping/product/internal/usecase/commands"
	"shopping/product/internal/usecase/queries"

	"github.com/google/wire"
	"github.com/rs/zerolog"
)

type Usecase struct {
	usecase.ServiceUsecase
	logger zerolog.Logger
}

var _ usecase.ServiceUsecase = (*Usecase)(nil)

var LoggingSet = wire.NewSet(LogApplicationAccess)

func LogApplicationAccess(application usecase.ServiceUsecase, logger zerolog.Logger) Usecase {
	return Usecase{
		ServiceUsecase: application,
		logger:         logger,
	}
}

func (a Usecase) CreateProduct(ctx context.Context, create commands.CreateProduct) (err error) {
	a.logger.Info().Msg("--> Product.CreateProduct")
	defer func() { a.logger.Info().Err(err).Msg("<-- Product.CreateProduct") }()
	return a.ServiceUsecase.CreateProduct(ctx, create)
}

func (a Usecase) GetProduct(ctx context.Context, get queries.GetProduct) (product *domain.ManagementProduct, err error) {
	a.logger.Info().Msg("--> Product.GetProduct")
	defer func() { a.logger.Info().Err(err).Msg("<-- Product.GetProduct") }()
	return a.ServiceUsecase.GetProduct(ctx, get)
}

func (a Usecase) DeleteProduct(ctx context.Context, delete commands.DeleteProduct) (err error) {
	a.logger.Info().Msg("--> Product.DeleteProduct")
	defer func() { a.logger.Info().Err(err).Msg("<-- Product.DeleteProduct") }()
	return a.ServiceUsecase.DeleteProduct(ctx, delete)
}

func (a Usecase) IncreasePrice(ctx context.Context, update commands.IncreasePrice) (err error) {
	a.logger.Info().Msg("--> Product.IncreasePrice")
	defer func() { a.logger.Info().Err(err).Msg("<-- Product.IncreasePrice") }()
	return a.ServiceUsecase.IncreasePriceProduct(ctx, update)
}

func (a Usecase) DecreasePrice(ctx context.Context, update commands.DecreasePrice) (err error) {
	a.logger.Info().Msg("--> Product.DecreasePrice")
	defer func() { a.logger.Info().Err(err).Msg("<-- Product.DecreasePrice") }()
	return a.ServiceUsecase.DecreasePriceProduct(ctx, update)
}
