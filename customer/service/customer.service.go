package service

import (
	"github.com/AngelVzla99/e-commerce-user-ms/customer/dto"
	"github.com/AngelVzla99/e-commerce-user-ms/customer/mapper"
	"github.com/AngelVzla99/e-commerce-user-ms/customer/repository"
)

func NewCustomerService() *CustomerService {
	return &CustomerService{
		repository: repository.NewCustomerRepository(),
	}
}

type CustomerService struct {
	repository *repository.CustomerRepository
}

func (s *CustomerService) CreateCustomers(dto dto.CreateCustomerDto) (string, error) {
	customerId, error := s.repository.CreateCustomers(mapper.MapCustomerDtoToDbe(dto))

	if error != nil {
		return "", error
	}

	return customerId, nil
}
