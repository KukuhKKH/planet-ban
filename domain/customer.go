package domain

import (
	"context"
	"time"
)

type Customer struct {
	ID       int       `db:"id"`
	Name     string    `db:"name"`
	Phone    string    `db:"phone"`
	CreadtAt time.Time `db:"created_at"`
}

type CustomerRepository interface {
	FindByID(ctx context.Context, id int) (Customer, error)
	FindByIds(ctx context.Context, ids []int) ([]Customer, error)
	FindByPhone(ctx context.Context, phone string) (Customer, error)
	Insert(ctx context.Context, customer *Customer) error
}

type CustomerService interface {
}
