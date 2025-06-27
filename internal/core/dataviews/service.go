package dataviews

import (
	"backend/internal/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRoute[K, Y any](routerGroup *gin.RouterGroup, info *DataviewModel[K, Y]) {
	g := routerGroup.Group("/dataviews")

	g.POST(info.DataviewKey, func(ctx *gin.Context) {
		c := ctx.Request.Context()

		db := db.GetDB()

		var params dataviewBodyReq[Y]

		if err := ctx.ShouldBindJSON(&params); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		q, err := info.Query(c, db, params.Parameters)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if params.PageNumber < 1 {
			params.PageNumber = 1
		}

		if params.PageSize < 1 {
			params.PageSize = 10
		}

		offset := (params.PageNumber - 1) * params.PageSize

		result := []*K{}

		err = q.Offset(offset).Limit(params.PageSize).Scan(&result).Error
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		ctx.JSON(http.StatusOK, result)
	})
}