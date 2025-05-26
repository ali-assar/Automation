package main

import (
	"backend/internal/api/databaseservice"
	"backend/internal/config"
	"backend/internal/core/action"
	"backend/internal/core/admin"
	"backend/internal/core/credentials"
	"backend/internal/core/person"
	"backend/internal/core/role"
	"backend/internal/db"
	"backend/internal/logger"
	"backend/internal/seeder" // Add seeder package
	"backend/pkg/services/router"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.Load()
	logger.Init(cfg.LogLevel, cfg.LogFormat)
	if err := db.Init(); err != nil {
		logger.Error("Failed to initialize database:", err)
		panic(err)
	}
	dbInstance := db.GetDB()

	actionService := action.NewService(dbInstance)
	roleService := role.NewService(dbInstance, actionService)

	adminService := admin.NewService(dbInstance, actionService, roleService)
	personService := person.NewService(dbInstance, actionService)
	credService := credentials.NewService(dbInstance, actionService)

	// Seed database for testing
	if cfg.IsTest {
		if err := seeder.Seed(true, actionService); err != nil {
			logger.Error("Failed to seed database:", err)
			panic(err)
		}
		logger.Info("Database seeded successfully")
	}

	handler := databaseservice.NewHandler(adminService, personService, credService)
	if err := router.InitRouter(cfg.AppHost, "8081", func(r *gin.Engine) {
		router.RegisterDatabaseRoutes(r, handler)
	}); err != nil {
		logger.Error("Failed to start DatabaseService:", err)
		panic(err)
	}
}
