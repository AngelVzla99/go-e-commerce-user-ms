package service

import (
	"fmt"

	"github.com/AngelVzla99/e-commerce-user-ms/customer/repository"
)

func Init() {
	fmt.Println("Init (Service)")
}

func NewCustomerService() *CustomerService {
	return &CustomerService{
		repository: repository.NewCustomerRepository(),
	}
}

type CustomerService struct {
	repository *repository.CustomerRepository
}

func (s *CustomerService) GetCustomers() {
	fmt.Println("GetCustomers (Service)")
	s.repository.GetCustomers()
}
