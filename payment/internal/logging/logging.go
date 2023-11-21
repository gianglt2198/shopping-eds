package logging

import (
	"context"
	"shopping/payment/internal/domain"
	"shopping/payment/internal/usecase"

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

func (a Usecase) CreateInvoice(ctx context.Context, create usecase.CreateInvoice) (err error) {
	a.logger.Info().Msg("--> Payments.CreateInvoice")
	defer func() { a.logger.Info().Err(err).Msg("<-- Payments.CreateInvoice") }()
	return a.ServiceUsecase.CreateInvoice(ctx, create)
}

func (a Usecase) GetInvoice(ctx context.Context, get usecase.GetInvoice) (invoice *domain.Invoice, err error) {
	a.logger.Info().Msg("--> Payments.GetInvoice")
	defer func() { a.logger.Info().Err(err).Msg("<-- Payments.GetInvoice") }()
	return a.ServiceUsecase.GetInvoice(ctx, get)
}

func (a Usecase) PayInvoice(ctx context.Context, pay usecase.PayInvoice) (err error) {
	a.logger.Info().Msg("--> Payments.PayInvoice")
	defer func() { a.logger.Info().Err(err).Msg("<-- Payments.PayInvoice") }()
	return a.ServiceUsecase.PayInvoice(ctx, pay)
}

func (a Usecase) CancelInvoice(ctx context.Context, delete usecase.CancelInvoice) (err error) {
	a.logger.Info().Msg("--> Payments.CancelInvoice")
	defer func() { a.logger.Info().Err(err).Msg("<-- Payments.CancelInvoice") }()
	return a.ServiceUsecase.CancelInvoice(ctx, delete)
}
