package main

import (
	"employee_crud/api"
	"employee_crud/config"
	"employee_crud/database"
	"log"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Failed to load configuration:", err)
	}

	pool, err := database.ConnectPostgres(cfg.DatabaseURL)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer pool.Close()

	router := api.SetupRouter(pool, cfg)
	router.Run(":" + cfg.Port)
}
