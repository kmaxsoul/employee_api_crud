package handlers

import (
	"employee_crud/repository"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type AttendanceInput struct {
	EmployeeID int  `json:"employee_id" binding:"required"`
	IsCheckIn  bool `json:"is_check_in"`
}

func CreateAttendance(pool *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input AttendanceInput
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if input.IsCheckIn {
			// Check in → create a new attendance record, server captures the time
			attendance, err := repository.CreateAttendance(pool, input.EmployeeID)
			if err != nil {
				if errors.Is(err, pgx.ErrNoRows) {
					c.JSON(http.StatusNotFound, gin.H{"error": "employee not found or has been deleted"})
					return
				}
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusCreated, attendance)
		} else {
			// Check out → close the latest open attendance record
			attendance, err := repository.CheckOutAttendance(pool, input.EmployeeID)
			if err != nil {
				if errors.Is(err, pgx.ErrNoRows) {
					c.JSON(http.StatusNotFound, gin.H{"error": "no open check-in found for this employee"})
					return
				}
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusOK, attendance)
		}
	}
}

func GetAllAttendances(pool *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Optional filter by employee_id query param
		employeeIDStr := c.Query("employee_id")

		if employeeIDStr != "" {
			employeeID, err := strconv.Atoi(employeeIDStr)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "invalid employee_id"})
				return
			}
			attendances, err := repository.GetAttendancesByEmployee(pool, employeeID)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusOK, attendances)
			return
		}

		attendances, err := repository.GetAllAttendances(pool)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, attendances)
	}
}
