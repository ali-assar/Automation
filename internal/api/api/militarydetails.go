package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateMilitaryDetails(s *HandlerService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req struct {
			RankID              int64  `json:"rank_id" binding:"required"`
			ServiceStartDate    int64  `json:"service_start_date" binding:"required"`
			ServiceDispatchDate int64  `json:"service_dispatch_date" binding:"required"`
			ServiceUnit         string `json:"service_unit" binding:"required"`
			BattalionUnit       string `json:"battalion_unit" binding:"required"`
			CompanyUnit         string `json:"company_unit" binding:"required"`
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
		id, err := s.MilitaryDetailsService.CreateMilitaryDetails(
			req.RankID, &req.ServiceStartDate, &req.ServiceDispatchDate,
			&req.ServiceUnit, &req.BattalionUnit, &req.CompanyUnit, actionBy,
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, gin.H{"id": id})
	}
}

func UpdateMilitaryDetails(s *HandlerService) gin.HandlerFunc {
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
		if err := s.MilitaryDetailsService.UpdateMilitaryDetails(id, updates, actionBy); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "military details updated"})
	}

}

func GetMilitaryDetailsByID(s *HandlerService) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
			return
		}
		militaryDetails, err := s.MilitaryDetailsService.GetMilitaryDetailsByID(id)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "military details not found"})
			return
		}
		c.JSON(http.StatusOK, militaryDetails)
	}
}

func GetAllMilitaryDetails(s *HandlerService) gin.HandlerFunc {
	return func(c *gin.Context) {
		militaryDetails, err := s.MilitaryDetailsService.GetAllMilitaryDetails()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, militaryDetails)
	}
}

func DeleteMilitaryDetails(s *HandlerService) gin.HandlerFunc {
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
		if err := s.MilitaryDetailsService.DeleteMilitaryDetails(id, actionBy); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "military details soft deleted"})
	}
}
