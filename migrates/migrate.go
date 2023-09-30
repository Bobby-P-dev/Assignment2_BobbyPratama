package main

import (
	"github.com/Bobby-P-dev/Assignment2_BobbyPratama.git/initializers"
	"github.com/Bobby-P-dev/Assignment2_BobbyPratama.git/models"
)

func init() {
	initializers.LoadEnvVariable()
	initializers.ConnectToDB()
}
func main() {
	initializers.DB.AutoMigrate(&models.Orders{})
	initializers.DB.AutoMigrate(&models.Items{})
}
