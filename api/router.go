package api

import (
	"employee_crud/config"
	"employee_crud/handlers"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func SetupRouter(pool *pgxpool.Pool, cfg *config.Config) *gin.Engine {
	router := gin.Default()

	employeeGroup := router.Group("/employees")
	{
		employeeGroup.POST("", handlers.CreateEmployee(pool))
		employeeGroup.GET("", handlers.GetAllEmployees(pool))
		employeeGroup.PUT("/:id", handlers.UpdateEmployee(pool))
		employeeGroup.DELETE("/:id", handlers.DeleteEmployee(pool))
	}

	attendanceGroup := router.Group("/attendance")
	{
		attendanceGroup.POST("", handlers.CreateAttendance(pool))
		attendanceGroup.GET("", handlers.GetAllAttendances(pool))
	}

	return router
}
