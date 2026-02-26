package repository

import (
	"context"
	"employee_crud/models"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

func CreateEmployee(pool *pgxpool.Pool, employee *models.Employee) (*models.Employee, error) {
	query := `
		WITH dept AS (
			SELECT name FROM crud_project.departments WHERE id = $5
		)
		INSERT INTO crud_project.employees (first_name, last_name, age, email, department_id, department_name)
		SELECT $1, $2, $3, $4, $5, dept.name FROM dept
		RETURNING id, first_name, last_name, age, email, department_id, department_name, deleted_at
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
		&created.DepartmentName,
		&created.DeletedAt,
	)
	if err != nil {
		return nil, err
	}

	return &created, nil
}

func GetAllEmployees(pool *pgxpool.Pool) ([]models.Employee, error) {
	query := `
		SELECT e.id, e.first_name, e.last_name, e.age, e.email, e.department_id, d.name, e.deleted_at
		FROM crud_project.employees e
		LEFT JOIN crud_project.departments d ON e.department_id = d.id
		WHERE e.deleted_at IS NULL
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
			&emp.DepartmentName,
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
		WITH dept AS (
			SELECT name FROM crud_project.departments WHERE id = $5
		)
		UPDATE crud_project.employees
		SET first_name = $1, last_name = $2, age = $3, email = $4, department_id = $5,
		    department_name = (SELECT name FROM dept)
		WHERE id = $6 AND deleted_at IS NULL
		RETURNING id, first_name, last_name, age, email, department_id, department_name, deleted_at
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
		&updated.DepartmentName,
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
