package main

import (
	"backend/internal/api/api"
	"backend/internal/config"
	"backend/internal/core/action"
	"backend/internal/core/admin"
	"backend/internal/core/bloodgroup"
	"backend/internal/core/contactinfo"
	"backend/internal/core/credentials"
	"backend/internal/core/education"
	educationlevel "backend/internal/core/educationLevel"
	"backend/internal/core/familyinfo"
	"backend/internal/core/gender"
	"backend/internal/core/hospitaldispatch"
	"backend/internal/core/medicalprofile"
	"backend/internal/core/medicines"
	"backend/internal/core/militarydetails"
	"backend/internal/core/person"
	"backend/internal/core/persontype"
	"backend/internal/core/physicalinfo"
	"backend/internal/core/physicalstatus"
	prescriptions "backend/internal/core/prescription"
	"backend/internal/core/psychologicalstatus"
	"backend/internal/core/rank"
	"backend/internal/core/religion"
	"backend/internal/core/role"
	"backend/internal/core/skills"
	"backend/internal/db"
	"backend/internal/logger"
	"backend/internal/seeder"
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
	educationLevelService := educationlevel.NewService(dbInstance, actionService)
	familyInfoService := familyinfo.NewService(dbInstance, actionService)
	militaryDetailsService := militarydetails.NewService(dbInstance, actionService)
	physicalInfoService := physicalinfo.NewService(dbInstance, actionService)
	physicalStatusService := physicalstatus.NewService(dbInstance, actionService)
	skillsService := skills.NewService(dbInstance, actionService)
	bloodGroupService := bloodgroup.NewService(dbInstance, actionService)
	rankService := rank.NewService(dbInstance, actionService)
	religionService := religion.NewService(dbInstance, actionService)
	personTypeService := persontype.NewService(dbInstance, actionService)
	genderService := gender.NewService(dbInstance, actionService)

	medicalProfile := medicalprofile.NewService(dbInstance, actionService)
	hospitalDispatch := hospitaldispatch.NewService(dbInstance, actionService)
	prescription := prescriptions.NewService(dbInstance, actionService)
	psychologicalStatus := psychologicalstatus.NewService(dbInstance, actionService)
	medicine := medicines.NewService(dbInstance, actionService)

	// Seed database for testing
	if cfg.IsTest {
		if err := seeder.Seed(true, actionService); err != nil {
			logger.Error("Failed to seed database:", err)
			panic(err)
		}
		logger.Info("Database seeded successfully")
	}

	// Initialize CoreServices
	coreServices := &api.CoreServices{
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
		BloodGroupService:      bloodGroupService,
		RankService:            rankService,
		ReligionService:        religionService,
		PersonTypeService:      personTypeService,
		EducationLevelService:  educationLevelService,
		GenderService:          genderService,
		MedicalProfile:         medicalProfile,
		HospitalDispatch:       hospitalDispatch,
		Prescription:           prescription,
		PsychologicalStatus:    psychologicalStatus,
		Medicine:               medicine,
	}
	gin.SetMode(gin.DebugMode)
	// Create HandlerService
	h := api.NewHandlerService(coreServices)
	// Initialize router
	if err := router.InitRouter(cfg.AppHost, cfg.AppPort, func(r *gin.Engine) {
		router.RegisterRoutes(r, h)
	}); err != nil {
		logger.Error("Failed to start API service:", err)
		panic(err)
	}
}
