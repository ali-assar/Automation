package user

import (
	"backend/internal/core/dataviews"
	"backend/internal/core/user"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AddRoutes(r *gin.RouterGroup) {
	g := r.Group("user")

	addDataviews(r)

	g.POST("/create", create)
	g.POST("/update", update)
	g.GET("/get/:id", get)
	g.DELETE("/delete/:id", delete)
}

func addDataviews(r *gin.RouterGroup) {
	dataviews.RegisterRoute(r, user.GetUserListDataview())
}

func create(c *gin.Context) {
	ctx := c.Request.Context()

	b := user.UserSaveReq{}
	if err := c.ShouldBindJSON(&b); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := user.Create(ctx, &b)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": id})
}

func update(c *gin.Context) {
	ctx := c.Request.Context()

	b := user.UserSaveReq{}
	if err := c.ShouldBindJSON(&b); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := user.Update(ctx, &b)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": id})
}

func get(c *gin.Context) {
	ctx := c.Request.Context()

	idStr := c.Param("id")

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	b, err := user.FetchForEdit(ctx, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, b)
}

func delete(c *gin.Context) {
	ctx := c.Request.Context()

	idStr := c.Param("id")

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	err = user.Delete(ctx, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	return
}
