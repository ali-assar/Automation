package handler

import (
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
)

// Services holds all core services used to initialize the handler.Service.
type Services struct {
	AdminService           *admin.Service
	ContactInfoService     *contactinfo.Service
	CredentialsService     *credentials.Service
	EducationService       *education.Service
	FamilyInfoService      *familyinfo.Service
	MilitaryDetailsService *militarydetails.Service
	PersonService          *person.Service
	PhysicalInfoService    *physicalinfo.Service
	PhysicalStatusService  *physicalstatus.Service
	RoleService            *role.Service
	SkillsService          *skills.Service
}

// Service aggregates all core services for the databaseservice handlers.
type Service struct {
	adminService           *admin.Service
	contactInfoService     *contactinfo.Service
	credentialsService     *credentials.Service
	educationService       *education.Service
	familyInfoService      *familyinfo.Service
	militaryDetailsService *militarydetails.Service
	personService          *person.Service
	physicalInfoService    *physicalinfo.Service
	physicalStatusService  *physicalstatus.Service
	roleService            *role.Service
	skillsService          *skills.Service
}

// NewService creates a new Service instance with the provided services.
func NewService(s *Services) *Service {
	return &Service{
		adminService:           s.AdminService,
		contactInfoService:     s.ContactInfoService,
		credentialsService:     s.CredentialsService,
		educationService:       s.EducationService,
		familyInfoService:      s.FamilyInfoService,
		militaryDetailsService: s.MilitaryDetailsService,
		personService:          s.PersonService,
		physicalInfoService:    s.PhysicalInfoService,
		physicalStatusService:  s.PhysicalStatusService,
		roleService:            s.RoleService,
		skillsService:          s.SkillsService,
	}
}
