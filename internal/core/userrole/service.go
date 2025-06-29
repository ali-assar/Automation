package userrole

import (
	"backend/internal/db"
	"context"

	"gorm.io/gorm"
)

func SaveUserRoles(ctx context.Context, userID int64, roleIds []int64) error {
	err := db.RunInTransaction(func(tx *gorm.DB) error {
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
	if err != nil {
		return err
	}

	return nil
}

func GetUserRoles(ctx context.Context, userID int64) ([]*UserRoleDTO, error) {
	db := db.GetDB()

	var result = []*UserRoleDTO{}

	err := db.Table("user_role ur").
		Joins("RIGHT JOIN roles r ON r.id = ur.role_id").
		Where("(CASE WHEN ur.id IS NULL THEN TRUE ELSE ur.user_id = ? END)", userID).
		Select(
			"r.id AS role_id",
			"r.title AS role_title",
			"(CASE WHEN ur.id IS NULL THEN false ELSE true END) AS checked",
		).
		Order("role_id").
		Scan(&result).Error
	if err != nil {
		return nil, err
	}

	return result, nil
}
