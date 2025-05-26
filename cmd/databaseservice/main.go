package main

import (
	"backend/internal/api/databaseservice/handler"
	"backend/internal/config"
	"backend/internal/core/action"
	"backend/internal/core/admin"
	"backend/internal/core/contactinfo"
	"backend/internal/core/credentials"
	"backend/internal/core/education"
	"backend/internal/core/familyinfo"
	"backend/internal/core/militarydetails"
	"backend/internal/core/person"
	"backend/internal/core/physicalinfo"
	"backend/internal/core/physicalstatus"
	"backend/internal/core/role"
	"backend/internal/core/skills"
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

	// Shared service
	actionService := action.NewService(dbInstance)

	// Core services
	roleService := role.NewService(dbInstance, actionService)
	adminService := admin.NewService(dbInstance, actionService, roleService)
	personService := person.NewService(dbInstance, actionService)
	credService := credentials.NewService(dbInstance, actionService)
	contactInfoService := contactinfo.NewService(dbInstance, actionService)
	educationService := education.NewService(dbInstance, actionService)
	familyInfoService := familyinfo.NewService(dbInstance, actionService)
	militaryDetailsService := militarydetails.NewService(dbInstance, actionService)
	physicalInfoService := physicalinfo.NewService(dbInstance, actionService)
	physicalStatusService := physicalstatus.NewService(dbInstance, actionService)
	skillsService := skills.NewService(dbInstance, actionService)
	// Seed database for testing
	if cfg.IsTest {
		if err := seeder.Seed(true, actionService); err != nil {
			logger.Error("Failed to seed database:", err)
			panic(err)
		}
		logger.Info("Database seeded successfully")
	}

	services := &handler.Services{
		AdminService:           adminService,
		ContactInfoService:     contactInfoService,
		CredentialsService:     credService,
		EducationService:       educationService,
		FamilyInfoService:      familyInfoService,
		MilitaryDetailsService: militaryDetailsService,
		PersonService:          personService,
		PhysicalInfoService:    physicalInfoService,
		PhysicalStatusService:  physicalStatusService,
		RoleService:            roleService,
		SkillsService:          skillsService,
	}

	h := handler.NewService(services)

	if err := router.InitRouter(cfg.AppHost, "8081", func(r *gin.Engine) {
		router.RegisterDatabaseRoutes(r, h)
	}); err != nil {
		logger.Error("Failed to start DatabaseService:", err)
		panic(err)
	}
}
