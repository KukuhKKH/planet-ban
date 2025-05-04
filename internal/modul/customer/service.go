package customer

import (
	"context"
	"kukuhkkh.id/learn/bengkel/domain"
	"log"
	"strconv"
	"time"
)

type service struct {
	customerRepository domain.CustomerRepository
}

func NewService(customerRepository domain.CustomerRepository) domain.CustomerService {
	return &service{customerRepository}
}

func (s service) ALl(ctx context.Context) domain.ApiResponse {
	customers, err := s.customerRepository.FindAll(ctx)

	if err != nil {
		log.Printf("Error fetching customers: %s", err.Error())

		return domain.ApiResponse{
			Code:    "500",
			Message: "Internal Server Error",
			Data:    nil,
		}
	}

	var customerData []domain.CustomerData
	for _, customer := range customers {
		customerData = append(customerData, domain.CustomerData{
			ID:    strconv.Itoa(customer.ID),
			Name:  customer.Name,
			Phone: customer.Phone,
		})
	}

	return domain.ApiResponse{
		Code:    "200",
		Message: "Success",
		Data:    customerData,
	}
}

func (s service) Save(ctx context.Context, customerData domain.CustomerData) domain.ApiResponse {
	customer := domain.Customer{
		Name:      customerData.Name,
		Phone:     customerData.Phone,
		CreatedAt: time.Now(),
	}

	err := s.customerRepository.Insert(ctx, &customer)
	if err != nil {
		log.Printf("Error saving customer: %s", err.Error())

		return domain.ApiResponse{
			Code:    "500",
			Message: "Internal Server Error",
			Data:    nil,
		}
	}

	return domain.ApiResponse{
		Code:    "200",
		Message: "Success",
		Data:    customer,
	}
}
