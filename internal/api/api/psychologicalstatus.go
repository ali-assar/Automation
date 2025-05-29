package api

import (
	"backend/internal/core/psychologicalstatus"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (s *HandlerService) CreatePsychologicalStatus() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req struct {
			Status string `json:"status" binding:"required"`
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
		status := &psychologicalstatus.PsychologicalStatus{
			Status: req.Status,
		}
		id, err := s.PsychologicalStatus.CreatePsychologicalStatus(status, actionBy)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, gin.H{"id": id})
	}
}

func (s *HandlerService) GetPsychologicalStatusByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
			return
		}
		status, err := s.PsychologicalStatus.GetPsychologicalStatusByID(id)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "psychological status not found"})
			return
		}
		c.JSON(http.StatusOK, status)
	}
}

func (s *HandlerService) GetAllPsychologicalStatuses() gin.HandlerFunc {
	return func(c *gin.Context) {
		statuses, err := s.PsychologicalStatus.GetAllPsychologicalStatuses()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, statuses)
	}
}

func (s *HandlerService) UpdatePsychologicalStatus() gin.HandlerFunc {
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
		if err := s.PsychologicalStatus.UpdatePsychologicalStatus(id, updates, actionBy); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "psychological status updated"})
	}
}

func (s *HandlerService) DeletePsychologicalStatus() gin.HandlerFunc {
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
		if err := s.PsychologicalStatus.DeletePsychologicalStatus(id, actionBy); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "psychological status deleted"})
	}
}
