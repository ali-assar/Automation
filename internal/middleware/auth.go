// internal/middleware/auth.go
package middleware

import (
	"backend/internal/api/personinfoservice/admin"
	"backend/internal/api/personinfoservice/credentials"
	"backend/pkg/security"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func StaticAuth(adminCtrl *admin.Controller, credCtrl *credentials.Controller) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "missing token"})
			c.Abort()
			return
		}
		claims, err := security.ValidateToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			c.Abort()
			return
		}
		adminID, err := uuid.Parse(claims["admin_id"].(string))
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid admin_id"})
			c.Abort()
			return
		}
		cred, err := credCtrl.GetByAdminID(adminID)
		if err != nil || cred.StaticToken != token {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid static token"})
			c.Abort()
			return
		}
		c.Set("adminID", adminID.String())
		c.Set("role", int(claims["role"].(float64)))
		c.Next()
	}
}

func DynamicAuth(adminCtrl *admin.Controller, credCtrl *credentials.Controller) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "missing token"})
			c.Abort()
			return
		}
		claims, err := security.ValidateToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			c.Abort()
			return
		}
		adminID, err := uuid.Parse(claims["admin_id"].(string))
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid admin_id"})
			c.Abort()
			return
		}
		cred, err := credCtrl.GetByAdminID(adminID)
		if err != nil || cred.DynamicToken != token {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid dynamic token"})
			c.Abort()
			return
		}
		c.Set("adminID", adminID.String())
		c.Set("role", int(claims["role"].(float64)))
		c.Next()
	}
}
