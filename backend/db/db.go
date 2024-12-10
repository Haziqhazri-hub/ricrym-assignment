package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)



func InitDB() {
	DB, err := sql.Open("postgres", "host=localhost port=5432 dbname=db sslmode=disable")
	if err != nil{
		log.Fatal("Failed to connect to database", err)
	}

	// defer DB.Close()

	err = createTable(DB)
	if err != nil {
		log.Fatalf("Error creating tables: %v", err)
	}

	log.Println("Database initialized and tables created successfully!")

}

func createTable(db *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS Account (
	acc_id SERIAL PRIMARY KEY,
	username VARCHAR(255) NOT NULL,
	email VARCHAR(255) NOT NULL
	);

	CREATE TABLE IF NOT EXISTS Character (
	char_id SERIAL PRIMARY KEY,
	acc_id INT NOT NULL,
	class_id INT NOT NULL,
	FOREIGN KEY (acc_id) REFERENCES Account (acc_id) ON DELETE CASCADE
	);

	CREATE TABLE IF NOT EXISTS Scores (
	score_id SERIAL PRIMARY KEY,
	char_id INT NOT NULL,
	reward_score INT NOT NULL,
	FOREIGN KEY (char_id) REFERENCES Character (char_id) ON DELETE CASCADE
	)
	`

	_, err := db.Exec(query)

	if err != nil {
		return fmt.Errorf("error creating table: %w", err)
	}
	return nil
}