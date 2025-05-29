package api

import (
	prescriptions "backend/internal/core/prescription"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (s *HandlerService) CreatePrescription() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req struct {
			VisitID    int64  `json:"visit_id" binding:"required"`
			MedicineID int64  `json:"medicine_id" binding:"required"`
			Dose       string `json:"dose" binding:"required"`
			Duration   string `json:"duration" binding:"required"`
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
		prescription := &prescriptions.Prescription{
			VisitID:    req.VisitID,
			MedicineID: req.MedicineID,
			Dose:       req.Dose,
			Duration:   req.Duration,
			DeletedAt:  0,
		}
		id, err := s.Prescription.CreatePrescription(prescription, actionBy)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, gin.H{"id": id})
	}
}

func (s *HandlerService) GetPrescriptionByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
			return
		}
		prescription, err := s.Prescription.GetPrescriptionByID(id)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "prescription not found"})
			return
		}
		c.JSON(http.StatusOK, prescription)
	}
}

func (s *HandlerService) GetAllPrescriptions() gin.HandlerFunc {
	return func(c *gin.Context) {
		prescriptions, err := s.Prescription.GetAllPrescriptions()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, prescriptions)
	}
}

func (s *HandlerService) UpdatePrescription() gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.ParseInt(idStr, 10, 64)
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
		if err := s.Prescription.UpdatePrescription(id, updates, actionBy); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "prescription updated"})
	}
}

func (s *HandlerService) DeletePrescription() gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
			return
		}
		actionBy := c.GetHeader("X-Action-By")
		if actionBy == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "missing X-Action-By header"})
			return
		}
		if err := s.Prescription.DeletePrescriptionSoft(id, actionBy); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "prescription soft deleted"})
	}
}

func (s *HandlerService) DeletePrescriptionHard() gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
			return
		}
		actionBy := c.GetHeader("X-Action-By")
		if actionBy == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "missing X-Action-By header"})
			return
		}
		if err := s.Prescription.DeletePrescriptionHard(id, actionBy); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "prescription hard deleted"})
	}
}
