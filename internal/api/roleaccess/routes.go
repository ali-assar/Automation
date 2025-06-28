package roleaccess

import (
	"backend/internal/core/roleaccess"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AddRoutes(r *gin.RouterGroup) {
	g := r.Group("/role-access")

	g.GET("/get/:roleid", get)
	g.POST("/save", save)

}

func save(c *gin.Context) {
	ctx := c.Request.Context()

	req := roleaccess.RoleAccessDTO{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := roleaccess.Upsert(ctx, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

}

func get(c *gin.Context) {
	ctx := c.Request.Context()

	idStr := c.Param("roleid")

	roleId, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	b, err := roleaccess.FetchAccess(ctx, roleId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, b)
}
