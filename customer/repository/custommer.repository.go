package repository

import "fmt"

type CustomerRepository struct {
}

func Init() {
	fmt.Println("Init (Repository)")
}

func NewCustomerRepository() *CustomerRepository {
	return &CustomerRepository{}
}

func (repository *CustomerRepository) GetCustomers() {
	fmt.Println("GetCustomers (Repository)")
}
