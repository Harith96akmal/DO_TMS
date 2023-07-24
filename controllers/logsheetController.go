package controllers

import (
	"Delivery_Order/database"
	"Delivery_Order/functions"
	"Delivery_Order/models"
	"strings"

	"github.com/gin-gonic/gin"
)

func CreateLogsheet(c *gin.Context) {

	// Capture parameters
	var params struct {
		TruckNo string
		Dockets []string
	}

	// Bind parameters
	c.Bind(&params)

	var findDocket models.Docket
	var newLogsheet []models.Logsheet
	var newLogsheetNo string = functions.GenerateLogsheetNo()
	var docketNotExists string = "Docket(s) "

	// Loop for each Docket in parameter array
	for _, Docket := range params.Dockets {

		// Search Docket in database
		database.DB.Table("dockets").Where("order_no = ?", Docket).Find(&findDocket)
		if findDocket == (models.Docket{}) {
			docketNotExists = docketNotExists + Docket + ", "
		}

		// If Docket exists
		if len(docketNotExists) == 10 {

			// Query update Docket
			database.DB.Model(&findDocket).Where("order_no = ?", Docket).Updates(models.Docket{
				TruckNo:    params.TruckNo,
				LogsheetNo: newLogsheetNo,
			})

			// Append to array of Logsheets to be created
			newLogsheet = append(newLogsheet, models.Logsheet{
				LogsheetNo: newLogsheetNo,
				TruckNo:    params.TruckNo,
				OrderNo:    Docket,
			})
		}

		// Empty findDocket model
		findDocket = models.Docket{}
	}

	// If one or more Dockets from parameter does not exist
	if len(docketNotExists) > 10 {
		// Return Bad Request if Docket(s) not found
		c.JSON(400, gin.H{
			"message": docketNotExists[0:len(docketNotExists)-2] + " not found",
		})
		return
	}

	// Query create new Logsheet
	query := database.DB.Create(&newLogsheet)

	if query.Error != nil {
		// Return Bad Request if fail
		c.JSON(400, gin.H{
			"message": "Failed to create Logsheet",
		})
		return
	}

	// Search Logsheet in database
	var showDockets []models.Docket
	database.DB.Table("dockets").Where("logsheet_no = ?", newLogsheetNo).Find(&showDockets)

	if len(showDockets) == 0 {
		// Return Bad Request if not found
		c.JSON(400, gin.H{
			"message": "Failed to create Logsheet",
		})
	} else {
		// Return OK if found
		c.JSON(200, showDockets)
	}
}

func FindLogsheet(c *gin.Context) {

	// Capture parameter
	id := c.Param("logsheetNo")

	// Search Logsheet in database
	var findLogsheets []models.Logsheet
	database.DB.Table("logsheets").Where("logsheet_no = ?", id).Find(&findLogsheets)
	if len(findLogsheets) == 0 {
		// Return Bad Request if not found
		c.JSON(400, gin.H{
			"message": "Logsheet not found",
		})
		return
	}

	// Append Dockets included in Logsheet
	var docketNos []string
	for _, findLogsheet := range findLogsheets {
		docketNos = append(docketNos, findLogsheet.OrderNo)
	}

	// Search Dockets matched to Logsheet in database
	var findDockets []models.Docket
	database.DB.Table("dockets").Where("order_no in ('" + strings.Join(docketNos, "','") + "')").Find(&findDockets)

	if len(findDockets) == 0 {
		// Return Bad Request if not found
		c.JSON(200, gin.H{
			"message": "No Dockets included in Logsheet",
		})
	} else {
		// Return OK if found
		c.JSON(200, findDockets)
	}
}
