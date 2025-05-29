package api

import (
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
)

// CoreServices holds pointers to all core services used to initialize HandlerService.
type CoreServices struct {
	AdminService           *admin.Service
	ContactInfoService     *contactinfo.Service
	CredentialsService     *credentials.Service
	EducationService       *education.Service
	EducationLevelService  *educationlevel.Service
	FamilyInfoService      *familyinfo.Service
	MilitaryDetailsService *militarydetails.Service
	PersonService          *person.Service
	PhysicalInfoService    *physicalinfo.Service
	PhysicalStatusService  *physicalstatus.Service
	RoleService            *role.Service
	SkillsService          *skills.Service
	RankService            *rank.Service
	BloodGroupService      *bloodgroup.Service
	ReligionService        *religion.Service
	PersonTypeService      *persontype.Service
	GenderService          *gender.Service

	MedicalProfile      *medicalprofile.Service
	HospitalDispatch    *hospitaldispatch.Service
	Prescription        *prescriptions.Service
	PsychologicalStatus *psychologicalstatus.Service
	Medicine            *medicines.Service
}

// HandlerService aggregates core services for API handlers.
type HandlerService struct {
	AdminService          *admin.Service
	ContactInfoService    *contactinfo.Service
	CredentialsService    *credentials.Service
	EducationService      *education.Service
	EducationLevelService *educationlevel.Service

	FamilyInfoService      *familyinfo.Service
	MilitaryDetailsService *militarydetails.Service
	PersonService          *person.Service
	PhysicalInfoService    *physicalinfo.Service
	PhysicalStatusService  *physicalstatus.Service
	RoleService            *role.Service
	SkillsService          *skills.Service
	BloodGroupService      *bloodgroup.Service
	RankService            *rank.Service
	ReligionService        *religion.Service
	PersonTypeService      *persontype.Service
	GenderService          *gender.Service

	MedicalProfile      *medicalprofile.Service
	HospitalDispatch    *hospitaldispatch.Service
	Prescription        *prescriptions.Service
	PsychologicalStatus *psychologicalstatus.Service
	Medicine            *medicines.Service
}

// NewHandlerService creates a new HandlerService instance with the provided core services.
func NewHandlerService(cs *CoreServices) *HandlerService {
	return &HandlerService{
		AdminService:           cs.AdminService,
		ContactInfoService:     cs.ContactInfoService,
		CredentialsService:     cs.CredentialsService,
		EducationService:       cs.EducationService,
		EducationLevelService:  cs.EducationLevelService,
		FamilyInfoService:      cs.FamilyInfoService,
		MilitaryDetailsService: cs.MilitaryDetailsService,
		PersonService:          cs.PersonService,
		PhysicalInfoService:    cs.PhysicalInfoService,
		PhysicalStatusService:  cs.PhysicalStatusService,
		RoleService:            cs.RoleService,
		SkillsService:          cs.SkillsService,
		BloodGroupService:      cs.BloodGroupService,
		RankService:            cs.RankService,
		ReligionService:        cs.ReligionService,
		GenderService:          cs.GenderService,

		MedicalProfile:      cs.MedicalProfile,
		HospitalDispatch:    cs.HospitalDispatch,
		Prescription:        cs.Prescription,
		PsychologicalStatus: cs.PsychologicalStatus,
		Medicine:            cs.Medicine,
		PersonTypeService:   cs.PersonTypeService,
	}
}
