package battalion

import (
	"backend/internal/db"
	"context"

	"gorm.io/gorm"
)

func Create(ctx context.Context, battlation *Battalion) (int64, error) {
	err := db.RunInTransaction(func(tx *gorm.DB) error {
		return tx.Create(battlation).Error
	})
	if err != nil {
		return 0, err
	}

	return battlation.ID, nil
}

func Update(ctx context.Context, battlation *Battalion) (int64, error) {
	err := db.RunInTransaction(func(tx *gorm.DB) error {
		return tx.Updates(battlation).Error
	})
	if err != nil {
		return 0, err
	}

	return battlation.ID, nil
}

func FetchForEdit(ctx context.Context, id int64) (*Battalion, error) {
	db := db.GetDB()

	b := Battalion{}

	err := db.Model(Battalion{}).Where("id = ?", id).Scan(&b).Error
	if err != nil {
		return nil, err
	}

	return &b, nil
}
