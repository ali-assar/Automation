package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateSkills(s *HandlerService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req struct {
			EducationID       int64  `json:"education_id" binding:"required"`
			Languages         string `json:"languages" binding:"required"`
			SkillsDescription string `json:"skills_description" binding:"required"`
			Certificates      string `json:"certificates"`
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
		id, err := s.SkillsService.CreateSkills(
			req.EducationID, req.Languages, req.SkillsDescription, req.Certificates, actionBy,
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, gin.H{"id": id})
	}
}

func GetSkillsByID(s *HandlerService) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
			return
		}
		skills, err := s.SkillsService.GetSkillsByID(id)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "skills not found"})
			return
		}
		c.JSON(http.StatusOK, skills)
	}
}

func GetSkillsByEducationID(s *HandlerService) gin.HandlerFunc {
	return func(c *gin.Context) {
		educationID, err := strconv.ParseInt(c.Param("education_id"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid education_id"})
			return
		}
		skills, err := s.SkillsService.GetSkillsByEducationID(educationID)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "skills not found"})
			return
		}
		c.JSON(http.StatusOK, skills)
	}
}

func GetAllSkills(s *HandlerService) gin.HandlerFunc {
	return func(c *gin.Context) {
		skills, err := s.SkillsService.GetAllSkills()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, skills)
	}
}

func UpdateSkills(s *HandlerService) gin.HandlerFunc {
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
		if err := s.SkillsService.UpdateSkills(id, updates, actionBy); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "skills updated"})
	}
}

func DeleteSkills(s *HandlerService) gin.HandlerFunc {
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
		if err := s.SkillsService.DeleteSkills(id, actionBy); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "skills deleted"})
	}
}
