package api

import (
	"backend/internal/core/medicalprofile"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateMedicalProfile(s *HandlerService)  gin.HandlerFunc {
	return func(c *gin.Context) {
		var req struct {
			PersonID              string `json:"person_id" binding:"required"`
			PhysicalInfoID        int64  `json:"physical_info_id" binding:"required"`
			Allergies             string `json:"allergies"`
			MedicalHistory        string `json:"medical_history"`
			Vaccinations          string `json:"vaccinations"`
			BloodTypeID           int64  `json:"blood_type_id" binding:"required"`
			PsychologicalStatusID int64  `json:"psychological_status_id" binding:"required"`
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
		profile := &medicalprofile.MedicalProfile{
			PersonID:              req.PersonID,
			PhysicalInfoID:        req.PhysicalInfoID,
			Allergies:             req.Allergies,
			MedicalHistory:        req.MedicalHistory,
			Vaccinations:          req.Vaccinations,
			BloodTypeID:           req.BloodTypeID,
			PsychologicalStatusID: req.PsychologicalStatusID,
			DeletedAt:             0,
		}
		id, err := s.MedicalProfile.CreateMedicalProfile(profile, actionBy)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, gin.H{"id": id})
	}
}

func GetMedicalProfileByID(s *HandlerService)  gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
			return
		}
		profile, err := s.MedicalProfile.GetMedicalProfileByID(id)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "medical profile not found"})
			return
		}
		c.JSON(http.StatusOK, profile)
	}
}

func GetMedicalProfileByPersonID(s *HandlerService)  gin.HandlerFunc {
	return func(c *gin.Context) {
		personID := c.Param("person_id")
		profile, err := s.MedicalProfile.GetMedicalProfileByPersonID(personID)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "medical profile not found"})
			return
		}
		c.JSON(http.StatusOK, profile)
	}
}

func GetAllMedicalProfiles(s *HandlerService)  gin.HandlerFunc {
	return func(c *gin.Context) {
		profiles, err := s.MedicalProfile.GetAllMedicalProfiles()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, profiles)
	}
}

func UpdateMedicalProfile(s *HandlerService)  gin.HandlerFunc {
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
		if err := s.MedicalProfile.UpdateMedicalProfile(id, updates, actionBy); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "medical profile updated"})
	}
}

func DeleteMedicalProfile(s *HandlerService)  gin.HandlerFunc {
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
		if err := s.MedicalProfile.DeleteMedicalProfileSoft(id, actionBy); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "medical profile soft deleted"})
	}
}

func DeleteMedicalProfileHard(s *HandlerService)  gin.HandlerFunc {
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
		if err := s.MedicalProfile.DeleteMedicalProfileHard(id, actionBy); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "medical profile hard deleted"})
	}
}
