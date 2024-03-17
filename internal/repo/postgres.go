package repo

import (
	"Postgre/internal/service"
	"database/sql"
	"fmt"
	"log"
	"strconv"
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

func InsertStudents(db *sql.DB, data []service.Student) error {

	for _, s := range data {

		query := `
			SELECT COUNT(*)
			FROM students
			WHERE id_student = $1
		`
		var count int
		err := db.QueryRow(query, s.ID).Scan(&count)
		if err != nil {
			return err
		}
		if count > 0 {
			continue
		}

		query = `
		INSERT INTO students (id_student, id_person, stud_number, id_group)
		VALUES ($1, $2, $3, $4)
	`
		_, err = db.Exec(query, s.ID, s.PersonID, s.StudNumber, s.GroupID)
		if err != nil {
			return err
		}
	}
	return nil
}

func InsertDepartments(db *sql.DB, data []service.Department) error {

	for _, s := range data {

		query := `
			SELECT COUNT(*)
			FROM departments
			WHERE id_department = $1
		`
		var count int
		err := db.QueryRow(query, s.ID).Scan(&count)
		if err != nil {
			return err
		}
		if count > 0 {
			continue
		}

		query = `
		INSERT INTO departments (id_department, department_name)
		VALUES ($1, $2)
	`
		_, err = db.Exec(query, s.ID, s.Name)
		if err != nil {
			return err
		}
	}
	return nil
}

func InsertSpecialties(db *sql.DB, data []service.Specialty) error {

	for _, s := range data {

		query := `
			SELECT COUNT(*)
			FROM specialties
			WHERE id_specialty = $1
		`
		var count int
		err := db.QueryRow(query, s.ID).Scan(&count)
		if err != nil {
			return err
		}
		if count > 0 {
			continue
		}

		query = `
		INSERT INTO specialties (id_specialty, specialty_name, id_department)
		VALUES ($1, $2, $3)
	`
		_, err = db.Exec(query, s.ID, s.Name, s.DepartmentID)
		if err != nil {
			return err
		}
	}
	return nil
}

func InsertGroups(db *sql.DB, data []service.Group) error {

	for _, s := range data {

		query := `
			SELECT COUNT(*)
			FROM groups
			WHERE id_group = $1
		`
		var count int
		err := db.QueryRow(query, s.ID).Scan(&count)
		if err != nil {
			return err
		}
		if count > 0 {
			continue
		}

		query = `
		INSERT INTO groups (id_group, group_number, education_form, id_specialty, course_number)
		VALUES ($1, $2, $3, $4, $5)
	`
		_, err = db.Exec(query, s.ID, s.Number, s.EducationForm, s.SpecialtyID, s.CourseNumber)
		if err != nil {
			return err
		}
	}
	return nil
}

func InsertPersons(db *sql.DB, data []service.Person) error {

	for _, s := range data {

		query := `
			SELECT COUNT(*)
			FROM persons
			WHERE id_person = $1
		`
		var count int
		err := db.QueryRow(query, s.ID).Scan(&count)
		if err != nil {
			return err
		}
		if count > 0 {
			continue
		}

		query = `
		INSERT INTO persons (id_person, first_name, last_name, birth_date, passport, citizenship)
		VALUES ($1, $2, $3, $4, $5, $6)
	`
		_, err = db.Exec(query, s.ID, s.FirstName, s.LastName, s.BirthDate, s.Passport, s.Citizenship)
		if err != nil {
			return err
		}
	}
	return nil
}

func InsertMovement(db *sql.DB, data []service.StudentMovement) error {

	for _, s := range data {

		query := `
			SELECT COUNT(*)
			FROM student_movement
			WHERE id_move = $1
		`
		var count int
		err := db.QueryRow(query, s.ID).Scan(&count)
		if err != nil {
			return err
		}
		if count > 0 {
			continue
		}

		query = `
		INSERT INTO student_movement (id_move, id_student, status, transfer_date, course_number, id_group)
		VALUES ($1, $2, $3, $4, $5, $6)
	`
		_, err = db.Exec(query, s.ID, s.StudentID, s.Status, s.TransferDate, s.CourseNumber, s.GroupId)
		if err != nil {
			return err
		}
	}
	return nil
}

func InsertAddresses(db *sql.DB, data []service.Address) error {

	for _, s := range data {

		query := `
			SELECT COUNT(*)
			FROM addresses
			WHERE id_address = $1
		`
		var count int
		err := db.QueryRow(query, s.ID).Scan(&count)
		if err != nil {
			return err
		}
		if count > 0 {
			continue
		}

		query = `
		INSERT INTO addresses (id_address, address, id_person, address_type)
		VALUES ($1, $2, $3, $4)
	`
		_, err = db.Exec(query, s.ID, s.Address, s.PersonID, s.AddressType)
		if err != nil {
			return err
		}
	}
	return nil
}

func InsertPhones(db *sql.DB, data []service.Phone) error {

	for _, s := range data {

		query := `
			SELECT COUNT(*)
			FROM phones
			WHERE id_phone = $1
		`
		var count int
		err := db.QueryRow(query, s.ID).Scan(&count)
		if err != nil {
			return err
		}
		if count > 0 {
			continue
		}

		query = `
		INSERT INTO phones (id_phone, phone, id_person)
		VALUES ($1, $2, $3)
	`
		_, err = db.Exec(query, s.ID, s.Phone, s.PersonID)
		if err != nil {
			return err
		}
	}
	return nil
}

func CreatePersonTable(db *sql.DB) error {
	var query = []string{
		`
		CREATE TABLE IF NOT EXISTS persons (
			id_person INT PRIMARY KEY,
			first_name VARCHAR(100),
		    last_name VARCHAR(100),
			birth_date DATE,
		    passport VARCHAR(100),
		    citizenship VARCHAR(100)
		)
		`,
		`
		CREATE TABLE IF NOT EXISTS departments (
    		id_department INT PRIMARY KEY,
    		department_name VARCHAR(50) UNIQUE
)
		`,
		`
		CREATE TABLE IF NOT EXISTS specialties (
			id_specialty INT PRIMARY KEY,
			specialty_name VARCHAR(50) UNIQUE,
			id_department INT,
			FOREIGN KEY (id_department) REFERENCES departments(id_department)
		)
		`,
		`
		CREATE TABLE IF NOT EXISTS groups (
			id_group INT PRIMARY KEY,
			group_number VARCHAR(10) UNIQUE,
		    education_form VARCHAR(20),
			id_specialty INT,
		    course_number INT,
			FOREIGN KEY (id_specialty) REFERENCES specialties(id_specialty)
		)
		`,
		`
		CREATE TABLE IF NOT EXISTS students (
			id_student INT PRIMARY KEY,
			id_person INT,
			stud_number INT,
			id_group INT,
			FOREIGN KEY (id_person) REFERENCES persons(id_person),
		    FOREIGN KEY (id_group) REFERENCES groups(id_group)
		)
		`,
		`
		CREATE TABLE IF NOT EXISTS student_movement (
			id_move SERIAL PRIMARY KEY,
			id_student INT,
			id_group INT,
			course_number INT,
			status VARCHAR(20),
			transfer_date DATE,
			FOREIGN KEY (id_student) REFERENCES students(id_student),
		    FOREIGN KEY (id_group) REFERENCES groups(id_group)
		)
		`,
		`
		CREATE TABLE IF NOT EXISTS addresses (
			id_address INT PRIMARY KEY,
			address VARCHAR(100),
		    address_type VARCHAR(100),
		    id_person INT,
		    FOREIGN KEY (id_person) REFERENCES persons(id_person)
		)
		`,
		`
		CREATE TABLE IF NOT EXISTS phones (
			id_phone INT PRIMARY KEY, --Нужен ли тут PRIMARY KEY
			phone VARCHAR(100),
		    id_person INT,
		    FOREIGN KEY (id_person) REFERENCES persons(id_person)
		)
		`,
	}

	for i, q := range query {
		_, err := db.Exec(q)
		if err != nil {
			fmt.Println("err "+strconv.Itoa(i), err)
			return err
		}
	}

	return nil
}
