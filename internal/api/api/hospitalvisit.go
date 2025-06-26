package api

import (
	"backend/internal/core/hospitalvisit"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateVisit creates a new hospital visit record
func CreateVisit(s *HandlerService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req struct {
			PersonID  string `json:"person_id" binding:"required"`
			Date      int64  `json:"date" binding:"required"`
			Reason    string `json:"reason"`
			Diagnosis string `json:"diagnosis"`
			Treatment string `json:"treatment"`
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
		visit := &hospitalvisit.Visit{
			PersonID:  req.PersonID,
			Date:      req.Date,
			Reason:    req.Reason,
			Diagnosis: req.Diagnosis,
			Treatment: req.Treatment,
			DeletedAt: 0,
		}
		id, err := s.HospitalVisit.CreateVisit(visit, actionBy)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, gin.H{"id": id})
	}
}

// GetVisitByID retrieves a hospital visit by its ID
func GetVisitByID(s *HandlerService) gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
			return
		}
		visit, err := s.HospitalVisit.GetVisitByID(id)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "visit not found"})
			return
		}
		c.JSON(http.StatusOK, visit)
	}
}

// GetVisitByPersonID retrieves hospital visits by person ID (national ID)
func GetVisitByPersonID(s *HandlerService) gin.HandlerFunc {
	return func(c *gin.Context) {
		personID := c.Param("person_id")
		visits, err := s.HospitalVisit.GetVisitsByPersonID(personID)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "visits not found for person"})
			return
		}
		c.JSON(http.StatusOK, visits)
	}
}

// GetAllVisits retrieves all hospital visits
func GetAllVisits(s *HandlerService) gin.HandlerFunc {
	return func(c *gin.Context) {
		visits, err := s.HospitalVisit.GetAllVisits()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, visits)
	}
}

// UpdateVisit updates a hospital visit by ID
func UpdateVisit(s *HandlerService) gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.ParseInt(idStr, 10, 64)
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
		if err := s.HospitalVisit.UpdateVisit(id, updates, actionBy); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "visit updated"})
	}
}

// DeleteVisit soft deletes a hospital visit by ID
func DeleteVisit(s *HandlerService) gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
			return
		}
		actionBy := c.GetHeader("X-Action-By")
		if actionBy == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "missing X-Action-By header"})
			return
		}
		if err := s.HospitalVisit.DeleteVisitSoft(id, actionBy); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "visit soft deleted"})
	}
}

// DeleteVisitHard permanently deletes a hospital visit by ID
func DeleteVisitHard(s *HandlerService) gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
			return
		}
		actionBy := c.GetHeader("X-Action-By")
		if actionBy == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "missing X-Action-By header"})
			return
		}
		if err := s.HospitalVisit.DeleteVisitHard(id, actionBy); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "visit hard deleted"})
	}
}
