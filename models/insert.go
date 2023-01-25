package models

import (
	"webapi/db"
)

func Insert(person Person) (id int64, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	sql := `INSERT INTO person (name, age) VALUES ($1, $2) RETURNING id`

	err = conn.QueryRow(sql, person.Name, person.Age).Scan(&id)

	return
}
