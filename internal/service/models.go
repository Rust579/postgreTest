package service

type Student struct {
	ID                 int
	Fio                string
	Passport           string
	ActiveDirections   int
	AcademicDirections int
}

type Department struct {
	ID   int
	Name string
}

type Specialty struct {
	ID           int
	Name         string
	DepartmentID int
}

type Group struct {
	ID          int
	Number      string
	SpecialtyID int
}

type StudentMovement struct {
	ID            int
	StudentID     int
	DepartmentID  int
	SpecialtyID   int
	GroupID       int
	Status        string
	EducationForm string
	TransferDate  string
}

type Person struct {
	ID        int
	FirstName string
	LastName  string
	Phone     string
	Age       int
	BirthDate string
	Passport  string
}
