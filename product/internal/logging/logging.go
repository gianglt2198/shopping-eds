package logging

import (
	"context"
	"shopping/product/internal/domain"
	"shopping/product/internal/usecase"

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

func (a Usecase) CreateProduct(ctx context.Context, create usecase.CreateProduct) (err error) {
	a.logger.Info().Msg("--> Product.CreateProduct")
	defer func() { a.logger.Info().Err(err).Msg("<-- Product.CreateProduct") }()
	return a.ServiceUsecase.CreateProduct(ctx, create)
}

func (a Usecase) GetProduct(ctx context.Context, get usecase.GetProduct) (product *domain.Product, err error) {
	a.logger.Info().Msg("--> Product.GetProduct")
	defer func() { a.logger.Info().Err(err).Msg("<-- Product.GetProduct") }()
	return a.ServiceUsecase.GetProduct(ctx, get)
}

func (a Usecase) UpdateProduct(ctx context.Context, update usecase.UpdateProduct) (err error) {
	a.logger.Info().Msg("--> Product.UpdateProduct")
	defer func() { a.logger.Info().Err(err).Msg("<-- Product.UpdateProduct") }()
	return a.ServiceUsecase.UpdateProduct(ctx, update)
}

func (a Usecase) DeleteProduct(ctx context.Context, delete usecase.DeleteProduct) (err error) {
	a.logger.Info().Msg("--> Product.DeleteProduct")
	defer func() { a.logger.Info().Err(err).Msg("<-- Product.DeleteProduct") }()
	return a.ServiceUsecase.DeleteProduct(ctx, delete)
}
