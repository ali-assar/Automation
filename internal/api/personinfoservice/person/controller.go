package person

import (
	"backend/internal/config"
	"backend/internal/core/person"
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	dbServiceURL string
}

func NewController() *Controller {
	cfg := config.Load()
	return &Controller{dbServiceURL: cfg.DBServiceUrl}
}

func (ctrl *Controller) CreatePerson(c *gin.Context) {
	var req struct {
		NationalIDNumber string `json:"national_id_number" binding:"required"`
		FirstName        string `json:"first_name" binding:"required"`
		LastName         string `json:"last_name" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	body, _ := json.Marshal(req)
	resp, err := http.Post(ctrl.dbServiceURL+"/person", "application/json", bytes.NewBuffer(body))
	if err != nil || resp.StatusCode != http.StatusCreated {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create person"})
		return
	}
	var person person.Person
	json.NewDecoder(resp.Body).Decode(&person)
	c.JSON(http.StatusCreated, person)
}

func (ctrl *Controller) GetPerson(c *gin.Context) {
	nationalID := c.Param("national_id")
	resp, err := http.Get(ctrl.dbServiceURL + "/person/" + nationalID)
	if err != nil || resp.StatusCode != http.StatusOK {
		c.JSON(http.StatusNotFound, gin.H{"error": "person not found"})
		return
	}
	defer resp.Body.Close()
	var person person.Person
	json.NewDecoder(resp.Body).Decode(&person)
	c.JSON(http.StatusOK, person)
}

func (ctrl *Controller) GetAllPersons(c *gin.Context) {
	resp, err := http.Get(ctrl.dbServiceURL + "/persons")
	if err != nil || resp.StatusCode != http.StatusOK {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch persons"})
		return
	}
	defer resp.Body.Close()
	var persons []person.Person
	json.NewDecoder(resp.Body).Decode(&persons)
	c.JSON(http.StatusOK, persons)
}

func (ctrl *Controller) UpdatePerson(c *gin.Context) {
	nationalID := c.Param("national_id")
	var updates map[string]interface{}
	if err := c.ShouldBindJSON(&updates); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	body, _ := json.Marshal(updates)
	req, _ := http.NewRequest(http.MethodPut, ctrl.dbServiceURL+"/person/"+nationalID, bytes.NewBuffer(body))
	resp, err := http.DefaultClient.Do(req)
	if err != nil || resp.StatusCode != http.StatusOK {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update person"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "person updated"})
}

func (ctrl *Controller) DeletePerson(c *gin.Context) {
	nationalID := c.Param("national_id")
	req, _ := http.NewRequest(http.MethodDelete, ctrl.dbServiceURL+"/person/"+nationalID, nil)
	resp, err := http.DefaultClient.Do(req)
	if err != nil || resp.StatusCode != http.StatusOK {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete person"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "person deleted"})
}
