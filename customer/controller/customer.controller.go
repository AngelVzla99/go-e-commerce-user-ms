package controller

import (
	"net/http"

	"github.com/AngelVzla99/e-commerce-user-ms/customer/dto"
	"github.com/AngelVzla99/e-commerce-user-ms/customer/service"
	"github.com/gin-gonic/gin"
)

type CustomerController struct {
	service *service.CustomerService
}

func NewCustomerController() *CustomerController {
	return &CustomerController{
		service: service.NewCustomerService(),
	}
}

func (c *CustomerController) CreateCustomers(request *gin.Context) {
	var dto dto.CreateCustomerDto
	if err := request.ShouldBindJSON(&dto); err != nil {
		request.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newCustomerId, error := c.service.CreateCustomers(dto)

	if error != nil {
		request.JSON(http.StatusInternalServerError, gin.H{"error": error.Error()})
		return
	}

	data := map[string]interface{}{
		"id": newCustomerId,
	}

	request.JSON(http.StatusCreated, gin.H{"message": "Customer created successfully", "data": data})
}
