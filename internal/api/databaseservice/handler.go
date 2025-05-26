package databaseservice

import (
	"backend/internal/core/admin"
	"backend/internal/core/credentials"
	"backend/internal/core/person"
	"backend/pkg/security"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Handler struct {
	adminService       *admin.Service
	personService      *person.Service
	credentialsService *credentials.Service
}

func NewHandler(adminService *admin.Service, personService *person.Service, credentialsService *credentials.Service) *Handler {
	return &Handler{
		adminService:       adminService,
		personService:      personService,
		credentialsService: credentialsService,
	}
}

// ### Admin Handlers

// **CreateAdmin** creates a new admin with additional fields and actionBy from the header.
func (h *Handler) CreateAdmin(c *gin.Context) {
	var req struct {
		NationalIDNumber string `json:"national_id_number" binding:"required"`
		UserName         string `json:"user_name" binding:"required"`
		Password         string `json:"password" binding:"required"`
		RoleID           int64  `json:"role_id" binding:"required"`
		CredentialsID    int64  `json:"credentials_id" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	actionBy := c.GetHeader("X-Action-By")
	if actionBy == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing X-Action-By header"})
		return
	}
	hash, err := security.HashPassword(req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to hash password"})
		return
	}
	id, err := h.adminService.CreateAdmin(req.NationalIDNumber, req.UserName, hash, req.RoleID, req.CredentialsID, actionBy)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"id": id})
}

// **GetAdminByID** retrieves an admin by their ID (unchanged).
func (h *Handler) GetAdminByID(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	admin, err := h.adminService.GetAdminByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "admin not found"})
		return
	}
	c.JSON(http.StatusOK, admin)
}

// **GetAdminByUsername** retrieves an admin by their username (unchanged).
func (h *Handler) GetAdminByUsername(c *gin.Context) {
	username := c.Param("username")
	admin, err := h.adminService.GetAdminByUsername(username)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "admin not found"})
		return
	}
	c.JSON(http.StatusOK, admin)
}

// **GetAllAdmins** retrieves all admins (unchanged).
func (h *Handler) GetAllAdmins(c *gin.Context) {
	admins, err := h.adminService.GetAllAdmins()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, admins)
}

// **UpdateAdmin** updates an admin with actionBy from the header.
func (h *Handler) UpdateAdmin(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	var updates map[string]interface{}
	if err := c.ShouldBindJSON(&updates); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	actionBy := c.GetHeader("X-Action-By")
	if actionBy == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing X-Action-By header"})
		return
	}
	if err := h.adminService.UpdateAdmin(id, updates, actionBy); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "admin updated"})
}

// **UpdateAdminPassword** updates an admin's password with actionBy.
func (h *Handler) UpdateAdminPassword(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	var req struct {
		Password string `json:"password" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	actionBy := c.GetHeader("X-Action-By")
	if actionBy == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing X-Action-By header"})
		return
	}
	if err := h.adminService.UpdateAdminPassword(id, req.Password, actionBy); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "password updated"})
}

// **DeleteAdmin** deletes an admin with actionBy (updated from Delete).
func (h *Handler) DeleteAdmin(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	actionBy := c.GetHeader("X-Action-By")
	if actionBy == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing X-Action-By header"})
		return
	}
	if err := h.adminService.DeleteAdmin(id, actionBy); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "admin deleted"})
}

// **SoftDeleteAdmin** soft deletes an admin with actionBy.
func (h *Handler) SoftDeleteAdmin(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	actionBy := c.GetHeader("X-Action-By")
	if actionBy == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing X-Action-By header"})
		return
	}
	if err := h.adminService.SoftDeleteAdmin(id, actionBy); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "admin soft deleted"})
}

// ### Person Handlers

// **CreatePerson** creates a person with full struct and actionBy.
func (h *Handler) CreatePerson(c *gin.Context) {
	var req struct {
		NationalIDNumber  string `json:"national_id_number" binding:"required"`
		FirstName         string `json:"first_name" binding:"required"`
		LastName          string `json:"last_name" binding:"required"`
		FamilyInfoID      int64  `json:"family_info_id" binding:"required"`
		PhysicalInfoID    int64  `json:"physical_info_id" binding:"required"`
		ContactInfoID     int64  `json:"contact_info_id" binding:"required"`
		SkillsID          int64  `json:"skills_id" binding:"required"`
		BirthDate         string `json:"birth_date" binding:"required"`
		ReligionID        int64  `json:"religion_id" binding:"required"`
		PersonTypeID      int64  `json:"person_type_id" binding:"required"`
		MilitaryDetailsID int64  `json:"military_details_id" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	actionBy := c.GetHeader("X-Action-By")
	if actionBy == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing X-Action-By header"})
		return
	}
	birthDate, err := time.Parse("2006-01-02", req.BirthDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid birth_date format, should be YYYY-MM-DD"})
		return
	}
	person := &person.Person{
		NationalIDNumber:  req.NationalIDNumber,
		FirstName:         req.FirstName,
		LastName:          req.LastName,
		FamilyInfoID:      req.FamilyInfoID,
		PhysicalInfoID:    req.PhysicalInfoID,
		ContactInfoID:     req.ContactInfoID,
		SkillsID:          req.SkillsID,
		BirthDate:         birthDate,
		ReligionID:        req.ReligionID,
		PersonTypeID:      req.PersonTypeID,
		MilitaryDetailsID: req.MilitaryDetailsID,
		DeletedAt:         0,
	}
	id, err := h.personService.CreatePerson(person, actionBy)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"national_id_number": id})
}

// **GetPersonByID** retrieves a person by national ID (updated method name).
func (h *Handler) GetPersonByID(c *gin.Context) {
	nationalID := c.Param("national_id")
	person, err := h.personService.GetPersonByID(nationalID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "person not found"})
		return
	}
	c.JSON(http.StatusOK, person)
}

// **GetAllPersons** retrieves all persons (updated method name).
func (h *Handler) GetAllPersons(c *gin.Context) {
	persons, err := h.personService.GetAllPersons()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, persons)
}

// **UpdatePerson** updates a person with actionBy (updated method name).
func (h *Handler) UpdatePerson(c *gin.Context) {
	nationalID := c.Param("national_id")
	var updates map[string]interface{}
	if err := c.ShouldBindJSON(&updates); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	actionBy := c.GetHeader("X-Action-By")
	if actionBy == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing X-Action-By header"})
		return
	}
	if err := h.personService.UpdatePerson(nationalID, updates, actionBy); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "person updated"})
}

// **UpdateContactInfo** updates a person's contact info with actionBy.
func (h *Handler) UpdateContactInfo(c *gin.Context) {
	nationalID := c.Param("national_id")
	var req struct {
		ContactInfoID int64 `json:"contact_info_id" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	actionBy := c.GetHeader("X-Action-By")
	if actionBy == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing X-Action-By header"})
		return
	}
	if err := h.personService.UpdateContactInfo(nationalID, req.ContactInfoID, actionBy); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "contact info updated"})
}

// **UpdateMilitaryDetails** updates a person's military details with actionBy.
func (h *Handler) UpdateMilitaryDetails(c *gin.Context) {
	nationalID := c.Param("national_id")
	var req struct {
		MilitaryDetailsID int64 `json:"military_details_id" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	actionBy := c.GetHeader("X-Action-By")
	if actionBy == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing X-Action-By header"})
		return
	}
	if err := h.personService.UpdateMilitaryDetails(nationalID, req.MilitaryDetailsID, actionBy); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "military details updated"})
}

// **DeletePerson** soft deletes a person with actionBy (updated method name).
func (h *Handler) DeletePerson(c *gin.Context) {
	nationalID := c.Param("national_id")
	actionBy := c.GetHeader("X-Action-By")
	if actionBy == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing X-Action-By header"})
		return
	}
	if err := h.personService.DeletePerson(nationalID, actionBy); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "person soft deleted"})
}

// **DeletePersonHard** hard deletes a person with actionBy.
func (h *Handler) DeletePersonHard(c *gin.Context) {
	nationalID := c.Param("national_id")
	actionBy := c.GetHeader("X-Action-By")
	if actionBy == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing X-Action-By header"})
		return
	}
	if err := h.personService.DeletePersonHard(nationalID, actionBy); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "person hard deleted"})
}

// **SearchPersonsByName** searches persons by name with actionBy.
func (h *Handler) SearchPersonsByName(c *gin.Context) {
	firstName := c.Query("first_name")
	lastName := c.Query("last_name")
	actionBy := c.GetHeader("X-Action-By")
	if actionBy == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing X-Action-By header"})
		return
	}
	persons, err := h.personService.SearchPersonsByName(firstName, lastName, actionBy)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, persons)
}

// **FilterPersonsByPersonType** filters persons by type with actionBy.
func (h *Handler) FilterPersonsByPersonType(c *gin.Context) {
	personTypeIDStr := c.Query("person_type_id")
	personTypeID, err := strconv.ParseInt(personTypeIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid person_type_id"})
		return
	}
	actionBy := c.GetHeader("X-Action-By")
	if actionBy == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing X-Action-By header"})
		return
	}
	persons, err := h.personService.FilterPersonsByPersonType(personTypeID, actionBy)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, persons)
}

// ### Credentials Handlers

// **CreateCredentials** creates credentials with actionBy (updated method name).
func (h *Handler) CreateCredentials(c *gin.Context) {
	var req struct {
		AdminID      string `json:"admin_id" binding:"required"`
		StaticToken  string `json:"static_token" binding:"required"`
		DynamicToken string `json:"dynamic_token" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	adminID, err := uuid.Parse(req.AdminID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid admin_id"})
		return
	}
	actionBy := c.GetHeader("X-Action-By")
	if actionBy == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing X-Action-By header"})
		return
	}
	cred, err := h.credentialsService.CreateCredentials(adminID, req.StaticToken, req.DynamicToken, actionBy)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, cred)
}

// **GetCredentialsByAdminID** retrieves credentials by admin ID (updated method name).
func (h *Handler) GetCredentialsByAdminID(c *gin.Context) {
	adminID, err := uuid.Parse(c.Param("admin_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid admin_id"})
		return
	}
	cred, err := h.credentialsService.GetCredentialsByAdminID(adminID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "credentials not found"})
		return
	}
	c.JSON(http.StatusOK, cred)
}

// **GetAllCredentials** retrieves all credentials.
func (h *Handler) GetAllCredentials(c *gin.Context) {
	creds, err := h.credentialsService.GetAllCredentials()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, creds)
}

// **GetSoftDeletedCredentials** retrieves soft-deleted credentials.
func (h *Handler) GetSoftDeletedCredentials(c *gin.Context) {
	creds, err := h.credentialsService.GetSoftDeletedCredentials()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, creds)
}

// **UpdateCredentials** updates credentials with actionBy.
func (h *Handler) UpdateCredentials(c *gin.Context) {
	adminID, err := uuid.Parse(c.Param("admin_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid admin_id"})
		return
	}
	var updates map[string]interface{}
	if err := c.ShouldBindJSON(&updates); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	actionBy := c.GetHeader("X-Action-By")
	if actionBy == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing X-Action-By header"})
		return
	}
	if err := h.credentialsService.UpdateCredentials(adminID, updates, actionBy); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "credentials updated"})
}

// **UpdateDynamicToken** updates the dynamic token with actionBy (updated method name).
func (h *Handler) UpdateDynamicToken(c *gin.Context) {
	adminID, err := uuid.Parse(c.Param("admin_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid admin_id"})
		return
	}
	var req struct {
		DynamicToken string `json:"dynamic_token" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	actionBy := c.GetHeader("X-Action-By")
	if actionBy == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing X-Action-By header"})
		return
	}
	if err := h.credentialsService.UpdateDynamicTokenByAdminID(adminID, req.DynamicToken, actionBy); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "dynamic token updated"})
}

// **DeleteCredentials** soft deletes credentials with actionBy (updated method name).
func (h *Handler) DeleteCredentials(c *gin.Context) {
	adminID, err := uuid.Parse(c.Param("admin_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid admin_id"})
		return
	}
	actionBy := c.GetHeader("X-Action-By")
	if actionBy == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing X-Action-By header"})
		return
	}
	if err := h.credentialsService.DeleteCredentials(adminID, actionBy); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "credentials soft deleted"})
}

// **DeleteCredentialsHard** hard deletes credentials with actionBy.
func (h *Handler) DeleteCredentialsHard(c *gin.Context) {
	adminID, err := uuid.Parse(c.Param("admin_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid admin_id"})
		return
	}
	actionBy := c.GetHeader("X-Action-By")
	if actionBy == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing X-Action-By header"})
		return
	}
	if err := h.credentialsService.DeleteCredentialsHard(adminID, actionBy); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "credentials hard deleted"})
}

// **GetStaticTokenByAdminID** retrieves the static token for an admin.
func (h *Handler) GetStaticTokenByAdminID(c *gin.Context) {
	adminID, err := uuid.Parse(c.Param("admin_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid admin_id"})
		return
	}
	token, err := h.credentialsService.GetStaticTokenByAdminID(adminID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "static token not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"static_token": token})
}

// **GetDynamicTokenByAdminID** retrieves the dynamic token for an admin.
func (h *Handler) GetDynamicTokenByAdminID(c *gin.Context) {
	adminID, err := uuid.Parse(c.Param("admin_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid admin_id"})
		return
	}
	token, err := h.credentialsService.GetDynamicTokenByAdminID(adminID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "dynamic token not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"dynamic_token": token})
}

// Add other person handlers (GetPersonByID, GetAllPersons, UpdatePerson, UpdateContactInfo, DeletePerson, DeletePersonHard, SearchPersonsByName, FilterPersonsByPersonType)

// Credentials, Role, Rank, MilitaryDetails handlers follow similar patterns
