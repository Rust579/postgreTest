package psg

import (
	"Postgre/internal/repo"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var gfsd = make(chan bool)

func Run() {

	db, err := repo.ConnectToDB()
	if err != nil {
		log.Fatalln("Нет соединения  с базой psg")
	}

	fmt.Println("Successfully connected!")

	if err = repo.CreateDatabase(db, "PsgDb"); err != nil {
		log.Fatalln("Ошибка создания базы")
	}
	fmt.Printf("Database %s created successfully.\n", repo.Dbname)

	<-gfsd
}
