package model

import (
	"database/sql"
	"fmt"
	"log"
)

type CharacterRank struct {
	Username  string `json:"username"`
	Character string `json:"character"`
	Rank      int    `json:"rank"`
	Score     int    `json:"score"`
}

func SearchUser(db *sql.DB, query string) ([]CharacterRank, error) {
	var ranks []CharacterRank

	searchQuery := `
SELECT 
    a.username,
    SUM(s.reward_score) AS total_score,
    RANK() OVER (ORDER BY SUM(s.reward_score) DESC) AS rank
FROM 
    account a
JOIN 
    character c ON a.acc_id = c.acc_id
JOIN 
    scores s ON c.char_id = s.char_id
WHERE 
    a.username ILIKE $1
GROUP BY 
    a.acc_id, a.username
ORDER BY 
    total_score DESC;
`
	rows, err := db.Query(searchQuery, "%"+query+"%")
	if err != nil {
		log.Println("Error querying the database:", err)
		return nil, fmt.Errorf("could not search for characters")
	}
	defer rows.Close()

	for rows.Next() {
		var rank CharacterRank
		if err := rows.Scan(&rank.Username, &rank.Character, &rank.Score, &rank.Rank); err != nil {
			log.Println("Error scanning row:", err)
			continue
		}
		ranks = append(ranks, rank)
	}

	if err := rows.Err(); err != nil {
		log.Println("Error with rows:", err)
		return nil, fmt.Errorf("could not retrieve search results")
	}

	return ranks, nil
}
