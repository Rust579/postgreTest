package service

const (
	ConstStudentStatusActive    = "active"
	ConstStudentStatusDismissed = "dismissed"
	ConstStudentStatusAcademic  = "academic"

	ConstEducFormFullTime = "Очная"
	ConstEducFormPartTime = "Заочная"
)

var Students = []Student{
	{
		ID:                 1,
		Fio:                "Иван Иванов",
		Passport:           "123456 0001",
		ActiveDirections:   1,
		AcademicDirections: 0,
	},
	{
		ID:                 2,
		Fio:                "Козодоев Руслан",
		Passport:           "123456 0002",
		ActiveDirections:   2,
		AcademicDirections: 0,
	},
	{
		ID:                 3,
		Fio:                "Листов Петр",
		Passport:           "123456 0003",
		ActiveDirections:   1,
		AcademicDirections: 1,
	},
}

var Departments = []Department{
	{
		ID:   1,
		Name: "Авионика",
	},
	{
		ID:   2,
		Name: "Информатика",
	},
	{
		ID:   3,
		Name: "Механика",
	},
}

var Specialties = []Specialty{
	{
		ID:           1,
		Name:         "СЭМС",
		DepartmentID: 1,
	},
	{
		ID:           2,
		Name:         "ИИТ",
		DepartmentID: 2,
	},
	{
		ID:           3,
		Name:         "РТ",
		DepartmentID: 3,
	},
}

var Groups = []Group{
	{
		ID:          1,
		Number:      "СЭМС-1",
		SpecialtyID: 1,
	},
	{
		ID:          2,
		Number:      "ИИТ-1",
		SpecialtyID: 2,
	},
	{
		ID:          3,
		Number:      "РТ-1",
		SpecialtyID: 3,
	},
}

var StudentMovements = []StudentMovement{
	{
		ID:            1,
		StudentID:     1,
		DepartmentID:  1,
		SpecialtyID:   1,
		GroupID:       1,
		Status:        ConstStudentStatusActive,
		EducationForm: ConstEducFormFullTime,
		TransferDate:  "2023-01-01",
	},
	{
		ID:            2,
		StudentID:     2,
		DepartmentID:  2,
		SpecialtyID:   2,
		GroupID:       2,
		Status:        ConstStudentStatusAcademic,
		EducationForm: ConstEducFormFullTime,
		TransferDate:  "2023-01-01",
	},
	{
		ID:            3,
		StudentID:     3,
		DepartmentID:  3,
		SpecialtyID:   3,
		GroupID:       3,
		Status:        ConstStudentStatusDismissed,
		EducationForm: ConstEducFormPartTime,
		TransferDate:  "2024-05-06",
	},
}

var Persons = []Person{
	{
		ID:        1,
		FirstName: "Иван",
		LastName:  "Иванов",
		Phone:     "89170000001",
		Age:       34,
		BirthDate: "1990-01-01",
		Passport:  "123456 0001",
	},
	{
		ID:        2,
		FirstName: "Руслан",
		LastName:  "Козодоев",
		Phone:     "89170000002",
		Age:       29,
		BirthDate: "1995-05-02",
		Passport:  "123456 0002",
	},
	{
		ID:        3,
		FirstName: "Петр",
		LastName:  "Листов",
		Phone:     "89170000002",
		Age:       39,
		BirthDate: "1985-12-10",
		Passport:  "123456 0003",
	},
}
