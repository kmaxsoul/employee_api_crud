package repository

import (
	"context"
	"employee_crud/models"
	"fmt"

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

func GetAllEmployees(pool *pgxpool.Pool) ([]models.Employee, error) {
	query := `
		SELECT id, first_name, last_name, age, email, department_id, deleted_at
		FROM crud_project.employees
		WHERE deleted_at IS NULL
	`
	rows, err := pool.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var employees []models.Employee
	for rows.Next() {
		var emp models.Employee
		err := rows.Scan(
			&emp.ID,
			&emp.FirstName,
			&emp.LastName,
			&emp.Age,
			&emp.Email,
			&emp.DepartmentID,
			&emp.DeletedAt,
		)
		if err != nil {
			return nil, err
		}
		employees = append(employees, emp)
	}

	return employees, nil

}

func UpdateEmployee(pool *pgxpool.Pool, employee *models.Employee) (*models.Employee, error) {
	query := `
		UPDATE crud_project.employees
		SET first_name = $1, last_name = $2, age = $3, email = $4, department_id = $5
		WHERE id = $6 AND deleted_at IS NULL
		RETURNING id, first_name, last_name, age, email, department_id, deleted_at
	`

	var updated models.Employee
	err := pool.QueryRow(
		context.Background(),
		query,
		employee.FirstName,
		employee.LastName,
		employee.Age,
		employee.Email,
		employee.DepartmentID,
		employee.ID,
	).Scan(
		&updated.ID,
		&updated.FirstName,
		&updated.LastName,
		&updated.Age,
		&updated.Email,
		&updated.DepartmentID,
		&updated.DeletedAt,
	)
	if err != nil {
		return nil, err
	}

	return &updated, nil
}

func DeleteEmployee(pool *pgxpool.Pool, id int) error {
	query := `
		UPDATE crud_project.employees
		SET deleted_at = NOW()
		WHERE id = $1 AND deleted_at IS NULL
	`

	result, err := pool.Exec(context.Background(), query, id)
	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return fmt.Errorf("employee not found")
	}

	return nil
}
