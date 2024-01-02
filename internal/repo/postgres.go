package repo

import (
	"database/sql"
	"fmt"
	"log"
)

const (
	Host     = "localhost"
	Port     = 5432
	User     = "postgres"
	Password = "postgres"
	Dbname   = "mydatabase"
)

func ConnectToDB() (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s Dbname=%s sslmode=disable",
		Host, Port, User, Password, Dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		fmt.Println("222")
		log.Fatal(err)
	}
	return db, nil
}

func CreateDatabase(db *sql.DB, dbname string) error {
	_, err := db.Exec(fmt.Sprintf("CREATE DATABASE %s", dbname))
	if err != nil {
		return err
	}
	return nil
}
