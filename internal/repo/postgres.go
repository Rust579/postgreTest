package repo

import (
	"database/sql"
	"fmt"
	"log"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "mydatabase"
)

func ConnectToDB() (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

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

/*func CreateDatabase(db *sql.DB, dbname string) error {
	_, err := db.Exec(fmt.Sprintf("CREATE DATABASE %s", dbname))
	if err != nil {
		return err
	}
	return nil
}*/

func CreatePersonTable(db *sql.DB) error {
	query := `
		CREATE TABLE IF NOT EXISTS person (
			id SERIAL PRIMARY KEY,
			first_name VARCHAR(50),
			last_name VARCHAR(50),
			age INT,
			birth_date DATE
		)
	`

	_, err := db.Exec(query)
	if err != nil {
		return err
	}
	return nil
}
