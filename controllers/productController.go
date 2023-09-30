package controllers

import (
	"github.com/Bobby-P-dev/Assignment2_BobbyPratama.git/initializers"
	"github.com/Bobby-P-dev/Assignment2_BobbyPratama.git/models"
	"github.com/gin-gonic/gin"
)

func GetAll(c *gin.Context) {
	var get []models.Orders

	err := initializers.DB.Preload("Items").Find(&get).Error

	if err != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"get": get,
	})

}

func GetByID(c *gin.Context) {
	id := c.Param("id")
	GetByid := models.Orders{}

	err := initializers.DB.Preload("Items").First(&GetByid, id).Error

	if err != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"getbyid": GetByid,
	})

}

func PostOrders(c *gin.Context) {
	var ordered struct {
		CustomerName string
	}
	c.Bind(&ordered)

	//creat a post
	order := models.Orders{CustomerName: ordered.CustomerName}
	result := initializers.DB.Create(&order)

	if result.Error != nil {
		c.Status(400)
		return
	}
	c.JSON(200, gin.H{
		"order": order,
	})
}

func PostItems(c *gin.Context) {
	var itemOrdered struct {
		ItemCode    string
		Description string
		Quantity    int
		OrdersID    int
	}
	c.Bind(&itemOrdered)

	//create a post
	itemsOrder := models.Items{ItemCode: itemOrdered.ItemCode, Description: itemOrdered.Description, Quantity: itemOrdered.Quantity, OrdersID: itemOrdered.OrdersID}
	result := initializers.DB.Create(&itemsOrder)

	if result.Error != nil {
		c.Status(400)
		return
	}
	c.JSON(200, gin.H{
		"items": itemOrdered,
	})
}

func PutByIDItems(c *gin.Context) {
	id := c.Param("id")

	var itemOrdered struct {
		ItemCode    string
		Description string
		Quantity    int
	}

	c.Bind(&itemOrdered)

	var update models.Items
	initializers.DB.First(&update, id)

	initializers.DB.Model(&update).Updates(models.Items{
		ItemCode: itemOrdered.ItemCode, Description: itemOrdered.Description, Quantity: itemOrdered.Quantity})

	c.JSON(200, gin.H{
		"update": update,
	})

}

func PutByIDOrders(c *gin.Context) {
	id := c.Param("id")

	var ordered struct {
		CustomerName string
	}
	c.Bind(&ordered)

	var update models.Orders
	initializers.DB.First(&update, id)

	initializers.DB.Model(&update).Updates(models.Orders{
		CustomerName: ordered.CustomerName})

	c.JSON(200, gin.H{
		"update": update,
	})
}

func DeleteByID(c *gin.Context) {
	id := c.Param("id")

	initializers.DB.Delete(&models.Orders{}, id)

	c.Status(200)
}
