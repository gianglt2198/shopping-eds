package usecase

import (
	"context"
	"shopping/customer/internal/domain"

	"github.com/google/wire"
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
		customers domain.CustomerRepository
	}
)

var _ ServiceUsecase = (*serviceUsecase)(nil)

var UseCaseSet = wire.NewSet(NewService)

func NewService(repo domain.CustomerRepository) ServiceUsecase {
	return &serviceUsecase{
		customers: repo,
	}
}

func (a *serviceUsecase) RegisterCustomer(ctx context.Context, register RegisterCustomer) error {
	customer, err := domain.RegisterCustomer(register.ID, register.Name, register.SmsNumber, register.Email)
	if err != nil {
		return err
	}

	return a.customers.Save(ctx, customer)
}

func (a *serviceUsecase) GetCustomer(ctx context.Context, get GetCustomer) (*domain.Customer, error) {
	return a.customers.Find(ctx, get.ID)
}
