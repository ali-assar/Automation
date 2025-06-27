package battalion

import (
	"backend/internal/core/battalion"
	"backend/internal/core/dataviews"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AddRoutes(r *gin.RouterGroup) {
	g := r.Group("battalion")

	addDataviews(r)

	g.POST("/create", createBattalion)
	g.POST("/update", updateBattalion)
	g.GET("/get/:id", getBattalion)
	g.DELETE("/delete/:id", deleteBattalion)
}

func addDataviews(r *gin.RouterGroup) {
	dataviews.RegisterRoute(r, battalion.GetListDataview())
}

func createBattalion(c *gin.Context) {
	ctx := c.Request.Context()

	b := battalion.Battalion{}
	if err := c.ShouldBindJSON(&b); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := battalion.Create(ctx, &b)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": id})
}

func updateBattalion(c *gin.Context) {
	ctx := c.Request.Context()

	b := battalion.Battalion{}
	if err := c.ShouldBindJSON(&b); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := battalion.Update(ctx, &b)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": id})
}

func getBattalion(c *gin.Context) {
	ctx := c.Request.Context()

	idStr := c.Param("id")

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	b, err := battalion.FetchForEdit(ctx, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, b)
}

func deleteBattalion(c *gin.Context) {
	ctx := c.Request.Context()

	idStr := c.Param("id")

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	err = battalion.Delete(ctx, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	return
}
