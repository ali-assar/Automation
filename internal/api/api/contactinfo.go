package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateContactInfo(s *HandlerService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req struct {
			Address              string `json:"address" binding:"required"`
			EmailAddress         string `json:"email_address" binding:"required,email"`
			SocialMedia          string `json:"social_media"`
			PhoneNumber          string `json:"phone_number" binding:"required"`
			EmergencyPhoneNumber string `json:"emergency_phone_number"`
			LandlinePhone        string `json:"landline_phone"`
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
		id, err := s.ContactInfoService.CreateContactInfo(
			req.Address, req.EmailAddress, req.SocialMedia,
			req.PhoneNumber, req.EmergencyPhoneNumber, req.LandlinePhone, actionBy,
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, gin.H{"id": id})
	}
}

func UpdateContactInfo(s *HandlerService) gin.HandlerFunc {
	return func(c *gin.Context) {
		nationalID := c.Param("national_id")
		var req struct {
			ContactInfoID int64 `json:"contact_info_id" binding:"required"`
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
		if err := s.PersonService.UpdateContactInfo(nationalID, req.ContactInfoID, actionBy); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "contact info updated"})
	}
}

func GetContactInfoByID(s *HandlerService) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
			return
		}
		contactInfo, err := s.ContactInfoService.GetContactInfoByID(id)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "contact info not found"})
			return
		}
		c.JSON(http.StatusOK, contactInfo)
	}
}

func GetContactInfoByEmail(s *HandlerService) gin.HandlerFunc {
	return func(c *gin.Context) {
		email := c.Param("email")
		contactInfo, err := s.ContactInfoService.GetContactInfoByEmail(email)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "contact info not found"})
			return
		}
		c.JSON(http.StatusOK, contactInfo)
	}
}

func GetAllContactInfos(s *HandlerService) gin.HandlerFunc {
	return func(c *gin.Context) {
		contactInfos, err := s.ContactInfoService.GetAllContactInfos()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, contactInfos)
	}
}

func DeleteContactInfo(s *HandlerService) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
			return
		}
		actionBy := c.GetHeader("X-Action-By")
		if actionBy == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "missing X-Action-By header"})
			return
		}
		if err := s.ContactInfoService.DeleteContactInfo(id, actionBy); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "contact info soft deleted"})
	}
}

func DeleteContactInfoHard(s *HandlerService) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
			return
		}
		actionBy := c.GetHeader("X-Action-By")
		if actionBy == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "missing X-Action-By header"})
			return
		}
		if err := s.ContactInfoService.DeleteContactInfoHard(id, actionBy); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "contact info hard deleted"})
	}
}
