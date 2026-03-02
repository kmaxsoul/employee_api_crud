package repository

import (
	"context"
	"employee_crud/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

func CreateAttendance(pool *pgxpool.Pool, employeeID int) (*models.Attendance, error) {
	query := `
		INSERT INTO crud_project.attendance (employee_id, first_name, last_name, department_name, check_in)
		SELECT e.id, e.first_name, e.last_name, e.department_name, NOW()
		FROM crud_project.employees e
		WHERE e.id = $1 AND e.deleted_at IS NULL
		RETURNING id, employee_id, first_name, last_name, department_name, check_in, check_out
	`

	var a models.Attendance
	err := pool.QueryRow(context.Background(), query, employeeID).Scan(
		&a.ID,
		&a.EmployeeID,
		&a.FirstName,
		&a.LastName,
		&a.DepartmentName,
		&a.CheckIn,
		&a.CheckOut,
	)
	if err != nil {
		return nil, err
	}

	return &a, nil
}

func CheckOutAttendance(pool *pgxpool.Pool, employeeID int) (*models.Attendance, error) {
	query := `
		UPDATE crud_project.attendance
		SET check_out = NOW()
		WHERE id = (
			SELECT id FROM crud_project.attendance
			WHERE employee_id = $1 AND check_out IS NULL
			ORDER BY check_in DESC
			LIMIT 1
		)
		RETURNING id, employee_id, first_name, last_name, department_name, check_in, check_out
	`

	var a models.Attendance
	err := pool.QueryRow(context.Background(), query, employeeID).Scan(
		&a.ID,
		&a.EmployeeID,
		&a.FirstName,
		&a.LastName,
		&a.DepartmentName,
		&a.CheckIn,
		&a.CheckOut,
	)
	if err != nil {
		return nil, err
	}

	return &a, nil
}

func GetAllAttendances(pool *pgxpool.Pool) ([]models.Attendance, error) {
	query := `
		SELECT id, employee_id, first_name, last_name, department_name, check_in, check_out
		FROM crud_project.attendance
		ORDER BY check_in DESC
	`

	rows, err := pool.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var attendances []models.Attendance
	for rows.Next() {
		var a models.Attendance
		err := rows.Scan(
			&a.ID,
			&a.EmployeeID,
			&a.FirstName,
			&a.LastName,
			&a.DepartmentName,
			&a.CheckIn,
			&a.CheckOut,
		)
		if err != nil {
			return nil, err
		}
		attendances = append(attendances, a)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return attendances, nil
}

func GetAttendancesByEmployee(pool *pgxpool.Pool, employeeID int) ([]models.Attendance, error) {
	query := `
		SELECT id, employee_id, first_name, last_name, department_name, check_in, check_out
		FROM crud_project.attendance
		WHERE employee_id = $1
		ORDER BY check_in DESC
	`

	rows, err := pool.Query(context.Background(), query, employeeID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var attendances []models.Attendance
	for rows.Next() {
		var a models.Attendance
		err := rows.Scan(
			&a.ID,
			&a.EmployeeID,
			&a.FirstName,
			&a.LastName,
			&a.DepartmentName,
			&a.CheckIn,
			&a.CheckOut,
		)
		if err != nil {
			return nil, err
		}
		attendances = append(attendances, a)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return attendances, nil
}
