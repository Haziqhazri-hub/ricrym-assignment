package db

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"os"

	"github.com/jaswdr/faker"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {
	dbURL := os.Getenv("DATABASE_URL")

	var err error
	DB, err = sql.Open("postgres", dbURL)
	if err != nil{
		log.Fatal("Failed to connect to database", err)
	}

	// defer DB.Close()
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	err = createTable(DB)
	if err != nil {
		log.Fatalf("Error creating tables: %v", err)
	}

	log.Println("Database initialized and tables created successfully!")

	generateFakeData()

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

func generateFakeData() {
    truncateQuery := "TRUNCATE TABLE Scores, Character, Account RESTART IDENTITY CASCADE"
    _, err := DB.Exec(truncateQuery)
    if err != nil {
        log.Fatalf("Failed to clear existing data: %v", err)
    }
    log.Println("Existing data cleared successfully.")

    faker := faker.New()

    for i := 1; i <= 900; i++ {
        username := faker.Person().FirstName()
        email := faker.Internet().Email()

        // Insert account and get acc_id
        var accID int
        query := "INSERT INTO Account (username, email) VALUES ($1, $2) RETURNING acc_id"
        err := DB.QueryRow(query, username, email).Scan(&accID)
        if err != nil {
            log.Fatalf("Failed to insert account: %v", err)
        }

        for j := 1; j <= 8; j++ {
            classID := rand.Intn(100) + 1

            // Insert character and get char_id
            var charID int
            charQuery := "INSERT INTO Character (acc_id, class_id) VALUES ($1, $2) RETURNING char_id"
            err := DB.QueryRow(charQuery, accID, classID).Scan(&charID)
            if err != nil {
                log.Fatalf("Failed to insert character: %v", err)
            }

            for k := 1; k <= 10; k++ {
                rewardScore := rand.Intn(1000) + 1

                // Insert scores
                _, err := DB.Exec("INSERT INTO Scores (char_id, reward_score) VALUES ($1, $2)", charID, rewardScore)
                if err != nil {
                    log.Fatalf("Failed to insert score: %v", err)
                }
            }
        }
    }

    log.Println("Fake data generated successfully!")
}
