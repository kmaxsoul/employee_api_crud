package repository

import (
	"context"
	"employee_crud/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

func CreateEmployee(pool *pgxpool.Pool, employee *models.Employee) (*models.Employee, error) {
	query := `
		INSERT INTO crud_project.employees (first_name, last_name, age, email, department_id)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, first_name, last_name, age, email, department_id, deleted_at
	`

	var created models.Employee
	err := pool.QueryRow(
		context.Background(),
		query,
		employee.FirstName,
		employee.LastName,
		employee.Age,
		employee.Email,
		employee.DepartmentID,
	).Scan(
		&created.ID,
		&created.FirstName,
		&created.LastName,
		&created.Age,
		&created.Email,
		&created.DepartmentID,
		&created.DeletedAt,
	)
	if err != nil {
		return nil, err
	}

	return &created, nil
}
