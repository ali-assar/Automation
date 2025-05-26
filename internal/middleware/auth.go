package middleware

import (
	"backend/internal/core/admin"
	"backend/internal/core/credentials"
	"backend/internal/logger"
	"backend/internal/response"
	"backend/pkg/security"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const (
	AuthorizationHeader = "Authorization"
	RoleParam           = "role"
	adminIDParam        = "admin_id"
)

type authMeta struct {
	token   string
	adminID uuid.UUID
	role    int
}

type AuthMiddleware struct {
	adminService       *admin.Service
	credentialsService *credentials.Service
}

func NewStaticProtectedRoute(adminService *admin.Service, credentialsService *credentials.Service) gin.HandlerFunc {
	m := &AuthMiddleware{
		adminService:       adminService,
		credentialsService: credentialsService,
	}
	return func(c *gin.Context) {
		meta, err := m.getRequestInfo(c)
		if err != nil {
			response.RespondWithError(c, http.StatusLocked, err.Error())
			return
		}

		if err := m.checkStaticWithDB(meta); err != nil {
			response.RespondWithError(c, http.StatusLocked, err.Error())
			return
		}

		if err := security.VerifyStaticToken(meta.token, meta.adminID, meta.role); err != nil {
			response.RespondWithError(c, http.StatusLocked, err.Error())
			return
		}

		if c.Request.URL.Path == "/check-static" {
			response.RespondWithSuccess(c, http.StatusAccepted, "Welcome back :)", nil, nil)
			return
		}

		c.Set("adminID", meta.adminID.String())
		c.Set("role", meta.role)
		c.Set("user_id", meta.adminID.String()) // For audit logging

		c.Next()
	}
}

func NewDynamicProtectedRoute(adminService *admin.Service, credentialsService *credentials.Service) gin.HandlerFunc {
	m := &AuthMiddleware{
		adminService:       adminService,
		credentialsService: credentialsService,
	}
	return func(c *gin.Context) {
		meta, err := m.getRequestInfo(c)
		if err != nil {
			response.RespondWithError(c, http.StatusLocked, err.Error())
			return
		}

		if err := m.checkDynamicWithDB(meta); err != nil {
			response.RespondWithError(c, http.StatusLocked, err.Error())
			return
		}

		if err := security.VerifyDynamicToken(meta.token, meta.adminID, meta.role); err != nil {
			response.RespondWithError(c, http.StatusLocked, err.Error())
			return
		}

		if c.Request.URL.Path == "/check-dynamic" {
			response.RespondWithSuccess(c, http.StatusAccepted, "Welcome back :)", nil, nil)
			return
		}

		c.Set("adminID", meta.adminID.String())
		c.Set("role", meta.role)
		c.Set("user_id", meta.adminID.String()) // For audit logging

		c.Next()
	}
}

func (m *AuthMiddleware) getRequestInfo(c *gin.Context) (*authMeta, error) {
	token := c.GetHeader(AuthorizationHeader)
	if token == "" {
		return nil, fmt.Errorf("missing token")
	}

	roleParam := c.Param(RoleParam)
	if roleParam == "" {
		return nil, fmt.Errorf("missing role")
	}

	role, err := strconv.Atoi(roleParam)
	if err != nil {
		return nil, fmt.Errorf("invalid role: %w", err)
	}

	if role > 5 {
		return nil, fmt.Errorf("invalid role: role ID exceeds maximum allowed")
	}

	adminID := c.Param(adminIDParam)
	if adminID == "" {
		return nil, fmt.Errorf("missing admin id")
	}

	adminUUID, err := uuid.Parse(adminID)
	if err != nil {
		return nil, fmt.Errorf("invalid admin id: %w", err)
	}

	admin, err := m.adminService.GetAdminByID(adminUUID)
	if err != nil {
		return nil, fmt.Errorf("failed to find admin: %w", err)
	}

	logger.Info(fmt.Sprintf("admin role ID: %d, param role: %d", admin.Role.ID, role))
	if admin.Role.ID != int64(role) {
		return nil, fmt.Errorf("role mismatch: admin has role ID %d, but %d was provided", admin.Role.ID, role)
	}

	return &authMeta{
		token:   token,
		adminID: admin.ID,
		role:    role,
	}, nil
}

func (m *AuthMiddleware) checkStaticWithDB(meta *authMeta) error {
	cred, err := m.credentialsService.GetStaticTokenByAdminID(meta.adminID)
	if err != nil {
		return fmt.Errorf("failed to get static token: %w", err)
	}

	if cred != meta.token {
		return fmt.Errorf("invalid static token")
	}

	return nil
}

func (m *AuthMiddleware) checkDynamicWithDB(meta *authMeta) error {
	cred, err := m.credentialsService.GetDynamicTokenByAdminID(meta.adminID)
	if err != nil {
		return fmt.Errorf("failed to get dynamic token: %w", err)
	}

	if cred != meta.token {
		return fmt.Errorf("invalid dynamic token")
	}

	return nil
}
