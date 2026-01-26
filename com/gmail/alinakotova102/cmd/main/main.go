package main

import (
	"log"

	"github.com/AlinaKlishyna/PostGIS-Study/com/gmail/alinakotova102/database"
)

func main() {
	db, err := database.GetDB()
	if err != nil {
		log.Fatal(err)
	}

	database.MigrateTables(db)

}
