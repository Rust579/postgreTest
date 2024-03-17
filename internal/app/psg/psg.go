package psg

import (
	"Postgre/internal/repo"
	"Postgre/internal/service"
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

	if err = repo.CreatePersonTable(db); err != nil {
		log.Fatalln("Ошибка создания базы")
	}
	fmt.Printf("Database %s created successfully.\n", "mydatabase")

	if err = repo.InsertPersons(db, service.Persons); err != nil {
		log.Fatal("err InsertPersons", err)
	}
	if err = repo.InsertDepartments(db, service.Departments); err != nil {
		log.Fatal("err InsertDepartments", err)
	}
	if err = repo.InsertSpecialties(db, service.Specialties); err != nil {
		log.Fatal("err InsertSpecialties", err)
	}
	if err = repo.InsertGroups(db, service.Groups); err != nil {
		log.Fatal("err InsertGroups", err)
	}
	if err = repo.InsertStudents(db, service.Students); err != nil {
		log.Fatal("err InsertStudents", err)
	}
	if err = repo.InsertMovement(db, service.StudentMovements); err != nil {
		log.Fatal("err InsertMovement", err)
	}
	if err = repo.InsertAddresses(db, service.Addresses); err != nil {
		log.Fatal("err InsertAddresses", err)
	}
	if err = repo.InsertPhones(db, service.Phones); err != nil {
		log.Fatal("err InsertPhones", err)
	}

	fmt.Println("Tables filled in successfully")

	<-gfsd
}
