package admin

import (
	"backend/internal/config"
	"backend/internal/core/admin"
	"backend/pkg/security"
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Controller struct {
	dbServiceURL string
}

func NewController() *Controller {
	cfg := config.Load()
	return &Controller{dbServiceURL: cfg.DBServiceUrl}
}



func (ctrl *Controller) Login(c *gin.Context) {
	var req struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resp, err := http.Get(ctrl.dbServiceURL + "/admin/username/" + req.Username)
	if err != nil || resp.StatusCode != http.StatusOK {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid username"})
		return
	}
	defer resp.Body.Close()
	var admin admin.Admin
	if err := json.NewDecoder(resp.Body).Decode(&admin); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to decode admin"})
		return
	}
	if ok, err := security.ComparePasswords(admin.HashPassword, req.Password); !ok || err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid password"})
		return
	}
	staticToken, err := security.GenerateStaticToken(admin.ID, 1) // Assuming role 1 for simplicity
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate static token"})
		return
	}
	dynamicToken, err := security.GenerateDynamicToken(admin.ID, 1)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate dynamic token"})
		return
	}
	credReq := map[string]string{
		"admin_id":      admin.ID.String(),
		"static_token":  staticToken,
		"dynamic_token": dynamicToken,
	}
	body, _ := json.Marshal(credReq)
	resp, err = http.Post(ctrl.dbServiceURL+"/credentials", "application/json", bytes.NewBuffer(body))
	if err != nil || resp.StatusCode != http.StatusCreated {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create credentials"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"admin":         admin,
		"static_token":  staticToken,
		"dynamic_token": dynamicToken,
	})
}

func (ctrl *Controller) CreateAdmin(c *gin.Context) {
	var req struct {
		NationalIDNumber string `json:"national_id_number" binding:"required"`
		UserName         string `json:"user_name" binding:"required"`
		Password         string `json:"password" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	hash, err := security.HashPassword(req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to hash password"})
		return
	}
	body, _ := json.Marshal(map[string]string{
		"national_id_number": req.NationalIDNumber,
		"user_name":          req.UserName,
		"password":           hash,
	})
	resp, err := http.Post(ctrl.dbServiceURL+"/admin", "application/json", bytes.NewBuffer(body))
	if err != nil || resp.StatusCode != http.StatusCreated {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create admin"})
		return
	}
	var admin admin.Admin
	json.NewDecoder(resp.Body).Decode(&admin)
	c.JSON(http.StatusCreated, admin)
}

func (ctrl *Controller) GetAdmin(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	resp, err := http.Get(ctrl.dbServiceURL + "/admin/" + id.String())
	if err != nil || resp.StatusCode != http.StatusOK {
		c.JSON(http.StatusNotFound, gin.H{"error": "admin not found"})
		return
	}
	defer resp.Body.Close()
	var admin admin.Admin
	json.NewDecoder(resp.Body).Decode(&admin)
	c.JSON(http.StatusOK, admin)
}

func (ctrl *Controller) GetAllAdmins(c *gin.Context) {
	resp, err := http.Get(ctrl.dbServiceURL + "/admins")
	if err != nil || resp.StatusCode != http.StatusOK {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch admins"})
		return
	}
	defer resp.Body.Close()
	var admins []admin.Admin
	json.NewDecoder(resp.Body).Decode(&admins)
	c.JSON(http.StatusOK, admins)
}
