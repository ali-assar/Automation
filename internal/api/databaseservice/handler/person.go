package handler

import (
	"backend/internal/core/person"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func CreatePerson(s *Service) gin.HandlerFunc {
	return func(c *gin.Context) {
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
		id, err := s.personService.CreatePerson(person, actionBy)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, gin.H{"national_id_number": id})
	}
}

func GetPersonByID(s *Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		nationalID := c.Param("national_id")
		person, err := s.personService.GetPersonByID(nationalID)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "person not found"})
			return
		}
		c.JSON(http.StatusOK, person)
	}
}

func GetAllPersons(s *Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		persons, err := s.personService.GetAllPersons()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, persons)
	}
}

func UpdatePerson(s *Service) gin.HandlerFunc {
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
		if err := s.personService.UpdatePerson(nationalID, updates, actionBy); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "person updated"})
	}
}

func UpdateContactInfo(s *Service) gin.HandlerFunc {
	return func(c *gin.Context) {
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
		if err := s.personService.UpdateContactInfo(nationalID, req.ContactInfoID, actionBy); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "contact info updated"})
	}
}

func UpdateMilitaryDetails(s *Service) gin.HandlerFunc {
	return func(c *gin.Context) {
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
		if err := s.personService.UpdateMilitaryDetails(nationalID, req.MilitaryDetailsID, actionBy); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "military details updated"})
	}
}

func DeletePerson(s *Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		nationalID := c.Param("national_id")
		actionBy := c.GetHeader("X-Action-By")
		if actionBy == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "missing X-Action-By header"})
			return
		}
		if err := s.personService.DeletePerson(nationalID, actionBy); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "person soft deleted"})
	}
}

func DeletePersonHard(s *Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		nationalID := c.Param("national_id")
		actionBy := c.GetHeader("X-Action-By")
		if actionBy == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "missing X-Action-By header"})
			return
		}
		if err := s.personService.DeletePersonHard(nationalID, actionBy); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "person hard deleted"})
	}
}

func SearchPersonsByName(s *Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		firstName := c.Query("first_name")
		lastName := c.Query("last_name")
		actionBy := c.GetHeader("X-Action-By")
		if actionBy == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "missing X-Action-By header"})
			return
		}
		persons, err := s.personService.SearchPersonsByName(firstName, lastName, actionBy)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, persons)
	}
}

func FilterPersonsByPersonType(s *Service) gin.HandlerFunc {
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
		persons, err := s.personService.FilterPersonsByPersonType(personTypeID, actionBy)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, persons)
	}
}
