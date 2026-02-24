package models

import "time"

type Department struct {
	ID   int    `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}

type Employee struct {
	ID           int        `json:"id" db:"id"`
	FirstName    string     `json:"first_name" db:"first_name"`
	LastName     string     `json:"last_name" db:"last_name"`
	Age          int        `json:"age" db:"age"`
	Email        string     `json:"email" db:"email"`
	DepartmentID int        `json:"department_id" db:"department_id"`
	DeletedAt    *time.Time `json:"deleted_at,omitempty" db:"deleted_at"`
}
