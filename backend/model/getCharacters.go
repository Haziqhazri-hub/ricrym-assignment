package model

import (
	"database/db"
)

type Character struct {
	Char_id          int64
	Acc_id int64    
	Class_id int64   
}

var character = []Character{}

func GetAllCharacter() ([]Character, error) {
	query := "SELECT char_id, acc_id, class_id FROM Character"
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var characters []Character

	for rows.Next() {
		var character Character
		err := rows.Scan(&character.Char_id, &character.Acc_id, &character.Class_id)
		if err != nil {
			return nil, err
		}

		characters = append(characters, character)
	}
	return characters, nil
}