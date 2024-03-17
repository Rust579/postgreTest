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
		ID:         1,
		PersonID:   1,
		StudNumber: 1,
		GroupID:    1,
	},
	{
		ID:         2,
		PersonID:   1,
		StudNumber: 2,
		GroupID:    2,
	},
	{
		ID:         3,
		PersonID:   3,
		StudNumber: 3,
		GroupID:    3,
	},
	{
		ID:         4,
		PersonID:   4,
		StudNumber: 4,
		GroupID:    2,
	},
	{
		ID:         5,
		PersonID:   5,
		StudNumber: 5,
		GroupID:    3,
	},
	{
		ID:         6,
		PersonID:   6,
		StudNumber: 6,
		GroupID:    5,
	},
	{
		ID:         7,
		PersonID:   7,
		StudNumber: 7,
		GroupID:    2,
	},
	{
		ID:         8,
		PersonID:   4,
		StudNumber: 8,
		GroupID:    1,
	},
	{
		ID:         9,
		PersonID:   7,
		StudNumber: 9,
		GroupID:    3,
	},
	{
		ID:         10,
		PersonID:   2,
		StudNumber: 10,
		GroupID:    5,
	},
	{
		ID:         11,
		PersonID:   1,
		StudNumber: 1,
		GroupID:    6,
	},
	{
		ID:         12,
		PersonID:   2,
		StudNumber: 2,
		GroupID:    6,
	},
	{
		ID:         13,
		PersonID:   1,
		StudNumber: 3,
		GroupID:    7,
	},
	{
		ID:         14,
		PersonID:   3,
		StudNumber: 4,
		GroupID:    7,
	},
	{
		ID:         15,
		PersonID:   2,
		StudNumber: 5,
		GroupID:    8,
	},
}

var Persons = []Person{
	{
		ID:          1,
		FirstName:   "Иван",
		LastName:    "Иванов",
		BirthDate:   "1990-01-01",
		Passport:    "123456 0001",
		Citizenship: "РФ",
	},
	{
		ID:          2,
		FirstName:   "Руслан",
		LastName:    "Козодоев",
		BirthDate:   "1995-05-02",
		Passport:    "123456 0002",
		Citizenship: "РФ",
	},
	{
		ID:          3,
		FirstName:   "Петр",
		LastName:    "Листов",
		BirthDate:   "1985-12-10",
		Passport:    "123456 0003",
		Citizenship: "РФ",
	},
	{
		ID:          4,
		FirstName:   "Авраам",
		LastName:    "Зеленый",
		BirthDate:   "2000-01-01",
		Passport:    "123456 0004",
		Citizenship: "РФ",
	},
	{
		ID:          5,
		FirstName:   "Галерей",
		LastName:    "Третьяков",
		BirthDate:   "1999-12-30",
		Passport:    "123456 0005",
		Citizenship: "РФ",
	},
	{
		ID:          6,
		FirstName:   "Салават",
		LastName:    "Уфимский",
		BirthDate:   "1990-12-10",
		Passport:    "123456 0006",
		Citizenship: "РФ",
	},
	{
		ID:          7,
		FirstName:   "Небов",
		LastName:    "Чистый",
		BirthDate:   "2005-10-02",
		Passport:    "123456 0007",
		Citizenship: "РФ",
	},
	{
		ID:          8,
		FirstName:   "Пикчер",
		LastName:    "Парамаунтов",
		BirthDate:   "2001-11-11",
		Passport:    "0000 0001",
		Citizenship: "Japan",
	},
	{
		ID:          9,
		FirstName:   "Тестер",
		LastName:    "Тестовый",
		BirthDate:   "1996-05-05",
		Passport:    "0000 0002",
		Citizenship: "USA",
	},
	{
		ID:          10,
		FirstName:   "Бетон",
		LastName:    "Железный",
		BirthDate:   "1980-03-03",
		Passport:    "0000 0003",
		Citizenship: "France",
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
	{
		ID:   4,
		Name: "Математика",
	},
	{
		ID:   5,
		Name: "Химия",
	},
	{
		ID:   6,
		Name: "Философия",
	},
	{
		ID:   7,
		Name: "Материаловедение",
	},
	{
		ID:   8,
		Name: "Сопромат",
	},
}

var Specialties = []Specialty{
	{
		ID:           1,
		Name:         "АВИО",
		DepartmentID: 1,
	},
	{
		ID:           2,
		Name:         "ИТ",
		DepartmentID: 2,
	},
	{
		ID:           3,
		Name:         "МЕХ",
		DepartmentID: 3,
	},
	{
		ID:           4,
		Name:         "МАТЕМ",
		DepartmentID: 4,
	},
	{
		ID:           5,
		Name:         "ХИМ",
		DepartmentID: 5,
	},
	{
		ID:           6,
		Name:         "ФИЛ",
		DepartmentID: 6,
	},
	{
		ID:           7,
		Name:         "МАТЕР",
		DepartmentID: 7,
	},
	{
		ID:           8,
		Name:         "СОПР",
		DepartmentID: 8,
	},
}

var Groups = []Group{
	{
		ID:            1,
		Number:        "АВИО-1",
		EducationForm: ConstEducFormFullTime,
		SpecialtyID:   1,
		CourseNumber:  1,
	},
	{
		ID:            2,
		Number:        "ИТ-1",
		EducationForm: ConstEducFormFullTime,
		SpecialtyID:   2,
		CourseNumber:  1,
	},
	{
		ID:            3,
		Number:        "МЕХ-1",
		EducationForm: ConstEducFormFullTime,
		SpecialtyID:   3,
		CourseNumber:  1,
	},
	{
		ID:            4,
		Number:        "МАТЕМ-2",
		EducationForm: ConstEducFormFullTime,
		SpecialtyID:   4,
		CourseNumber:  2,
	},
	{
		ID:            5,
		Number:        "ХИМ-2",
		EducationForm: ConstEducFormPartTime,
		SpecialtyID:   5,
		CourseNumber:  2,
	},
	{
		ID:            6,
		Number:        "ФИЛ-2",
		EducationForm: ConstEducFormFullTime,
		SpecialtyID:   6,
		CourseNumber:  2,
	},
	{
		ID:            7,
		Number:        "МАТЕР-2",
		EducationForm: ConstEducFormFullTime,
		SpecialtyID:   7,
		CourseNumber:  2,
	},
	{
		ID:            8,
		Number:        "СОПР-2",
		EducationForm: ConstEducFormPartTime,
		SpecialtyID:   8,
		CourseNumber:  2,
	},
}

var StudentMovements = []StudentMovement{
	{
		ID:           1,
		StudentID:    1,
		GroupId:      1,
		CourseNumber: 1,
		Status:       ConstStudentStatusActive,
		TransferDate: "2023-01-01",
	},
	{
		ID:           2,
		StudentID:    2,
		GroupId:      2,
		CourseNumber: 1,
		Status:       ConstStudentStatusActive,
		TransferDate: "2023-01-01",
	},
	{
		ID:           3,
		StudentID:    3,
		GroupId:      3,
		CourseNumber: 1,
		Status:       ConstStudentStatusActive,
		TransferDate: "2024-05-06",
	},
	{
		ID:           4,
		StudentID:    4,
		GroupId:      4,
		CourseNumber: 2,
		Status:       ConstStudentStatusDismissed,
		TransferDate: "2024-01-01",
	},
	{
		ID:           5,
		StudentID:    5,
		GroupId:      5,
		CourseNumber: 2,
		Status:       ConstStudentStatusActive,
		TransferDate: "2023-02-03",
	},
	{
		ID:           6,
		StudentID:    6,
		GroupId:      6,
		CourseNumber: 2,
		Status:       ConstStudentStatusActive,
		TransferDate: "2023-04-04",
	},
	{
		ID:           7,
		StudentID:    7,
		GroupId:      7,
		CourseNumber: 2,
		Status:       ConstStudentStatusAcademic,
		TransferDate: "2023-09-06",
	},
	{
		ID:           8,
		StudentID:    8,
		GroupId:      8,
		CourseNumber: 2,
		Status:       ConstStudentStatusActive,
		TransferDate: "2023-06-06",
	},
	{
		ID:           9,
		StudentID:    9,
		GroupId:      2,
		CourseNumber: 3,
		Status:       ConstStudentStatusAcademic,
		TransferDate: "2023-02-10",
	},
	{
		ID:           10,
		StudentID:    10,
		GroupId:      8,
		CourseNumber: 3,
		Status:       ConstStudentStatusActive,
		TransferDate: "2024-02-03",
	},
	{
		ID:           11,
		StudentID:    1,
		GroupId:      8,
		CourseNumber: 3,
		Status:       ConstStudentStatusActive,
		TransferDate: "2023-01-01",
	},
	{
		ID:           12,
		StudentID:    10,
		GroupId:      3,
		CourseNumber: 2,
		Status:       ConstStudentStatusActive,
		TransferDate: "2023-01-01",
	},
	{
		ID:           13,
		StudentID:    13,
		GroupId:      5,
		CourseNumber: 1,
		Status:       ConstStudentStatusActive,
		TransferDate: "2024-05-06",
	},
	{
		ID:           14,
		StudentID:    12,
		GroupId:      7,
		CourseNumber: 3,
		Status:       ConstStudentStatusDismissed,
		TransferDate: "2024-01-01",
	},
	{
		ID:           15,
		StudentID:    11,
		GroupId:      1,
		CourseNumber: 1,
		Status:       ConstStudentStatusActive,
		TransferDate: "2023-02-03",
	},
}

var Addresses = []Address{
	{
		ID:          1,
		Address:     "Адрес РФ 1-1",
		AddressType: "Постоянная прописка",
		PersonID:    1,
	},
	{
		ID:          2,
		Address:     "Адрес РФ 1-2",
		AddressType: "Постоянная прописка",
		PersonID:    1,
	},
	{
		ID:          3,
		Address:     "Адрес РФ 2",
		AddressType: "Постоянная прописка",
		PersonID:    2,
	},
	{
		ID:          4,
		Address:     "Адрес РФ 3",
		AddressType: "Временная прописка",
		PersonID:    3,
	},
	{
		ID:          5,
		Address:     "Адрес РФ 4",
		AddressType: "Временная прописка",
		PersonID:    4,
	},
	{
		ID:          6,
		Address:     "Адрес РФ 5",
		AddressType: "Временная прописка",
		PersonID:    5,
	},
	{
		ID:          7,
		Address:     "Адрес РФ 6",
		AddressType: "Временная прописка",
		PersonID:    6,
	},
	{
		ID:          8,
		Address:     "Адрес РФ 7-1",
		AddressType: "Место жительства",
		PersonID:    7,
	},
	{
		ID:          9,
		Address:     "Адрес РФ 7-2",
		AddressType: "Место жительства",
		PersonID:    7,
	},
	{
		ID:          10,
		Address:     "Адрес РФ 7-3",
		AddressType: "Место жительства",
		PersonID:    8,
	},
	{
		ID:          11,
		Address:     "Адрес иностранный 1",
		AddressType: "Место жительства",
		PersonID:    9,
	},
	{
		ID:          12,
		Address:     "Адрес иностранный 2",
		AddressType: "Место жительства",
		PersonID:    10,
	},
	{
		ID:          13,
		Address:     "Адрес иностранный 3-1",
		AddressType: "Временная прописка",
		PersonID:    9,
	},
	{
		ID:          14,
		Address:     "Адрес иностранный 3-2",
		AddressType: "Временная прописка",
		PersonID:    9,
	},
	{
		ID:          15,
		Address:     "Адрес иностранный 3-3",
		AddressType: "Временная прописка",
		PersonID:    10,
	},
}

var Phones = []Phone{
	{
		ID:       1,
		Phone:    "89990000001",
		PersonID: 1,
	},
	{
		ID:       2,
		Phone:    "89990000002",
		PersonID: 2,
	},
	{
		ID:       3,
		Phone:    "89990000003",
		PersonID: 3,
	},
	{
		ID:       4,
		Phone:    "89990000004",
		PersonID: 3,
	},
	{
		ID:       5,
		Phone:    "89990000005",
		PersonID: 3,
	},
	{
		ID:       6,
		Phone:    "89990000006",
		PersonID: 4,
	},
	{
		ID:       7,
		Phone:    "89990000007",
		PersonID: 5,
	},
	{
		ID:       8,
		Phone:    "89990000008",
		PersonID: 6,
	},
	{
		ID:       9,
		Phone:    "89990000009",
		PersonID: 7,
	},
	{
		ID:       10,
		Phone:    "89990000010",
		PersonID: 7,
	},
	{
		ID:       11,
		Phone:    "55500001",
		PersonID: 8,
	},
	{
		ID:       12,
		Phone:    "55500002",
		PersonID: 9,
	},
	{
		ID:       13,
		Phone:    "55500003",
		PersonID: 10,
	},
	{
		ID:       14,
		Phone:    "55500004",
		PersonID: 9,
	},
	{
		ID:       15,
		Phone:    "55500005",
		PersonID: 10,
	},
}
