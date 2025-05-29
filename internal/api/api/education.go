package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateEducation(s *HandlerService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req struct {
			EducationLevelID int64  `json:"education_level_id" binding:"required"`
			FieldOfStudy     int64  `json:"field_of_study" binding:"required"`
			Description      string `json:"description"`
			University       string `json:"university" binding:"required"`
			StartDate        int64  `json:"start_date" binding:"required"`
			EndDate          int64  `json:"end_date" binding:"required"`
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
		id, err := s.EducationService.CreateEducation(
			req.EducationLevelID, req.Description,
			req.University, req.StartDate, req.EndDate, actionBy,
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, gin.H{"id": id})
	}
}

func GetEducationByID(s *HandlerService) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
			return
		}
		education, err := s.EducationService.GetEducationByID(id)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "education not found"})
			return
		}
		c.JSON(http.StatusOK, education)
	}
}

func GetAllEducations(s *HandlerService) gin.HandlerFunc {
	return func(c *gin.Context) {
		educations, err := s.EducationService.GetAllEducations()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, educations)
	}
}

func UpdateEducation(s *HandlerService) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
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
		if err := s.EducationService.UpdateEducation(id, updates, actionBy); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "education updated"})
	}
}

func DeleteEducation(s *HandlerService) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
			return
		}
		actionBy := c.GetHeader("X-Action-By")
		if actionBy == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "missing X-Action-By header"})
			return
		}
		if err := s.EducationService.DeleteEducation(id, actionBy); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "education soft deleted"})
	}
}

func SearchEducationsByUniversity(s *HandlerService) gin.HandlerFunc {
	return func(c *gin.Context) {
		university := c.Query("university")
		actionBy := c.GetHeader("X-Action-By")
		if actionBy == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "missing X-Action-By header"})
			return
		}
		educations, err := s.EducationService.SearchEducationsByUniversity(university, actionBy)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, educations)
	}
}
