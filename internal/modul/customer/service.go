package customer

import "kukuhkkh.id/learn/bengkel/domain"

type service struct {
	customerRepository domain.CustomerRepository
}

func NewService(customerRepository domain.CustomerRepository) domain.CustomerService {
	return &service{customerRepository}
}
