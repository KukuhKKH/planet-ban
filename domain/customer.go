package domain

import "time"

type Customer struct {
	ID       int       `db:"id"`
	Name     string    `db:"name"`
	Phone    string    `db:"phone"`
	CreadtAt time.Time `db:"created_at"`
}

type CustomerRepository interface {
}

type CustomerService interface {
}
