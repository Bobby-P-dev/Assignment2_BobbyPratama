package main

import (
	"github.com/Bobby-P-dev/Assignment2_BobbyPratama.git/controllers"
	"github.com/Bobby-P-dev/Assignment2_BobbyPratama.git/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariable()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()
	r.GET("/api/order", controllers.GetAll)
	r.GET("/api/order/:id", controllers.GetByID)
	r.POST("/api/order", controllers.PostOrders)
	r.POST("/api/order/items", controllers.PostItems)
	r.PUT("/api/items/:id", controllers.PutByIDItems)
	r.PUT("/api/order/:id", controllers.PutByIDOrders)
	r.DELETE("/api/order/:id", controllers.DeleteByID)

	r.Run()
}
