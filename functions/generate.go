package functions

import (
	"DO_TMS/database"
	"DO_TMS/models"
	"fmt"
	"log"
	"strconv"
)

func GenerateOrderNo() string {
	var err error
	var lastDocket models.Docket
	var lastNo int
	var nextNo int
	var newOrderNo string

	// Search latest Docket entry
	database.DB.Table("dockets").Order("order_no desc").First(&lastDocket)

	if lastDocket == (models.Docket{}) {
		// Generate first Docket No if database empty
		newOrderNo = "TDN0001"
	} else {
		// Increment Docket No by 1 and generate new
		lastNo, err = (strconv.Atoi(lastDocket.OrderNo[3:]))
		nextNo = lastNo + 1
		newOrderNo = "TDN" + fmt.Sprintf("%04s", strconv.Itoa(nextNo))
	}

	if err != nil {
		log.Fatal(err)
	}

	return newOrderNo
}

func GenerateLogsheetNo() string {
	var err error
	var lastLogsheet models.Logsheet
	var lastNo int
	var nextNo int
	var newLogsheetNo string

	// Search latest Logsheet entry
	database.DB.Table("logsheets").Order("logsheet_no desc").First(&lastLogsheet)

	if lastLogsheet == (models.Logsheet{}) {
		// Generate first Logsheet No if database empty
		newLogsheetNo = "DT0001"
	} else {
		// Increment Logsheet No by 1 and generate new
		lastNo, err = (strconv.Atoi(lastLogsheet.LogsheetNo[2:]))
		nextNo = lastNo + 1
		newLogsheetNo = "DT" + fmt.Sprintf("%04s", strconv.Itoa(nextNo))
	}

	if err != nil {
		log.Fatal(err)
	}

	return newLogsheetNo
}
