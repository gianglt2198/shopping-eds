package usecase

import (
	"context"
	"shopping/payment/internal/domain"

	"github.com/google/wire"
)

type (
	CreateInvoice struct {
		ID         string
		OrderID    string
		CustomerID string
		Amount     float64
	}

	GetInvoice struct {
		ID string
	}

	PayInvoice struct {
		ID string
	}

	CancelInvoice struct {
		ID string
	}

	ServiceUsecase interface {
		CreateInvoice(context.Context, CreateInvoice) error
		GetInvoice(context.Context, GetInvoice) (*domain.Invoice, error)
		PayInvoice(context.Context, PayInvoice) error
		CancelInvoice(context.Context, CancelInvoice) error
	}

	serviceUsecase struct {
		payments domain.PaymentRepository
		orders   domain.OrderRepository
	}
)

var _ ServiceUsecase = (*serviceUsecase)(nil)

var UseCaseSet = wire.NewSet(NewService)

func NewService(repo domain.PaymentRepository, orders domain.OrderRepository) ServiceUsecase {
	return &serviceUsecase{
		payments: repo,
		orders:   orders,
	}
}

func (a *serviceUsecase) CreateInvoice(ctx context.Context, create CreateInvoice) error {
	invoice, err := domain.CreateInvoice(create.ID, create.OrderID, create.CustomerID, create.Amount)
	if err != nil {
		return err
	}

	if err := a.payments.Save(ctx, invoice); err != nil {
		return err
	}

	if err := a.orders.ReadyOrder(ctx, invoice.OrderID, invoice.ID); err != nil {
		return err
	}

	return nil
}

func (a *serviceUsecase) GetInvoice(ctx context.Context, get GetInvoice) (*domain.Invoice, error) {
	return a.payments.Find(ctx, get.ID)
}

func (a *serviceUsecase) PayInvoice(ctx context.Context, update PayInvoice) error {
	// TODO: acts payment for this invoice
	payment, err := a.payments.Find(ctx, update.ID)
	if err != nil {
		return err
	}

	if err = payment.Pay(); err != nil {
		return err
	}

	if err := a.orders.CompleteOrder(ctx, payment.OrderID); err != nil {
		return err
	}

	return a.payments.Update(ctx, payment)
}

func (a *serviceUsecase) CancelInvoice(ctx context.Context, delete CancelInvoice) error {
	payment, err := a.payments.Find(ctx, delete.ID)
	if err != nil {
		return err
	}

	if err = payment.Cancel(); err != nil {
		return err
	}

	return a.payments.Update(ctx, payment)
}
