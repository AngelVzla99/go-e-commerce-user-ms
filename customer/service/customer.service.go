package service

import (
	"fmt"

	"github.com/yourusername/my-gin-project/customer/repository"
)

type CustomerService struct {
	repository *repository.CustomerRepository
}

func (service *CustomerService) GetCustomers() {
	fmt.Println("GetCustomers (Service)")
}
