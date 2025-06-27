package api

import (
	"backend/internal/core/person"
	"database/sql"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func CreatePerson(s *HandlerService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req struct {
			NationalIDNumber  string `json:"national_id_number" binding:"required"`
			FirstName         string `json:"first_name" binding:"required"`
			LastName          string `json:"last_name" binding:"required"`
			FamilyInfoID      int64  `json:"family_info_id"`
			ContactInfoID     int64  `json:"contact_info_id"`
			SkillsID          int64  `json:"skills_id"`
			PhysicalInfoID    int64  `json:"physical_info_id"`
			BirthDate         string `json:"birth_date" binding:"required"`
			ReligionID        int64  `json:"religion_id"`
			PersonTypeID      int64  `json:"person_type_id"`
			MilitaryDetailsID int64  `json:"military_details_id"`
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
			NationalIDNumber: req.NationalIDNumber,
			FirstName:        req.FirstName,
			LastName:         req.LastName,
			BirthDate:        birthDate,
			DeletedAt:        0,
		}
		// Set nullable IDs if provided
		if req.FamilyInfoID != 0 {
			person.SetFamilyInfoID(req.FamilyInfoID)
		}
		if req.ContactInfoID != 0 {
			person.SetContactInfoID(req.ContactInfoID)
		}
		if req.SkillsID != 0 {
			person.SetSkillsID(req.SkillsID)
		}
		if req.PhysicalInfoID != 0 {
			person.SetPhysicalInfoID(req.PhysicalInfoID)
		}
		if req.ReligionID != 0 {
			person.SetReligionID(req.ReligionID)
		}
		if req.PersonTypeID != 0 {
			person.SetPersonTypeID(req.PersonTypeID)
		}
		if req.MilitaryDetailsID != 0 {
			person.SetMilitaryDetailsID(req.MilitaryDetailsID)
		}
		id, err := s.PersonService.CreatePerson(person, actionBy)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, gin.H{"national_id_number": id})
	}
}
func GetPersonByID(s *HandlerService) gin.HandlerFunc {
	return func(c *gin.Context) {
		nationalID := c.Param("national_id")
		person, err := s.PersonService.GetPersonByID(nationalID)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "person not found"})
			return
		}
		c.JSON(http.StatusOK, person)
	}
}

func GetAllPersons(s *HandlerService) gin.HandlerFunc {
	return func(c *gin.Context) {
		persons, err := s.PersonService.GetAllPersons()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, persons)
	}
}

func UpdatePerson(s *HandlerService) gin.HandlerFunc {
	return func(c *gin.Context) {
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
		// Handle nullable fields in updates
		if familyInfoID, ok := updates["family_info_id"].(int64); ok && familyInfoID != 0 {
			updates["family_info_id"] = sql.NullInt64{Int64: familyInfoID, Valid: true}
		}
		if contactInfoID, ok := updates["contact_info_id"].(int64); ok && contactInfoID != 0 {
			updates["contact_info_id"] = sql.NullInt64{Int64: contactInfoID, Valid: true}
		}
		if skillsID, ok := updates["skills_id"].(int64); ok && skillsID != 0 {
			updates["skills_id"] = sql.NullInt64{Int64: skillsID, Valid: true}
		}
		if physicalInfoID, ok := updates["physical_info_id"].(int64); ok && physicalInfoID != 0 {
			updates["physical_info_id"] = sql.NullInt64{Int64: physicalInfoID, Valid: true}
		}
		if religionID, ok := updates["religion_id"].(int64); ok && religionID != 0 {
			updates["religion_id"] = sql.NullInt64{Int64: religionID, Valid: true}
		}
		if personTypeID, ok := updates["person_type_id"].(int64); ok && personTypeID != 0 {
			updates["person_type_id"] = sql.NullInt64{Int64: personTypeID, Valid: true}
		}
		if militaryDetailsID, ok := updates["military_details_id"].(int64); ok && militaryDetailsID != 0 {
			updates["military_details_id"] = sql.NullInt64{Int64: militaryDetailsID, Valid: true}
		}
		if err := s.PersonService.UpdatePerson(nationalID, updates, actionBy); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "person updated"})
	}
}
func DeletePerson(s *HandlerService) gin.HandlerFunc {
	return func(c *gin.Context) {
		nationalID := c.Param("national_id")
		actionBy := c.GetHeader("X-Action-By")
		if actionBy == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "missing X-Action-By header"})
			return
		}
		if err := s.PersonService.DeletePerson(nationalID, actionBy); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "person soft deleted"})
	}
}

func DeletePersonHard(s *HandlerService) gin.HandlerFunc {
	return func(c *gin.Context) {
		nationalID := c.Param("national_id")
		actionBy := c.GetHeader("X-Action-By")
		if actionBy == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "missing X-Action-By header"})
			return
		}
		if err := s.PersonService.DeletePersonHard(nationalID, actionBy); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "person hard deleted"})
	}
}

func SearchPersonsByName(s *HandlerService) gin.HandlerFunc {
	return func(c *gin.Context) {
		firstName := c.Query("first_name")
		lastName := c.Query("last_name")
		actionBy := c.GetHeader("X-Action-By")
		if actionBy == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "missing X-Action-By header"})
			return
		}
		persons, err := s.PersonService.SearchPersonsByName(firstName, lastName, actionBy)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, persons)
	}
}

func FilterPersonsByPersonType(s *HandlerService) gin.HandlerFunc {
	return func(c *gin.Context) {
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
		persons, err := s.PersonService.FilterPersonsByPersonType(personTypeID, actionBy)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, persons)
	}
}
