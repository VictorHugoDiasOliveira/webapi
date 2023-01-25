package models

import "webapi/db"

func Get(id int64) (person Person, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	row := conn.QueryRow(`SELECT * FROM person WHERE id=$1`, id)

	err = row.Scan(&person.ID, &person.Name, &person.Age)

	return person, err
}
