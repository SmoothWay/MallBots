package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/SmoothWay/MallBots/customers/internal/domain"
)

type CustomerRepository struct {
	tableName string
	db        *sql.DB
}

var _ domain.CustomerRepository = (*CustomerRepository)(nil)

func NewCustomerRepository(tableName string, db *sql.DB) *CustomerRepository {
	return &CustomerRepository{
		tableName: "customers",
		db:        db,
	}
}

func (r *CustomerRepository) Save(ctx context.Context, customer *domain.Customer) error {
	const query = `INSERT INTO %s (id, name, sms_number, enabled) VALUES ($1, $2, $3, $4)`

	_, err := r.db.ExecContext(ctx, r.table(query), customer.ID, customer.Name, customer.SmsNumber, customer.Enabled)
	if err != nil {
		return err
	}
	return nil
}

func (r *CustomerRepository) Find(ctx context.Context, customerId string) (*domain.Customer, error) {
	const query = `SELECT id, name, sms_number, enabled FROM %s WHERE id = $1`

	var customer domain.Customer

	err := r.db.QueryRowContext(ctx, r.table(query), customerId).Scan(&customer.ID, &customer.Name, &customer.SmsNumber, &customer.Enabled)
	if err != nil {
		return nil, err
	}
	return &customer, nil
}

func (r *CustomerRepository) Update(ctx context.Context, customer *domain.Customer) error {
	const query = `UPDATE %s SET name = $1, sms_number = $2, enabled = $3 WHERE id = $4`

	_, err := r.db.ExecContext(ctx, r.table(query), customer.Name, customer.SmsNumber, customer.Enabled, customer.ID)
	if err != nil {
		return err
	}
	return nil
}

func (r *CustomerRepository) table(query string) string {
	return fmt.Sprintf(query, r.tableName)
}
