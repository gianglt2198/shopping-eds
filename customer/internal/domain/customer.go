package domain

import (
	"shopping/internal/ddd"

	"github.com/stackus/errors"
)

const CustomerAggregate = "customers.CustomerAggregate"

type Customer struct {
	ddd.Aggregate
	Name      string
	SmsNumber string
	Email     string
	Active    bool
}

func NewCustomer(id string) *Customer {
	return &Customer{
		Aggregate: ddd.NewAggregate(id, CustomerAggregate),
	}
}

func RegisterCustomer(id, name, smsNumber, email string) (*Customer, error) {
	if id == "" {
		return nil, errors.Wrap(errors.ErrBadRequest, "the customer id cannot be blank")
	}

	if name == "" {
		return nil, errors.Wrap(errors.ErrBadRequest, "the customer name cannot be blank")
	}

	if smsNumber == "" {
		return nil, errors.Wrap(errors.ErrBadRequest, "the SMS number cannot be blank")
	}

	if email == "" {
		return nil, errors.Wrap(errors.ErrBadRequest, "the email cannot be blank")
	}

	customer := NewCustomer(id)
	customer.Name = name
	customer.Active = true
	customer.Email = email
	customer.SmsNumber = smsNumber

	customer.AddEvent(CustomerRegisteredEvent, &CustomerRegistered{
		Customer: customer,
	})

	return customer, nil
}
