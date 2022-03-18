package main

import (
	"checkoutservice/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	checkoutController := controllers.NewCheckoutController()

	router := gin.Default()
	router.POST("/checkouts", checkoutController.Create)
	router.Run()
}
