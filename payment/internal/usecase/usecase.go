package usecase

import (
	"context"
	"shopping/internal/ddd"
	"shopping/payment/internal/domain"
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
		payments  domain.PaymentRepository
		publisher ddd.EventPublisher[ddd.Event]
	}
)

var _ ServiceUsecase = (*serviceUsecase)(nil)

func NewService(repo domain.PaymentRepository, publisher ddd.EventPublisher[ddd.Event]) ServiceUsecase {
	return &serviceUsecase{
		payments:  repo,
		publisher: publisher,
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

	if err = a.publisher.Publish(ctx, ddd.NewEvent(domain.InvoiceCreatedEvent, &domain.InvoiceCreated{
		ID:      invoice.ID,
		OrderID: invoice.OrderID,
	})); err != nil {
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

	if err = a.publisher.Publish(ctx, ddd.NewEvent(domain.InvoicePaidEvent, &domain.InvoicePaid{
		ID:      payment.ID,
		OrderID: payment.OrderID,
	})); err != nil {
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
