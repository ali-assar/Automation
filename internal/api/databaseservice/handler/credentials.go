package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CreateCredentials(s *Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req struct {
			AdminID      string `json:"admin_id" binding:"required"`
			StaticToken  string `json:"static_token" binding:"required"`
			DynamicToken string `json:"dynamic_token" binding:"required"`
		}
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		adminID, err := uuid.Parse(req.AdminID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid admin_id"})
			return
		}
		actionBy := c.GetHeader("X-Action-By")
		if actionBy == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "missing X-Action-By header"})
			return
		}
		cred, err := s.credentialsService.CreateCredentials(adminID, req.StaticToken, req.DynamicToken, actionBy)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, cred)
	}
}

func GetCredentialsByAdminID(s *Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		adminID, err := uuid.Parse(c.Param("admin_id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid admin_id"})
			return
		}
		cred, err := s.credentialsService.GetCredentialsByAdminID(adminID)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "credentials not found"})
			return
		}
		c.JSON(http.StatusOK, cred)
	}
}

func GetAllCredentials(s *Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		creds, err := s.credentialsService.GetAllCredentials()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, creds)
	}
}

func GetSoftDeletedCredentials(s *Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		creds, err := s.credentialsService.GetSoftDeletedCredentials()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, creds)
	}
}

func UpdateCredentials(s *Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		adminID, err := uuid.Parse(c.Param("admin_id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid admin_id"})
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
		if err := s.credentialsService.UpdateCredentials(adminID, updates, actionBy); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "credentials updated"})
	}
}

func UpdateDynamicToken(s *Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		adminID, err := uuid.Parse(c.Param("admin_id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid admin_id"})
			return
		}
		var req struct {
			DynamicToken string `json:"dynamic_token" binding:"required"`
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
		if err := s.credentialsService.UpdateDynamicTokenByAdminID(adminID, req.DynamicToken, actionBy); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "dynamic token updated"})
	}
}

func DeleteCredentials(s *Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		adminID, err := uuid.Parse(c.Param("admin_id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid admin_id"})
			return
		}
		actionBy := c.GetHeader("X-Action-By")
		if actionBy == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "missing X-Action-By header"})
			return
		}
		if err := s.credentialsService.DeleteCredentials(adminID, actionBy); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "credentials soft deleted"})
	}
}

func DeleteCredentialsHard(s *Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		adminID, err := uuid.Parse(c.Param("admin_id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid admin_id"})
			return
		}
		actionBy := c.GetHeader("X-Action-By")
		if actionBy == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "missing X-Action-By header"})
			return
		}
		if err := s.credentialsService.DeleteCredentialsHard(adminID, actionBy); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "credentials hard deleted"})
	}
}

func GetStaticTokenByAdminID(s *Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		adminID, err := uuid.Parse(c.Param("admin_id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid admin_id"})
			return
		}
		token, err := s.credentialsService.GetStaticTokenByAdminID(adminID)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "static token not found"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"static_token": token})
	}
}

func GetDynamicTokenByAdminID(s *Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		adminID, err := uuid.Parse(c.Param("admin_id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid admin_id"})
			return
		}
		token, err := s.credentialsService.GetDynamicTokenByAdminID(adminID)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "dynamic token not found"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"dynamic_token": token})
	}
}
