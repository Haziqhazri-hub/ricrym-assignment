package model

import (
	"database/db"
)

type Account struct {
	Acc_id          int64
	Username string    `binding:"required"`
	Email string    `binding:"required"`
}

var account = []Account{}

func GetAllAccount() ([]Account, error) {
	query := "SELECT acc_id, username, email FROM Account"
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var accounts []Account

	for rows.Next() {
		var account Account
		err := rows.Scan(&account.Acc_id, &account.Username, &account.Email)
		if err != nil {
			return nil, err
		}

		accounts = append(accounts, account)
	}
	return accounts, nil
}