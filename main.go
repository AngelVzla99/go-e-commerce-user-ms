package main

import (
	customerRoutes "github.com/AngelVzla99/e-commerce-user-ms/customer/route"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	customerRoutes.SetupRouter(router)

	router.Run(":8080")
}
