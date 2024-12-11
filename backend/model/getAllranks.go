package model

import (
	"fmt"

	"github.com/Haziqhazri-hub/ricrym-assignment/db"
)

func GetAllRanks() ([]Rank, error) {
	query := `
		SELECT a.acc_id, a.username, a.email, COALESCE(SUM(s.reward_score), 0) AS total_score
		FROM Account AS a
		LEFT JOIN Character AS c ON a.acc_id = c.acc_id
		LEFT JOIN Scores AS s ON c.char_id = s.char_id
		GROUP BY a.acc_id, a.username, a.email
		ORDER BY total_score DESC
	`
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error fetching all ranks: %v", err)
	}
	defer rows.Close()

	var ranks []Rank
	var globalRank int64
	globalRank = 0

	for rows.Next() {
		globalRank++
		var r Rank
		err := rows.Scan(&r.Acc_id, &r.Username, &r.Email, &r.TotalScore)
		if err != nil {
			return nil, fmt.Errorf("error scanning rank data: %v", err)
		}

		r.Rank = globalRank
		ranks = append(ranks, r)
	}

	return ranks, nil
}
