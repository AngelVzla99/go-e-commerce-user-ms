package controller

import (
	"fmt"

	"github.com/AngelVzla99/e-commerce-user-ms/customer/service"
	"github.com/gin-gonic/gin"
)

func Init() {
	fmt.Println("Init (Controller)")
}

type CustomerController struct {
	service *service.CustomerService
}

func NewCustomerController() *CustomerController {
	return &CustomerController{
		service: service.NewCustomerService(),
	}
}

func (c *CustomerController) GetCustomers(request *gin.Context) {
	fmt.Println("GetCustomers (Controller)")
	c.service.GetCustomers()
}
