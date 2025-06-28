package userrole

import (
	"backend/internal/db"
	"context"

	"gorm.io/gorm"
)

func SaveUserRoles(ctx context.Context, userID int64, roleIds []int64) error {
	db.RunInTransaction(func(tx *gorm.DB) error {
		err := tx.Exec("DELETE FROM user_role WHERE user_id = ?", userID).Error
		if err != nil {
			return err
		}

		resultToSave := []*UserRole{}

		for _, roleID := range roleIds {
			resultToSave = append(resultToSave, &UserRole{
				UserID: userID,
				RoleID: roleID,
			})
		}

		if len(resultToSave) > 0 {
			err = tx.Create(&resultToSave).Error
			if err != nil {
				return err
			}
		}

		return nil
	})

	return nil
}
