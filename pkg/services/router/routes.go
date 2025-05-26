// pkg/services/router/routes.go
package router

import (
	"backend/internal/api/databaseservice"
	"backend/internal/api/personinfoservice/admin"
	"backend/internal/api/personinfoservice/credentials"
	"backend/internal/api/personinfoservice/person"
	"backend/internal/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterDatabaseRoutes(r *gin.Engine, h *databaseservice.Handler) {
	r.Use(middleware.SetHeaders())
	// Admin Routes
	r.POST("/admin", h.CreateAdmin)
	r.GET("/admin/:id", h.GetAdminByID)
	r.GET("/admin/username/:username", h.GetAdminByUsername)
	r.GET("/admins", h.GetAllAdmins)
	r.PUT("/admin/:id", h.UpdateAdmin)
	r.PUT("/admin/:id/password", h.UpdateAdminPassword)
	r.DELETE("/admin/:id", h.DeleteAdmin)
	r.DELETE("/admin/:id/soft", h.SoftDeleteAdmin)

	// Person Routes
	r.POST("/person", h.CreatePerson)
	r.GET("/person/:national_id", h.GetPersonByID)
	r.GET("/persons", h.GetAllPersons)
	r.PUT("/person/:national_id", h.UpdatePerson)
	r.PUT("/person/:national_id/contact", h.UpdateContactInfo)
	r.PUT("/person/:national_id/military", h.UpdateMilitaryDetails)
	r.DELETE("/person/:national_id", h.DeletePerson)
	r.DELETE("/person/:national_id/hard", h.DeletePersonHard)
	r.GET("/persons/search", h.SearchPersonsByName)       // Query params: first_name, last_name
	r.GET("/persons/filter", h.FilterPersonsByPersonType) // Query param: person_type_id

	// Credentials Routes
	r.POST("/credentials", h.CreateCredentials)
	r.GET("/credentials/:admin_id", h.GetCredentialsByAdminID)
	r.GET("/credentials", h.GetAllCredentials)
	r.GET("/credentials/deleted", h.GetSoftDeletedCredentials)
	r.PUT("/credentials/:admin_id", h.UpdateCredentials)
	r.POST("/credentials/:admin_id/dynamic", h.UpdateDynamicToken)
	r.DELETE("/credentials/:admin_id", h.DeleteCredentials)
	r.DELETE("/credentials/:admin_id/hard", h.DeleteCredentialsHard)
	r.GET("/credentials/:admin_id/static-token", h.GetStaticTokenByAdminID)
	r.GET("/credentials/:admin_id/dynamic-token", h.GetDynamicTokenByAdminID)
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
