package models

import (
	"database/sql"
	"log"
	_ "github.com/lib/pq"
)

const databaseURL    = "postgres://url_shortener:password@localhost/url_shortener?sslmode=disable"
const databaseDriver = "postgres"

var databaseTables = [...]string{
	"WEB_URL(ID SERIAL PRIMARY KEY, URL TEXT NOT NULL)",
}

func InitializeDatabase() (*sql.DB, error) {
	var exception error
	log.Println(databaseTables)

	database , exception := sql.Open(databaseDriver, databaseURL)

	if exception != nil {
		log.Println(exception)
		return nil, exception
	}

	for index, databaseTable := range databaseTables {
		// Prepare statement for database create
		statement, exception := database .Prepare("CREATE TABLE " + databaseTable + ";")

		if exception != nil {
			log.Println(exception)
			return nil, exception
		}

		// Execute statement
		result, exception := statement.Exec()

		if exception != nil {
			log.Println(exception)
			return nil, exception
		}

		log.Println(index, result)
	}

	// Return database handle
	return database , nil
}