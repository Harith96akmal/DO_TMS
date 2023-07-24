package main

import (
	"Delivery_Order/controllers"
	"Delivery_Order/database"
	"Delivery_Order/database/migration"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
	database.ConnectionDB()
	migration.MigrateDB()
}

func main() {
	r := gin.Default()

	r.POST("/docket", controllers.CreateDocket)
	r.GET("/docket/:orderNo", controllers.FindDocket)
	r.GET("/docket", controllers.ReadDocket)

	r.POST("/logsheet", controllers.CreateLogsheet)
	r.GET("/logsheet/:logsheetNo", controllers.FindLogsheet)

	r.Run() // listen and serve on localhost:8080
}
