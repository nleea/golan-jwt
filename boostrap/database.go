package boostrap

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func DatabaseConnect(env *Env) *sql.DB {

	configDB := "postgresql://%s:%s@%s:%s/%s?sslmode=disable"
	configDB = fmt.Sprintf(configDB, env.DBUser, env.DBPassword, env.DBHost, env.DBPort, env.DBName)
	db, err := sql.Open("postgres", configDB)

	if err != nil {
		log.Fatal("Error to conect to db", err)
	}

	return db

}

func DatabaseDisconnect(sql *sql.DB) {
	if sql == nil {
		return
	}

	err := sql.Close()

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connection to postgrest is closed.")

}
