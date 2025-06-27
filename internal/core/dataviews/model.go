package dataviews

import (
	"context"

	"gorm.io/gorm"
)

type DataviewModel[K, Y any] struct {
	DataviewKey        string
	Query              func(context.Context, *gorm.DB, *Y) (*gorm.DB, error)
	ModelForPatch      K
	ModelForParameters Y
}

type dataviewBodyReq[Y any] struct {
	Parameters *Y
	PageSize   int
	PageNumber int
}