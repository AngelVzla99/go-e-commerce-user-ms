package controller

import "fmt"

type CustomerController struct {
}

func (controller *CustomerController) GetCustomers() {
	fmt.Println("GetCustomers (Controller)")
}
