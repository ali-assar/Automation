package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateRole(s *HandlerService) gin.HandlerFunc {
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
		id, err := s.RoleService.CreateRole(req.TypeName, actionBy)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, gin.H{"id": id})
	}
}

func GetRoleByID(s *HandlerService) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
			return
		}
		role, err := s.RoleService.GetRoleByID(id)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "role not found"})
			return
		}
		c.JSON(http.StatusOK, role)
	}
}

func GetRoleByType(s *HandlerService) gin.HandlerFunc {
	return func(c *gin.Context) {
		typeName := c.Param("type")
		role, err := s.RoleService.GetRoleByType(typeName)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "role not found"})
			return
		}
		c.JSON(http.StatusOK, role)
	}
}

func GetAllRoles(s *HandlerService) gin.HandlerFunc {
	return func(c *gin.Context) {
		roles, err := s.RoleService.GetAllRoles()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, roles)
	}
}

func UpdateRole(s *HandlerService) gin.HandlerFunc {
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
		if err := s.RoleService.UpdateRole(id, req.TypeName, actionBy); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "role updated"})
	}
}

func DeleteRole(s *HandlerService) gin.HandlerFunc {
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
		if err := s.RoleService.DeleteRole(id, actionBy); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "role deleted"})
	}
}
