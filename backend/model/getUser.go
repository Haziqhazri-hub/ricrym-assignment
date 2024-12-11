package model

import (
	"database/sql"
	"fmt"
	"strings"
)

type UserRanking struct {
	Username   string `json:"username"`
	Rank       int    `json:"rank"`
	TotalScore int    `json:"total_score"`
}

func SearchUser(db *sql.DB, username string) ([]UserRanking, error) {
	query := `
		SELECT a.username, 
		       RANK() OVER (ORDER BY SUM(s.reward_score) DESC) AS rank, 
		       SUM(s.reward_score) AS total_score
		FROM account a
		LEFT JOIN character c ON a.acc_id = c.acc_id
		LEFT JOIN scores s ON c.char_id = s.char_id
		WHERE a.username ILIKE $1
		GROUP BY a.username
		ORDER BY rank
		LIMIT 10;
	`

	rows, err := db.Query(query, fmt.Sprintf("%%%s%%", strings.TrimSpace(username)))
	if err != nil {
		return nil, fmt.Errorf("error querying database: %v", err)
	}
	defer rows.Close()

	var rankings []UserRanking
	for rows.Next() {
		var ranking UserRanking
		if err := rows.Scan(&ranking.Username, &ranking.Rank, &ranking.TotalScore); err != nil {
			return nil, fmt.Errorf("error scanning row: %w", err)
		}
		rankings = append(rankings, ranking)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error processing rows: %w", err)
	}

	return rankings, nil
}