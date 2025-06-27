package role

import (
	"backend/internal/core/dataviews"
	"context"

	"gorm.io/gorm"
)

func GetRoleListsDataview() *dataviews.DataviewModel[Role, any] {
	return &dataviews.DataviewModel[Role, any]{
		Query: func(ctx context.Context, db *gorm.DB, a *any) (*gorm.DB, error) {
			return db.Table("roles").Select("*"), nil
		},
		DataviewKey: "roles_all",
	}
}
