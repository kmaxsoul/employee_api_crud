package handlers

import (
	"employee_crud/models"
	"employee_crud/repository"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

type CreateEmployeeInput struct {
	FirstName    string `json:"first_name" binding:"required"`
	LastName     string `json:"last_name" binding:"required"`
	Age          int    `json:"age" binding:"required"`
	Email        string `json:"email" binding:"required,email"`
	DepartmentID int    `json:"department_id" binding:"required"`
}

func CreateEmployee(pool *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input CreateEmployeeInput
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		emplpoyee, err := repository.CreateEmployee(pool, &models.Employee{
			FirstName:    input.FirstName,
			LastName:     input.LastName,
			Age:          input.Age,
			Email:        input.Email,
			DepartmentID: input.DepartmentID,
		})
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(201, emplpoyee)
	}

}

func GetAllEmployees(pool *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {
		employees, err := repository.GetAllEmployees(pool)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, employees)
	}
}

func UpdateEmployee(pool *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input CreateEmployeeInput
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		employeeID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(400, gin.H{"error": "invalid employee id"})
			return
		}
		updatedEmployee, err := repository.UpdateEmployee(pool, &models.Employee{
			ID:           employeeID,
			FirstName:    input.FirstName,
			LastName:     input.LastName,
			Age:          input.Age,
			Email:        input.Email,
			DepartmentID: input.DepartmentID,
		})
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, updatedEmployee)
	}
}

func DeleteEmployee(pool *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {
		employeeID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(400, gin.H{"error": "invalid employee id"})
			return
		}

		err = repository.DeleteEmployee(pool, employeeID)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.Status(204)
	}
}
