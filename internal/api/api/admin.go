package api

import (
	"backend/pkg/security"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func Login(s *HandlerService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req struct {
			Username string `json:"username" binding:"required"`
			Password string `json:"password" binding:"required"`
		}
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		admin, err := s.AdminService.GetAdminByUsername(req.Username)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid username"})
			return
		}
		if ok, err := security.ComparePasswords(admin.HashPassword, req.Password); !ok || err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid password"})
			return
		}
		staticToken, err := security.GenerateStaticToken(admin.ID, int(admin.RoleID)) // Role 1 = admin
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate static token"})
			return
		}
		dynamicToken, err := security.GenerateDynamicToken(admin.ID, int(admin.RoleID))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate dynamic token"})
			return
		}

		_, err = s.CredentialsService.CreateCredentials(admin.ID, staticToken, dynamicToken, admin.UserName)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create credentials"})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"admin":         admin,
			"static_token":  staticToken,
			"dynamic_token": dynamicToken,
		})
	}
}
func CreateAdmin(s *HandlerService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req struct {
			NationalIDNumber string `json:"national_id_number" binding:"required"`
			UserName         string `json:"user_name" binding:"required"`
			Password         string `json:"password" binding:"required"`
			RoleID           int64  `json:"role_id" binding:"required"`
			CredentialsID    int64  `json:"credentials_id" binding:"required"`
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
		hash, err := security.HashPassword(req.Password)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to hash password"})
			return
		}
		id, err := s.AdminService.CreateAdmin(req.NationalIDNumber, req.UserName, hash, req.RoleID, req.CredentialsID, actionBy)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, gin.H{"id": id})
	}
}

func GetAdminByID(s *HandlerService) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := uuid.Parse(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
			return
		}
		admin, err := s.AdminService.GetAdminByID(id)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "admin not found"})
			return
		}
		c.JSON(http.StatusOK, admin)
	}
}

func GetAdminByUsername(s *HandlerService) gin.HandlerFunc {
	return func(c *gin.Context) {
		username := c.Param("username")
		admin, err := s.AdminService.GetAdminByUsername(username)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "admin not found"})
			return
		}
		c.JSON(http.StatusOK, admin)
	}
}

func GetAllAdmins(s *HandlerService) gin.HandlerFunc {
	return func(c *gin.Context) {
		admins, err := s.AdminService.GetAllAdmins()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, admins)
	}
}

func UpdateAdmin(s *HandlerService) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := uuid.Parse(c.Param("id"))
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
		if err := s.AdminService.UpdateAdmin(id, updates, actionBy); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "admin updated"})
	}
}

func UpdateAdminPassword(s *HandlerService) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := uuid.Parse(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
			return
		}
		var req struct {
			Password string `json:"password" binding:"required"`
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
		if err := s.AdminService.UpdateAdminPassword(id, req.Password, actionBy); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "password updated"})
	}
}

func DeleteAdmin(s *HandlerService) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := uuid.Parse(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
			return
		}
		actionBy := c.GetHeader("X-Action-By")
		if actionBy == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "missing X-Action-By header"})
			return
		}
		if err := s.AdminService.DeleteAdmin(id, actionBy); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "admin deleted"})
	}
}

func SoftDeleteAdmin(s *HandlerService) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := uuid.Parse(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
			return
		}
		actionBy := c.GetHeader("X-Action-By")
		if actionBy == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "missing X-Action-By header"})
			return
		}
		if err := s.AdminService.SoftDeleteAdmin(id, actionBy); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "admin soft deleted"})
	}
}
