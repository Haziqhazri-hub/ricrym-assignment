package model

import (
	"fmt"
	"math"

	"github.com/Haziqhazri-hub/ricrym-assignment/db"
)

type Rank struct {
	Rank       int64   `json:"rank"`
	Acc_id     int64   `json:"acc_id"`
	Username   string  `json:"username"`
	Email      string  `json:"email"`
	TotalScore int64   `json:"total_score"`
}

func GetPaginatedRank(page int, pageSize int) ([]Rank, int, error) {
	offset := (page - 1) * pageSize
	query := `
		SELECT a.acc_id, a.username, a.email, COALESCE(SUM(s.reward_score), 0) AS total_score
		FROM Account AS a
		LEFT JOIN Character AS c ON a.acc_id = c.acc_id
		LEFT JOIN Scores AS s ON c.char_id = s.char_id
		GROUP BY a.acc_id, a.username, a.email
		ORDER BY total_score DESC
		LIMIT $1 OFFSET $2
	`
	rows, err := db.DB.Query(query, pageSize, offset)
	if err != nil {
		return nil, 0, fmt.Errorf("error fetching paginated ranks: %v", err)
	}
	defer rows.Close()

	var ranks []Rank

	var globalRank int64
	globalRank = int64((page - 1) * pageSize)

	for rows.Next() {
		globalRank++
		var r Rank
		err := rows.Scan(&r.Acc_id, &r.Username, &r.Email, &r.TotalScore)
		if err != nil {
			return nil, 0, fmt.Errorf("error scanning rank data: %v", err)
		}

		r.Rank = globalRank
		ranks = append(ranks, r)
	}

	var totalRecords int
	countQuery := `
		SELECT COUNT(DISTINCT a.acc_id)
		FROM Account AS a
		LEFT JOIN Character AS c ON a.acc_id = c.acc_id
		LEFT JOIN Scores AS s ON c.char_id = s.char_id
	`
	err = db.DB.QueryRow(countQuery).Scan(&totalRecords)
	if err != nil {
		return nil, 0, fmt.Errorf("error getting total records: %v", err)
	}

	totalPages := int(math.Ceil(float64(totalRecords) / float64(pageSize)))

	return ranks, totalPages, nil
}
