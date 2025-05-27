package api

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

// CoreServices holds pointers to all core services used to initialize HandlerService.
type CoreServices struct {
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

// HandlerService aggregates core services for API handlers.
type HandlerService struct {
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

// NewHandlerService creates a new HandlerService instance with the provided core services.
func NewHandlerService(cs *CoreServices) *HandlerService {
	return &HandlerService{
		AdminService:           cs.AdminService,
		ContactInfoService:     cs.ContactInfoService,
		CredentialsService:     cs.CredentialsService,
		EducationService:       cs.EducationService,
		FamilyInfoService:      cs.FamilyInfoService,
		MilitaryDetailsService: cs.MilitaryDetailsService,
		PersonService:          cs.PersonService,
		PhysicalInfoService:    cs.PhysicalInfoService,
		PhysicalStatusService:  cs.PhysicalStatusService,
		RoleService:            cs.RoleService,
		SkillsService:          cs.SkillsService,
	}
}
