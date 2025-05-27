package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreatePhysicalInfo(s *HandlerService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req struct {
			Height           int    `json:"height" binding:"required"`
			Weight           int    `json:"weight" binding:"required"`
			EyeColor         string `json:"eye_color" binding:"required"`
			BloodGroupID     int64  `json:"blood_group_id" binding:"required"`
			GenderID         int64  `json:"gender_id" binding:"required"`
			PhysicalStatusID int64  `json:"physical_status_id" binding:"required"`
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
		id, err := s.PhysicalInfoService.CreatePhysicalInfo(
			req.Height, req.Weight, req.EyeColor,
			req.BloodGroupID, req.GenderID, req.PhysicalStatusID, actionBy,
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, gin.H{"id": id})
	}
}

func GetPhysicalInfoByID(s *HandlerService) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
			return
		}
		physicalInfo, err := s.PhysicalInfoService.GetPhysicalInfoByID(id)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "physical info not found"})
			return
		}
		c.JSON(http.StatusOK, physicalInfo)
	}
}

func GetAllPhysicalInfos(s *HandlerService) gin.HandlerFunc {
	return func(c *gin.Context) {
		physicalInfos, err := s.PhysicalInfoService.GetAllPhysicalInfos()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, physicalInfos)
	}
}

func UpdatePhysicalInfo(s *HandlerService) gin.HandlerFunc {
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
		if err := s.PhysicalInfoService.UpdatePhysicalInfo(id, updates, actionBy); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "physical info updated"})
	}
}

func DeletePhysicalInfo(s *HandlerService) gin.HandlerFunc {
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
		if err := s.PhysicalInfoService.DeletePhysicalInfo(id, actionBy); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "physical info soft deleted"})
	}
}
