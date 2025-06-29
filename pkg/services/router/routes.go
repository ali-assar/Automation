package router

import (
	"backend/internal/api"
	"backend/internal/api/battalion"
	"backend/internal/api/role"
	"backend/internal/api/roleaccess"
	"backend/internal/api/user"
	"backend/internal/api/userrole"
	"backend/internal/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, s *api.HandlerService) {

	r.Use(middleware.CORSMiddleware())

	apiRouterGroup := r.Group("/api")

	// Public Routes
	public := apiRouterGroup.Group("/personinfo")
	{
		public.POST("/login", api.Login(s))
	}

	// Static Protected Routes (GET operations)
	staticGroup := public.Group("/static").Use(middleware.StaticAuth(s))
	{
		// Admin
		staticGroup.GET("/admins", api.GetAllAdmins(s))
		staticGroup.GET("/admin/:id", api.GetAdminByID(s))
		staticGroup.GET("/admin/username/:username", api.GetAdminByUsername(s))
		// ContactInfo
		staticGroup.GET("/contactinfo/:id", api.GetContactInfoByID(s))
		staticGroup.GET("/contactinfo/email/:email", api.GetContactInfoByEmail(s))
		staticGroup.GET("/contactinfos", api.GetAllContactInfos(s))
		// Credentials
		staticGroup.GET("/credentials/:admin_id", api.GetCredentialsByAdminID(s))
		staticGroup.GET("/credentials", api.GetAllCredentials(s))
		staticGroup.GET("/credentials/softdeleted", api.GetSoftDeletedCredentials(s))
		staticGroup.GET("/credentials/static_token/:admin_id", api.GetStaticTokenByAdminID(s))
		staticGroup.GET("/credentials/dynamic_token/:admin_id", api.GetDynamicTokenByAdminID(s))
		// Education
		staticGroup.GET("/education/:id", api.GetEducationByID(s))
		staticGroup.GET("/educations", api.GetAllEducations(s))
		staticGroup.GET("/education/search", api.SearchEducationsByUniversity(s))
		// FamilyInfo
		staticGroup.GET("/familyinfo/:id", api.GetFamilyInfoByID(s))
		staticGroup.GET("/familyinfos", api.GetAllFamilyInfos(s))
		// MilitaryDetails
		staticGroup.GET("/militarydetails/:id", api.GetMilitaryDetailsByID(s))
		staticGroup.GET("/militarydetails", api.GetAllMilitaryDetails(s))
		// Person
		staticGroup.GET("/persons", api.GetAllPersons(s))
		staticGroup.GET("/person/:national_id", api.GetPersonByID(s))
		staticGroup.GET("/person/search", api.SearchPersonsByName(s))
		staticGroup.GET("/person/filter", api.FilterPersonsByPersonType(s))
		// PhysicalInfo
		staticGroup.GET("/physicalinfo/:id", api.GetPhysicalInfoByID(s))
		staticGroup.GET("/physicalinfos", api.GetAllPhysicalInfos(s))

		// Skills
		staticGroup.GET("/skills/:id", api.GetSkillsByID(s))
		staticGroup.GET("/skills/education/:education_id", api.GetSkillsByEducationID(s))
		staticGroup.GET("/skills", api.GetAllSkills(s))

		// Get Info for person creation
		staticGroup.GET("/static-tables", api.GetStaticTables(s))

		staticGroup.GET("/hospitaldispatch/:id", api.GetHospitalDispatchByID(s))
		staticGroup.GET("/hospitaldispatches", api.GetAllHospitalDispatches(s))
		staticGroup.GET("/medicalprofile/:id", api.GetMedicalProfileByID(s))
		staticGroup.GET("/medicalprofile/person/:person_id", api.GetMedicalProfileByPersonID(s))
		staticGroup.GET("/medicalprofiles", api.GetAllMedicalProfiles(s))
		staticGroup.GET("/medicine/:id", api.GetMedicineByID(s))
		staticGroup.GET("/medicines", api.GetAllMedicines(s))
		staticGroup.GET("/prescription/:id", api.GetPrescriptionByID(s))
		staticGroup.GET("/prescriptions", api.GetAllPrescriptions(s))
		staticGroup.GET("/psychologicalstatus/:id", api.GetPsychologicalStatusByID(s))
		staticGroup.GET("/psychologicalstatuses", api.GetAllPsychologicalStatuses(s))
		staticGroup.GET("/visit/:id", api.GetVisitByID(s))
		staticGroup.GET("/visit/person/:person_id", api.GetVisitByPersonID(s))
		staticGroup.GET("/visits", api.GetAllVisits(s))
		staticGroup.GET("/prescription/visit/:visit_id", api.GetPrescriptionsByVisitID(s))
	}

	// Dynamic Protected Routes (POST, PUT, DELETE)
	dynamicGroup := public.Group("/dynamic").Use(middleware.DynamicAuth(s))
	{

		// Create a full person
		dynamicGroup.POST("/persons/full", api.CreateFullPerson(s)) // Admin
		dynamicGroup.POST("/admin", api.CreateAdmin(s))
		dynamicGroup.PUT("/admin/:id", api.UpdateAdmin(s))
		dynamicGroup.PUT("/admin/password/:id", api.UpdateAdminPassword(s))
		dynamicGroup.DELETE("/admin/:id", api.DeleteAdmin(s))
		dynamicGroup.DELETE("/admin/soft/:id", api.SoftDeleteAdmin(s))
		// ContactInfo
		dynamicGroup.POST("/contactinfo", api.CreateContactInfo(s))
		dynamicGroup.PUT("/contactinfo/:id", api.UpdateContactInfo(s))
		dynamicGroup.DELETE("/contactinfo/:id", api.DeleteContactInfo(s))
		dynamicGroup.DELETE("/contactinfo/hard/:id", api.DeleteContactInfoHard(s))
		// Credentials
		dynamicGroup.POST("/credentials", api.CreateCredentials(s))
		dynamicGroup.PUT("/credentials/:admin_id", api.UpdateCredentials(s))
		dynamicGroup.PUT("/credentials/dynamic_token/:admin_id", api.UpdateDynamicToken(s))
		dynamicGroup.DELETE("/credentials/:admin_id", api.DeleteCredentials(s))
		dynamicGroup.DELETE("/credentials/hard/:admin_id", api.DeleteCredentialsHard(s))
		// Education
		dynamicGroup.POST("/education", api.CreateEducation(s))
		dynamicGroup.PUT("/education/:id", api.UpdateEducation(s))
		dynamicGroup.DELETE("/education/:id", api.DeleteEducation(s))
		// FamilyInfo
		dynamicGroup.POST("/familyinfo", api.CreateFamilyInfo(s))
		dynamicGroup.PUT("/familyinfo/:id", api.UpdateFamilyInfo(s))
		dynamicGroup.DELETE("/familyinfo/:id", api.DeleteFamilyInfo(s))
		// MilitaryDetails
		dynamicGroup.POST("/militarydetails", api.CreateMilitaryDetails(s))
		dynamicGroup.PUT("/militarydetails/:id", api.UpdateMilitaryDetails(s))
		dynamicGroup.DELETE("/militarydetails/:id", api.DeleteMilitaryDetails(s))
		// Person
		dynamicGroup.POST("/person", api.CreatePerson(s))
		dynamicGroup.PUT("/person/:national_id", api.UpdatePerson(s))
		dynamicGroup.PUT("/person/contactinfo/:national_id", api.UpdateContactInfo(s))
		dynamicGroup.PUT("/person/militarydetails/:national_id", api.UpdateMilitaryDetails(s))
		dynamicGroup.DELETE("/person/:national_id", api.DeletePerson(s))
		dynamicGroup.DELETE("/person/hard/:national_id", api.DeletePersonHard(s))
		// PhysicalInfo
		dynamicGroup.POST("/physicalinfo", api.CreatePhysicalInfo(s))
		dynamicGroup.PUT("/physicalinfo/:id", api.UpdatePhysicalInfo(s))
		dynamicGroup.DELETE("/physicalinfo/:id", api.DeletePhysicalInfo(s))

		// Skills
		dynamicGroup.POST("/skills", api.CreateSkills(s))
		dynamicGroup.PUT("/skills/:id", api.UpdateSkills(s))
		dynamicGroup.DELETE("/skills/:id", api.DeleteSkills(s))
		// Medical-related dynamic routes
		dynamicGroup.POST("/hospitaldispatch", api.CreateHospitalDispatch(s))
		dynamicGroup.PUT("/hospitaldispatch/:id", api.UpdateHospitalDispatch(s))
		dynamicGroup.DELETE("/hospitaldispatch/:id", api.DeleteHospitalDispatch(s))
		dynamicGroup.DELETE("/hospitaldispatch/hard/:id", api.DeleteHospitalDispatchHard(s))
		dynamicGroup.POST("/medicalprofile", api.CreateMedicalProfile(s))
		dynamicGroup.PUT("/medicalprofile/:id", api.UpdateMedicalProfile(s))
		dynamicGroup.DELETE("/medicalprofile/:id", api.DeleteMedicalProfile(s))
		dynamicGroup.DELETE("/medicalprofile/hard/:id", api.DeleteMedicalProfileHard(s))
		dynamicGroup.POST("/medicine", api.CreateMedicine(s))
		dynamicGroup.PUT("/medicine/:id", api.UpdateMedicine(s))
		dynamicGroup.DELETE("/medicine/:id", api.DeleteMedicine(s))
		dynamicGroup.DELETE("/medicine/hard/:id", api.DeleteMedicineHard(s))
		dynamicGroup.POST("/prescription", api.CreatePrescription(s))
		dynamicGroup.PUT("/prescription/:id", api.UpdatePrescription(s))
		dynamicGroup.DELETE("/prescription/:id", api.DeletePrescription(s))
		dynamicGroup.DELETE("/prescription/hard/:id", api.DeletePrescriptionHard(s))
		dynamicGroup.POST("/psychologicalstatus", api.CreatePsychologicalStatus(s))
		dynamicGroup.PUT("/psychologicalstatus/:id", api.UpdatePsychologicalStatus(s))
		dynamicGroup.DELETE("/psychologicalstatus/:id", api.DeletePsychologicalStatus(s))

		dynamicGroup.POST("/visit", api.CreateVisit(s))
		dynamicGroup.PUT("/visit/:id", api.UpdateVisit(s))
		dynamicGroup.DELETE("/visit/:id", api.DeleteVisit(s))
		dynamicGroup.DELETE("/visit/hard/:id", api.DeleteVisitHard(s))
	}

	battalion.AddRoutes(apiRouterGroup)
	role.AddRoutes(apiRouterGroup)
	roleaccess.AddRoutes(apiRouterGroup)
	user.AddRoutes(apiRouterGroup)
	userrole.AddRoutes(apiRouterGroup)

}
