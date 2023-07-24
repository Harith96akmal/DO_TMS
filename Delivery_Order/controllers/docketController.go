package controllers

import (
	"Delivery_Order/database"
	"Delivery_Order/functions"
	"Delivery_Order/models"

	"github.com/gin-gonic/gin"
)

func CreateDocket(c *gin.Context) {
	//Capture parameters
	var params struct {
		Customer      string
		PickUpPoint   string
		DeliveryPoint string
		Quantity      float64
		Volume        float64
	}

	//Bind parameters
	c.Bind(&params)

	// Generate new Docket
	var newDocketNo string = functions.GenerateOrderNo()
	newDocket := models.Docket{
		OrderNo:       newDocketNo,
		Customer:      params.Customer,
		PickUpPoint:   params.PickUpPoint,
		DeliveryPoint: params.DeliveryPoint,
		Quantity:      params.Quantity,
		Volume:        params.Volume,
		Status:        "Created",
		TruckNo:       "",
		LogsheetNo:    "",
	}

	// Query create new Docket
	query := database.DB.Create(&newDocket)

	if query.Error != nil {
		// Return Bad Request if fail
		c.JSON(400, gin.H{
			"message": "Failed to create Docket",
		})
		return
	}

	// Search Docket in database
	var findDocket models.Docket
	database.DB.Where("order_no = ?", newDocketNo).First(&findDocket)
	if findDocket == (models.Docket{}) {
		// Return Bad Request if not found
		c.JSON(400, gin.H{
			"message": "Failed to create Docket",
		})
	} else {
		// Return OK if found
		c.JSON(200, findDocket)
	}
}

func FindDocket(c *gin.Context) {
	// Capture parameter
	id := c.Param("orderNo")

	// Search Docket in database
	var findDocket models.Docket
	database.DB.Where("order_no = ?", id).First(&findDocket)
	if findDocket == (models.Docket{}) {
		// Return Bad Request if not found
		c.JSON(400, gin.H{
			"message": "Docket not found",
		})
		return
	} else {
		// Return OK if found
		c.JSON(200, findDocket)
	}
}

func ReadDocket(c *gin.Context) {
	// Search Docket in database
	var readDockets []models.Docket
	database.DB.Find(&readDockets)

	if len(readDockets) == 0 {
		// Return Bad Request if not found
		c.JSON(200, gin.H{
			"message": "No Dockets found",
		})
	} else {
		// Return OK if found
		c.JSON(200, readDockets)
	}
}
