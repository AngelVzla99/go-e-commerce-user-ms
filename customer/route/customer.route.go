package route

import (
	controller "github.com/AngelVzla99/e-commerce-user-ms/customer/controller"
	"github.com/gin-gonic/gin"
)

func Init() {}

func SetupRouter(router *gin.Engine) {
	customer := router.Group("/customer")
	controllerStruct := controller.NewCustomerController()

	customer.POST("", controllerStruct.CreateCustomers)
	customer.GET("/:id", controllerStruct.GetCustomer)
}
