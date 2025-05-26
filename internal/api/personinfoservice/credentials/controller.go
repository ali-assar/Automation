package credentials

import (
	"backend/internal/config"
	"backend/internal/core/credentials"
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

func (ctrl *Controller) GetByAdminID(adminID uuid.UUID) (*credentials.Credentials, error) {
	resp, err := http.Get(ctrl.dbServiceURL + "/credentials/" + adminID.String())
	if err != nil || resp.StatusCode != http.StatusOK {
		return nil, err
	}
	defer resp.Body.Close()
	var cred credentials.Credentials
	if err := json.NewDecoder(resp.Body).Decode(&cred); err != nil {
		return nil, err
	}
	return &cred, nil
}

func (ctrl *Controller) UpdateDynamicToken(c *gin.Context) {
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
	body, _ := json.Marshal(req)
	resp, err := http.Post(ctrl.dbServiceURL+"/credentials/"+adminID.String()+"/dynamic", "application/json", bytes.NewBuffer(body))
	if err != nil || resp.StatusCode != http.StatusOK {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update dynamic token"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "dynamic token updated"})
}
