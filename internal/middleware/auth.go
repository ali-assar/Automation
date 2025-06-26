package middleware

import (
	"backend/internal/api/api"
	"backend/pkg/security"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func StaticAuth(s *api.HandlerService) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
			c.Abort()
			return
		}
		claims, err := security.ValidateToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token: " + err.Error()})
			c.Abort()
			return
		}
		adminID, err := uuid.Parse(claims["admin_id"].(string))
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid admin_id in token"})
			c.Abort()
			return
		}
		cred, err := s.CredentialsService.GetCredentialsByAdminID(adminID)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Credentials not found for admin_id"})
			c.Abort()
			return
		}
		if cred.StaticToken != token {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Static token mismatch"})
			c.Abort()
			return
		}
		c.Set("adminID", adminID.String())
		c.Set("role", int(claims["role"].(float64)))
		c.Next()
	}
}

func DynamicAuth(s *api.HandlerService) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
			c.Abort()
			return
		}
		claims, err := security.ValidateToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token: " + err.Error()})
			c.Abort()
			return
		}
		adminID, err := uuid.Parse(claims["admin_id"].(string))
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid admin_id in token"})
			c.Abort()
			return
		}
		cred, err := s.CredentialsService.GetCredentialsByAdminID(adminID)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Credentials not found for admin_id"})
			c.Abort()
			return
		}
		if cred.DynamicToken != token {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Dynamic token mismatch"})
			c.Abort()
			return
		}
		role := int(claims["role"].(float64))
		if c.Request.Method != "GET" && role != 1 {
			c.JSON(http.StatusForbidden, gin.H{"error": "Admin role (ID 1) required for write operations"})
			c.Abort()
			return
		}
		c.Set("adminID", adminID.String())
		c.Set("role", role)
		c.Next()
	}
}
