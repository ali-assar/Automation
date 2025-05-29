package router

import (
	"backend/internal/api/api"
	"backend/internal/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, s *api.HandlerService) {
	r.Use(middleware.SetHeaders())

	// Public Routes
	public := r.Group("/api/personinfo")
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
		// Role
		staticGroup.GET("/role/:id", api.GetRoleByID(s))
		staticGroup.GET("/role/type/:type", api.GetRoleByType(s))
		staticGroup.GET("/roles", api.GetAllRoles(s))
		// Skills
		staticGroup.GET("/skills/:id", api.GetSkillsByID(s))
		staticGroup.GET("/skills/education/:education_id", api.GetSkillsByEducationID(s))
		staticGroup.GET("/skills", api.GetAllSkills(s))

		// Get Info for person creation
		staticGroup.GET("/static-tables", api.GetStaticTables(s))
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
		// Role
		dynamicGroup.POST("/role", api.CreateRole(s))
		dynamicGroup.PUT("/role/:id", api.UpdateRole(s))
		dynamicGroup.DELETE("/role/:id", api.DeleteRole(s))
		// Skills
		dynamicGroup.POST("/skills", api.CreateSkills(s))
		dynamicGroup.PUT("/skills/:id", api.UpdateSkills(s))
		dynamicGroup.DELETE("/skills/:id", api.DeleteSkills(s))
	}
}
