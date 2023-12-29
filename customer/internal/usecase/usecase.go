package usecase

import (
	"context"
	"shopping/customer/internal/domain"
	"shopping/internal/ddd"
)

type (
	RegisterCustomer struct {
		ID        string
		Name      string
		SmsNumber string
		Email     string
	}

	GetCustomer struct {
		ID string
	}

	ServiceUsecase interface {
		RegisterCustomer(context.Context, RegisterCustomer) error
		GetCustomer(context.Context, GetCustomer) (*domain.Customer, error)
	}

	serviceUsecase struct {
		customers       domain.CustomerRepository
		domainPublisher ddd.EventPublisher[ddd.AggregateEvent]
	}
)

var _ ServiceUsecase = (*serviceUsecase)(nil)

func NewService(repo domain.CustomerRepository, domainPublisher ddd.EventPublisher[ddd.AggregateEvent]) ServiceUsecase {
	return &serviceUsecase{
		customers:       repo,
		domainPublisher: domainPublisher,
	}
}

func (a *serviceUsecase) RegisterCustomer(ctx context.Context, register RegisterCustomer) error {
	customer, err := domain.RegisterCustomer(register.ID, register.Name, register.SmsNumber, register.Email)
	if err != nil {
		return err
	}

	if err = a.customers.Save(ctx, customer); err != nil {
		return err
	}

	// publish domain events
	if err = a.domainPublisher.Publish(ctx, customer.Events()...); err != nil {
		return err
	}

	return nil
}

func (a *serviceUsecase) GetCustomer(ctx context.Context, get GetCustomer) (*domain.Customer, error) {
	return a.customers.Find(ctx, get.ID)
}
