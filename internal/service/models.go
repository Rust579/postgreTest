package service

type Student struct {
	ID         int
	PersonID   int
	StudNumber int
	GroupID    int
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
	ID            int
	Number        string
	EducationForm string
	SpecialtyID   int
	CourseNumber  int
}

type StudentMovement struct {
	ID           int
	StudentID    int
	CourseNumber int
	Status       string
	TransferDate string
	GroupId      int
}

type Person struct {
	ID          int
	FirstName   string
	LastName    string
	BirthDate   string
	Passport    string
	Citizenship string
}

type Address struct {
	ID          int
	Address     string
	PersonID    int
	AddressType string
}

type Phone struct {
	ID       int
	Phone    string
	PersonID int
}

type Email struct {
	ID       int
	Email    string
	PersonID int
}
