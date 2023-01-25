package models

import "webapi/db"

func Update(id int64, person Person) (int64, error) {

	conn, err := db.OpenConnection()
	if err != nil {
		return 0, err
	}
	defer conn.Close()

	res, err := conn.Exec(`UPDATE person SET name=$1, age=$2 WHERE id=$3`, person.Name, person.Age, id)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}
