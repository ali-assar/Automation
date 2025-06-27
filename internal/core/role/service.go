package role

import (
	"backend/internal/db"
	"context"

	"gorm.io/gorm"
)

func Create(ctx context.Context, role *Role) (int64, error) {
	err := db.RunInTransaction(func(tx *gorm.DB) error {
		return tx.Create(role).Error
	})
	if err != nil {
		return 0, err
	}

	return role.ID, nil
}

func Update(ctx context.Context, role *Role) (int64, error) {
	err := db.RunInTransaction(func(tx *gorm.DB) error {
		return tx.Updates(role).Error
	})
	if err != nil {
		return 0, err
	}

	return role.ID, nil
}

func FetchForEdit(ctx context.Context, id int64) (*Role, error) {
	db := db.GetDB()

	b := Role{}

	err := db.Model(Role{}).Where("id = ?", id).Scan(&b).Error
	if err != nil {
		return nil, err
	}

	return &b, nil
}

func Delete(ctx context.Context, id int64) error {
	db := db.GetDB()

	err := db.Exec("DELETE FROM roles WHERE id = ?", id).Error
	if err != nil {
		return err
	}

	return nil
}
