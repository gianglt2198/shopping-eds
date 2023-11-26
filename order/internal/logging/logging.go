package logging

import (
	"context"
	"shopping/order/internal/domain"
	"shopping/order/internal/usecase"
	"shopping/order/internal/usecase/commands"
	"shopping/order/internal/usecase/queries"

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

func (a Usecase) CreateOrder(ctx context.Context, cmd commands.CreateOrder) (err error) {
	a.logger.Info().Msg("--> Ordering.CreateOrder")
	defer func() { a.logger.Info().Err(err).Msg("<-- Ordering.CreateOrder") }()
	return a.ServiceUsecase.CreateOrder(ctx, cmd)
}

func (a Usecase) CancelOrder(ctx context.Context, cmd commands.CancelOrder) (err error) {
	a.logger.Info().Msg("--> Ordering.CancelOrder")
	defer func() { a.logger.Info().Err(err).Msg("<-- Ordering.CancelOrder") }()
	return a.ServiceUsecase.CancelOrder(ctx, cmd)
}

func (a Usecase) CheckoutOrder(ctx context.Context, cmd commands.CheckoutOrder) (err error) {
	a.logger.Info().Msg("--> Ordering.CheckoutOrder")
	defer func() { a.logger.Info().Err(err).Msg("<-- Ordering.CheckoutOrder") }()
	return a.ServiceUsecase.CheckoutOrder(ctx, cmd)
}

func (a Usecase) ReadyOrder(ctx context.Context, cmd commands.ReadyOrder) (err error) {
	a.logger.Info().Msg("--> Ordering.ReadyOrder")
	defer func() { a.logger.Info().Err(err).Msg("<-- Ordering.ReadyOrder") }()
	return a.ServiceUsecase.ReadyOrder(ctx, cmd)
}

func (a Usecase) CompleteOrder(ctx context.Context, cmd commands.CompleteOrder) (err error) {
	a.logger.Info().Msg("--> Ordering.CompleteOrder")
	defer func() { a.logger.Info().Err(err).Msg("<-- Ordering.CompleteOrder") }()
	return a.ServiceUsecase.CompleteOrder(ctx, cmd)
}

func (a Usecase) GetOrder(ctx context.Context, query queries.GetOrder) (order *domain.Order, err error) {
	a.logger.Info().Msg("--> Ordering.GetOrder")
	defer func() { a.logger.Info().Err(err).Msg("<-- Ordering.GetOrder") }()
	return a.ServiceUsecase.GetOrder(ctx, query)
}

func (a Usecase) AddItem(ctx context.Context, cmd commands.AddItem) (err error) {
	a.logger.Info().Msg("--> Ordering.AddItem")
	defer func() { a.logger.Info().Err(err).Msg("<-- Ordering.AddItem") }()
	return a.ServiceUsecase.AddItem(ctx, cmd)
}
