package models

import "webapi/db"

func GetAll() (persons []Person, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}

	defer conn.Close()

	rows, err := conn.Query(`SELECT * FROM person`)
	if err != nil {
		return
	}

	defer rows.Close()
	for rows.Next() {
		var person Person

		err = rows.Scan(&person.ID, &person.Name, &person.Age)
		if err != nil {
			continue
		}

		persons = append(persons, person)
	}

	return persons, err
}
