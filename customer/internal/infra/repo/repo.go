package repo

import (
	"context"
	"database/sql"
	"fmt"
	"shopping/customer/internal/domain"

	"github.com/google/wire"
)

type CustomerRepository struct {
	tableName string
	db        *sql.DB
}

var _ domain.CustomerRepository = (*CustomerRepository)(nil)

var RepositorySet = wire.NewSet(NewCustomerRepository)

func NewCustomerRepository(tableName string, db *sql.DB) domain.CustomerRepository {
	return CustomerRepository{
		tableName: tableName,
		db:        db,
	}
}

func (r CustomerRepository) Save(ctx context.Context, customer *domain.Customer) error {
	const query = "INSERT INTO %s (id, name, sms_number, email, active) VALUES ($1, $2, $3, $4, $5)"

	_, err := r.db.ExecContext(ctx, r.table(query), customer.ID, customer.Name, customer.SmsNumber, customer.Email, customer.Active)

	return err
}

func (r CustomerRepository) Find(ctx context.Context, customerID string) (*domain.Customer, error) {
	const query = "SELECT name, sms_number, email, active FROM %s WHERE id = $1 LIMIT 1"

	customer := &domain.Customer{
		ID: customerID,
	}

	err := r.db.QueryRowContext(ctx, r.table(query), customerID).Scan(&customer.Name, &customer.SmsNumber, &customer.Email, &customer.Active)

	return customer, err
}

func (r CustomerRepository) table(query string) string {
	return fmt.Sprintf(query, r.tableName)
}
