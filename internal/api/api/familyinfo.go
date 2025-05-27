package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateFamilyInfo(s *HandlerService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req struct {
			FatherDetails  string `json:"father_details" binding:"required"`
			MotherDetails  string `json:"mother_details" binding:"required"`
			ChildsDetails  string `json:"childs_details"`
			HusbandDetails string `json:"husband_details"`
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
		id, err := s.FamilyInfoService.CreateFamilyInfo(
			req.FatherDetails, req.MotherDetails, req.ChildsDetails, req.HusbandDetails, actionBy,
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, gin.H{"id": id})
	}
}

func GetFamilyInfoByID(s *HandlerService) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
			return
		}
		familyInfo, err := s.FamilyInfoService.GetFamilyInfoByID(id)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "family info not found"})
			return
		}
		c.JSON(http.StatusOK, familyInfo)
	}
}

func GetAllFamilyInfos(s *HandlerService) gin.HandlerFunc {
	return func(c *gin.Context) {
		familyInfos, err := s.FamilyInfoService.GetAllFamilyInfos()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, familyInfos)
	}
}

func UpdateFamilyInfo(s *HandlerService) gin.HandlerFunc {
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
		if err := s.FamilyInfoService.UpdateFamilyInfo(id, updates, actionBy); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "family info updated"})
	}
}

func DeleteFamilyInfo(s *HandlerService) gin.HandlerFunc {
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
		if err := s.FamilyInfoService.DeleteFamilyInfo(id, actionBy); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "family info soft deleted"})
	}
}
