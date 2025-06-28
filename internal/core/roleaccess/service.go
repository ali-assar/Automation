package roleaccess

import (
	"backend/internal/db"
	"context"
	"errors"

	"gorm.io/gorm"
)

func Upsert(ctx context.Context, req *RoleAccessDTO) error {

	if req.RoleID == 0 {
		return errors.New("role id is required")
	}

	db.RunInTransaction(func(tx *gorm.DB) error {
		err := tx.Exec("DELETE FROM role_access WHERE role_id = ?", req.RoleID).Error
		if err != nil {
			return err
		}

		accessToSave := []*RoleAccess{}

		for _, item := range req.Details {
			accessToSave = append(accessToSave, &RoleAccess{
				RoleID:      req.RoleID,
				ResourceKey: item.ResourceKey,
			})
		}

		err = tx.Create(&accessToSave).Error
		if err != nil {
			return err
		}

		return nil
	})

	return nil
}

func FetchAccess(ctx context.Context, roleID int64) (*RoleAccessDTO, error) {
	db := db.GetDB()

	if roleID == 0 {
		return nil, errors.New("role id is required")
	}

	accessList := []*RoleDetail{}

	var keys []string
	err := db.Model(&RoleAccess{}).Where("role_id = ?", roleID).Pluck("resource_key", &keys).Error
	if err != nil {
		return nil, err
	}

	for _, item := range keys {
		accessList = append(accessList, &RoleDetail{
			ResourceKey: item,
		})
	}

	return &RoleAccessDTO{
		RoleID:  roleID,
		Details: accessList,
	}, nil
}
