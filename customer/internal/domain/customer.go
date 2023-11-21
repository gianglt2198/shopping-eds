package domain

import (
	"github.com/stackus/errors"
)

type Customer struct {
	ID        string
	Name      string
	SmsNumber string
	Email     string
	Active    bool
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

	return &Customer{
		ID:        id,
		Name:      name,
		SmsNumber: smsNumber,
		Email:     email,
		Active:    true,
	}, nil
}
