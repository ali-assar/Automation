// pkg/services/router/routes.go
package router

import (
	"backend/internal/api/databaseservice/handler"
	"backend/internal/api/personinfoservice/admin"
	"backend/internal/api/personinfoservice/credentials"
	"backend/internal/api/personinfoservice/person"
	"backend/internal/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterDatabaseRoutes(r *gin.Engine, s *handler.Service) {
	r.Use(middleware.SetHeaders())
	// Admin Routes
	adminGroup := r.Group("/admin")
	{
		adminGroup.POST("", handler.CreateAdmin(s))
		adminGroup.GET(":id", handler.GetAdminByID(s))
		adminGroup.GET("/username/:username", handler.GetAdminByUsername(s))
		adminGroup.GET("", handler.GetAllAdmins(s))
		adminGroup.PUT(":id", handler.UpdateAdmin(s))
		adminGroup.PUT("/password/:id", handler.UpdateAdminPassword(s))
		adminGroup.DELETE(":id", handler.DeleteAdmin(s))
		adminGroup.DELETE("/soft/:id", handler.SoftDeleteAdmin(s))
	}

	// ContactInfo Routes
	contactInfoGroup := r.Group("/contactinfo")
	{
		contactInfoGroup.POST("", handler.CreateContactInfo(s))
		contactInfoGroup.GET(":id", handler.GetContactInfoByID(s))
		contactInfoGroup.GET("/email/:email", handler.GetContactInfoByEmail(s))
		contactInfoGroup.GET("", handler.GetAllContactInfos(s))
		contactInfoGroup.PUT(":id", handler.UpdateContactInfo(s))
		contactInfoGroup.DELETE(":id", handler.DeleteContactInfo(s))
		contactInfoGroup.DELETE("/hard/:id", handler.DeleteContactInfoHard(s))
	}

	// Credentials Routes
	credentialsGroup := r.Group("/credentials")
	{
		credentialsGroup.POST("", handler.CreateCredentials(s))
		credentialsGroup.GET(":admin_id", handler.GetCredentialsByAdminID(s))
		credentialsGroup.GET("", handler.GetAllCredentials(s))
		credentialsGroup.GET("/softdeleted", handler.GetSoftDeletedCredentials(s))
		credentialsGroup.PUT(":admin_id", handler.UpdateCredentials(s))
		credentialsGroup.PUT("/dynamic_token/:admin_id", handler.UpdateDynamicToken(s))
		credentialsGroup.DELETE(":admin_id", handler.DeleteCredentials(s))
		credentialsGroup.DELETE("/hard/:admin_id", handler.DeleteCredentialsHard(s))
		credentialsGroup.GET("/static_token/:admin_id", handler.GetStaticTokenByAdminID(s))
		credentialsGroup.GET("/dynamic_token/:admin_id", handler.GetDynamicTokenByAdminID(s))
	}

	// Education Routes
	educationGroup := r.Group("/education")
	{
		educationGroup.POST("", handler.CreateEducation(s))
		educationGroup.GET(":id", handler.GetEducationByID(s))
		educationGroup.GET("", handler.GetAllEducations(s))
		educationGroup.PUT(":id", handler.UpdateEducation(s))
		educationGroup.DELETE(":id", handler.DeleteEducation(s))
		educationGroup.GET("/search", handler.SearchEducationsByUniversity(s))
	}

	// FamilyInfo Routes
	familyInfoGroup := r.Group("/familyinfo")
	{
		familyInfoGroup.POST("", handler.CreateFamilyInfo(s))
		familyInfoGroup.GET(":id", handler.GetFamilyInfoByID(s))
		familyInfoGroup.GET("", handler.GetAllFamilyInfos(s))
		familyInfoGroup.PUT(":id", handler.UpdateFamilyInfo(s))
		familyInfoGroup.DELETE(":id", handler.DeleteFamilyInfo(s))
	}

	// MilitaryDetails Routes
	militaryDetailsGroup := r.Group("/militarydetails")
	{
		militaryDetailsGroup.POST("", handler.CreateMilitaryDetails(s))
		militaryDetailsGroup.GET(":id", handler.GetMilitaryDetailsByID(s))
		militaryDetailsGroup.GET("", handler.GetAllMilitaryDetails(s))
		militaryDetailsGroup.PUT(":id", handler.UpdateMilitaryDetails(s))
		militaryDetailsGroup.DELETE(":id", handler.DeleteMilitaryDetails(s))
	}

	// Person Routes
	personGroup := r.Group("/person")
	{
		personGroup.POST("", handler.CreatePerson(s))
		personGroup.GET(":national_id", handler.GetPersonByID(s))
		personGroup.GET("", handler.GetAllPersons(s))
		personGroup.PUT(":national_id", handler.UpdatePerson(s))
		personGroup.PUT("/contactinfo/:national_id", handler.UpdateContactInfo(s))
		personGroup.PUT("/militarydetails/:national_id", handler.UpdateMilitaryDetails(s))
		personGroup.DELETE(":national_id", handler.DeletePerson(s))
		personGroup.DELETE("/hard/:national_id", handler.DeletePersonHard(s))
		personGroup.GET("/search", handler.SearchPersonsByName(s))
		personGroup.GET("/filter", handler.FilterPersonsByPersonType(s))
	}

	// PhysicalInfo Routes
	physicalInfoGroup := r.Group("/physicalinfo")
	{

		physicalInfoGroup.POST("", handler.CreatePhysicalInfo(s))
		physicalInfoGroup.GET(":id", handler.GetPhysicalInfoByID(s))
		physicalInfoGroup.GET("", handler.GetAllPhysicalInfos(s))
		physicalInfoGroup.PUT(":id", handler.UpdatePhysicalInfo(s))
		physicalInfoGroup.DELETE(":id", handler.DeletePhysicalInfo(s))
	}

	// PhysicalStatus Routes
	physicalStatusGroup := r.Group("/physicalstatus")
	{
		physicalStatusGroup.POST("", handler.CreatePhysicalStatus(s))
		physicalStatusGroup.GET(":id", handler.GetPhysicalStatusByID(s))
		physicalStatusGroup.GET("", handler.GetAllPhysicalStatuses(s))
		physicalStatusGroup.PUT(":id", handler.UpdatePhysicalStatus(s))
		physicalStatusGroup.DELETE(":id", handler.DeletePhysicalStatus(s))
	}

	// Role Routes
	roleGroup := r.Group("/role")
	{
		roleGroup.POST("", handler.CreateRole(s))
		roleGroup.GET(":id", handler.GetRoleByID(s))
		roleGroup.GET("/type/:type", handler.GetRoleByType(s))
		roleGroup.GET("", handler.GetAllRoles(s))
		roleGroup.PUT(":id", handler.UpdateRole(s))
		roleGroup.DELETE(":id", handler.DeleteRole(s))
	}

	// Skills Routes
	skillsGroup := r.Group("/skills")
	{
		skillsGroup.POST("", handler.CreateSkills(s))
		skillsGroup.GET(":id", handler.GetSkillsByID(s))
		skillsGroup.GET("/education/:education_id", handler.GetSkillsByEducationID(s))
		skillsGroup.GET("", handler.GetAllSkills(s))
		skillsGroup.PUT(":id", handler.UpdateSkills(s))
		skillsGroup.DELETE(":id", handler.DeleteSkills(s))
	}
}

func RegisterPersonInfoRoutes(r *gin.Engine, adminCtrl *admin.Controller, personCtrl *person.Controller, credCtrl *credentials.Controller) {
	r.Use(middleware.SetHeaders())
	// Public Routes
	r.POST("/login", adminCtrl.Login)
	// Static Protected Routes
	staticGroup := r.Group("/static")
	staticGroup.Use(middleware.StaticAuth(adminCtrl, credCtrl))
	{
		staticGroup.GET("/admins", adminCtrl.GetAllAdmins)
		staticGroup.GET("/admin/:id", adminCtrl.GetAdmin)
		staticGroup.GET("/persons", personCtrl.GetAllPersons)
		staticGroup.GET("/person/:national_id", personCtrl.GetPerson)
	}
	// Dynamic Protected Routes
	dynamicGroup := r.Group("/dynamic")
	dynamicGroup.Use(middleware.DynamicAuth(adminCtrl, credCtrl))
	{
		dynamicGroup.POST("/admin", adminCtrl.CreateAdmin)
		dynamicGroup.POST("/person", personCtrl.CreatePerson)
		dynamicGroup.PUT("/person/:national_id", personCtrl.UpdatePerson)
		dynamicGroup.DELETE("/person/:national_id", personCtrl.DeletePerson)
		dynamicGroup.POST("/credentials/:admin_id/dynamic", credCtrl.UpdateDynamicToken)
	}
}
