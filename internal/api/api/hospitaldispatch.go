package api

import (
	"backend/internal/core/hospitaldispatch"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (s *HandlerService) CreateHospitalDispatch() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req struct {
			VisitID       int64  `json:"visit_id" binding:"required"`
			DispatchDate  int64  `json:"dispatch_date" binding:"required"`
			DoctorComment string `json:"doctor_comment"`
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

		dispatchInfo := &hospitaldispatch.HospitalDispatch{
			DispatchDate:  req.DispatchDate,
			DoctorComment: req.DoctorComment,
			VisitID:       req.VisitID,
		}

		id, err := s.HospitalDispatch.CreateHospitalDispatch(dispatchInfo, actionBy)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, gin.H{"id": id})
	}
}

func (s *HandlerService) GetHospitalDispatchByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
			return
		}
		dispatch, err := s.HospitalDispatch.GetHospitalDispatchByID(id)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "hospital dispatch not found"})
			return
		}
		c.JSON(http.StatusOK, dispatch)
	}
}

func (s *HandlerService) GetAllHospitalDispatches() gin.HandlerFunc {
	return func(c *gin.Context) {
		dispatches, err := s.HospitalDispatch.GetAllHospitalDispatches()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, dispatches)
	}
}

func (s *HandlerService) UpdateHospitalDispatch() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
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
		if err := s.HospitalDispatch.UpdateHospitalDispatch(id, updates, actionBy); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "hospital dispatch updated"})
	}
}

func (s *HandlerService) DeleteHospitalDispatch() gin.HandlerFunc {
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
		if err := s.HospitalDispatch.DeleteHospitalDispatchSoft(id, actionBy); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "hospital dispatch soft deleted"})
	}
}

func (s *HandlerService) DeleteHospitalDispatchHard() gin.HandlerFunc {
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
		if err := s.HospitalDispatch.DeleteHospitalDispatchHard(id, actionBy); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "hospital dispatch hard deleted"})
	}
}
