package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateRole(s *Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req struct {
			TypeName string `json:"type_name" binding:"required"`
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
		id, err := s.roleService.CreateRole(req.TypeName, actionBy)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, gin.H{"id": id})
	}
}

func GetRoleByID(s *Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
			return
		}
		role, err := s.roleService.GetRoleByID(id)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "role not found"})
			return
		}
		c.JSON(http.StatusOK, role)
	}
}

func GetRoleByType(s *Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		typeName := c.Param("type")
		role, err := s.roleService.GetRoleByType(typeName)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "role not found"})
			return
		}
		c.JSON(http.StatusOK, role)
	}
}

func GetAllRoles(s *Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		roles, err := s.roleService.GetAllRoles()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, roles)
	}
}

func UpdateRole(s *Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
			return
		}
		var req struct {
			TypeName string `json:"type_name" binding:"required"`
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
		if err := s.roleService.UpdateRole(id, req.TypeName, actionBy); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "role updated"})
	}
}

func DeleteRole(s *Service) gin.HandlerFunc {
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
		if err := s.roleService.DeleteRole(id, actionBy); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "role deleted"})
	}
}
