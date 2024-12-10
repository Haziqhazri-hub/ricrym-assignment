package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)



func InitDB() {
	DB, err := sql.Open("postgres", "host=localhost port=5432 user=user password=user name=db sslmode=disable")
	if err != nil{
		log.Fatal("Failed to connect to database", err)
	}

	DB.Close()

}