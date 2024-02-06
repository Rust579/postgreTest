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
			log.Fatal(err)
		}
		if count > 0 {
			continue
		}

		query = `
		INSERT INTO students (id_student, fio_student, passport, active_directions, academic_directions)
		VALUES ($1, $2, $3, $4, $5)
	`
		_, err = db.Exec(query, s.ID, s.Fio, s.Passport, s.ActiveDirections, s.AcademicDirections)
		if err != nil {
			log.Fatal(err)
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
			log.Fatal(err)
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
			log.Fatal(err)
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
			log.Fatal(err)
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
			log.Fatal(err)
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
			log.Fatal(err)
		}
		if count > 0 {
			continue
		}

		query = `
		INSERT INTO groups (id_group, group_number, id_specialty)
		VALUES ($1, $2, $3)
	`
		_, err = db.Exec(query, s.ID, s.Number, s.SpecialtyID)
		if err != nil {
			log.Fatal(err)
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
			log.Fatal(err)
		}
		if count > 0 {
			continue
		}

		query = `
		INSERT INTO persons (id_person, first_name, last_name, phone, age, birth_date, passport)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`
		_, err = db.Exec(query, s.ID, s.FirstName, s.LastName, s.Phone, s.Age, s.BirthDate, s.Passport)
		if err != nil {
			log.Fatal(err)
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
			log.Fatal(err)
		}
		if count > 0 {
			continue
		}

		query = `
		INSERT INTO student_movement (id_move, id_student, id_department, id_specialty, id_group, status, education_form, transfer_date)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	`
		_, err = db.Exec(query, s.ID, s.StudentID, s.DepartmentID, s.SpecialtyID, s.GroupID, s.Status, s.EducationForm, s.TransferDate)
		if err != nil {
			log.Fatal(err)
		}
	}
	return nil
}

func CreatePersonTable(db *sql.DB) error {
	var query = []string{
		`
		CREATE TABLE IF NOT EXISTS students (
			id_student INT PRIMARY KEY,
			fio_student VARCHAR(100),
		    passport VARCHAR(100),
		    active_directions INT,
		    academic_directions INT
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
			id_specialty INT,
			FOREIGN KEY (id_specialty) REFERENCES specialties(id_specialty)
		)
		`,
		`
		CREATE TABLE IF NOT EXISTS student_movement (
			id_move SERIAL PRIMARY KEY,
			id_student INT,
			id_department INT,
			id_specialty INT,
			id_group INT,
			status VARCHAR(20),
		    education_form VARCHAR(20),
			transfer_date DATE,
			FOREIGN KEY (id_student) REFERENCES students(id_student),
			FOREIGN KEY (id_department) REFERENCES departments(id_department),
			FOREIGN KEY (id_specialty) REFERENCES specialties(id_specialty),
			FOREIGN KEY (id_group) REFERENCES groups(id_group)
		)
		`,
		`
		CREATE TABLE IF NOT EXISTS persons (
			id_person INT PRIMARY KEY,
			first_name VARCHAR(100),
		    last_name VARCHAR(100),
		    phone VARCHAR(100),
		    age VARCHAR(100),
			birth_date DATE,
		    passport VARCHAR(100)
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
