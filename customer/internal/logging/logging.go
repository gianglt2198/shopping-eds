package logging

import (
	"context"
	"shopping/customer/internal/domain"
	"shopping/customer/internal/usecase"

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

func (a Usecase) RegisterCustomer(ctx context.Context, register usecase.RegisterCustomer) (err error) {
	a.logger.Info().Msg("--> Customers.RegisterCustomer")
	defer func() { a.logger.Info().Err(err).Msg("<-- Customers.RegisterCustomer") }()
	return a.ServiceUsecase.RegisterCustomer(ctx, register)
}

func (a Usecase) GetCustomer(ctx context.Context, get usecase.GetCustomer) (customer *domain.Customer, err error) {
	a.logger.Info().Msg("--> Customers.GetCustomer")
	defer func() { a.logger.Info().Err(err).Msg("<-- Customers.GetCustomer") }()
	return a.ServiceUsecase.GetCustomer(ctx, get)
}
