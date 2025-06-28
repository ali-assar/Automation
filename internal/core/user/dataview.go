package user

import (
	"backend/internal/core/dataviews"
	"context"

	"gorm.io/gorm"
)

func GetUserListDataview() *dataviews.DataviewModel[UserDTO, any] {
	return &dataviews.DataviewModel[UserDTO, any]{
		Query: func(ctx context.Context, db *gorm.DB, a *any) (*gorm.DB, error) {
			return db.Table("users u").
				Joins("JOIN person p ON p.national_id_number = u.national_id_number").
				Select("u.*, p.first_name, p.last_name"), nil
		},
		DataviewKey: "users_all",
	}
}
