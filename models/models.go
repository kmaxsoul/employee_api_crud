package models

import "time"

type Department struct {
	ID   int    `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}

type Employee struct {
	ID             int        `json:"id" db:"id"`
	FirstName      string     `json:"first_name" db:"first_name"`
	LastName       string     `json:"last_name" db:"last_name"`
	Age            int        `json:"age" db:"age"`
	Email          string     `json:"email" db:"email"`
	DepartmentID   int        `json:"department_id" db:"department_id"`
	DepartmentName string     `json:"department_name,omitempty" db:"department_name"`
	DeletedAt      *time.Time `json:"deleted_at,omitempty" db:"deleted_at"`
}

type Attendance struct {
	ID             int        `json:"id" db:"id"`
	EmployeeID     int        `json:"employee_id" db:"employee_id"`
	FirstName      string     `json:"first_name" db:"first_name"`
	LastName       string     `json:"last_name" db:"last_name"`
	DepartmentName string     `json:"department_name" db:"department_name"`
	CheckIn        time.Time  `json:"check_in" db:"check_in"`
	CheckOut       *time.Time `json:"check_out,omitempty" db:"check_out"`
}
