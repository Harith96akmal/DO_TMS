package migration

import (
	"DO_TMS/database"
	"io/ioutil"
	"log"
)

func main() {
	var err error
	createDockets, err := ioutil.ReadFile("database/migration/dockets.sql")
	if err != nil {
		log.Fatal(err)
	}

	createLogsheets, err := ioutil.ReadFile("database/migration/logsheets.sql")
	if err != nil {
		log.Fatal(err)
	}

	// database.DB.AutoMigrate()

	// Create tables in database using SQL files
	database.DB.Exec(string(createDockets))
	database.DB.Exec(string(createLogsheets))
}

func MigrateDB() {
	main()
}
