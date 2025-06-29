package userrole

import (
	"backend/internal/core/dataviews"
	"strconv"

	"backend/internal/core/userrole"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddRoutes(r *gin.RouterGroup) {
	g := r.Group("user-role")

	addDataviews(r)

	g.POST("/save", save)
	g.GET("/get-user-roles/:userId", getUserRoles)
}

func addDataviews(r *gin.RouterGroup) {
	dataviews.RegisterRoute(r, userrole.GetUserListDataview())
}

func save(c *gin.Context) {
	ctx := c.Request.Context()

	b := struct {
		UserID  int64   `json:"userId"`
		RoleIds []int64 `json:"roleIds"`
	}{}
	if err := c.ShouldBindJSON(&b); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := userrole.SaveUserRoles(ctx, b.UserID, b.RoleIds)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

func getUserRoles(c *gin.Context) {
	ctx := c.Request.Context()

	idStr := c.Param("userId")

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	roles, err := userrole.GetUserRoles(ctx, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, roles)
}
