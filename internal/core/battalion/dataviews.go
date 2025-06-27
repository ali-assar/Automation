package battalion

import (
	"backend/internal/core/dataviews"
	"context"

	"gorm.io/gorm"
)

func GetListDataview() *dataviews.DataviewModel[Battalion, any] {
	return &dataviews.DataviewModel[Battalion, any]{
		Query: func(ctx context.Context, db *gorm.DB, a *any) (*gorm.DB, error) {
			return db.Table("battalion").Select("*"), nil
		},
		DataviewKey: "battalion_all",
	}
}
