package db

import (
	"database/sql"
	"fmt"
	"log"
	"webapi/configs"

	_ "github.com/lib/pq"
)

func OpenConnection() (*sql.DB, error) {
	conf := configs.GetDB()

	sc := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", conf.Host, conf.Port, conf.User, conf.Pass, conf.Database)

	conn, err := sql.Open("postgres", sc)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	err = conn.Ping()
	if err != nil {
		log.Fatal(err)
	}

	return conn, err
}
