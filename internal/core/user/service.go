package user

import (
	"backend/internal/db"
	"backend/pkg/security"
	"context"
	"errors"

	"gorm.io/gorm"
)

func validateSaveReq(u *UserSaveReq) error {
	if u.Password == "" {
		return errors.New("password is empty")
	}

	if u.Username == "" {
		return errors.New("username is empty")
	}

	if u.NationalIDNumber == "" {
		return errors.New("national id number is empty")
	}

	return nil
}

func Create(ctx context.Context, u *UserSaveReq) (int64, error) {
	err := validateSaveReq(u)
	if err != nil {
		return 0, err
	}

	hashPassword, err := security.HashPassword(u.Password)
	if err != nil {
		return 0, err
	}

	err = db.RunInTransaction(func(tx *gorm.DB) error {
		userForInsert := User{
			IsAdmin:          u.IsAdmin,
			NationalIDNumber: u.NationalIDNumber,
			Username:         u.Username,
			HashPassword:     hashPassword,
		}

		err = tx.Create(&userForInsert).Error
		if err != nil {
			return err
		}

		u.ID = userForInsert.ID

		return nil
	})
	if err != nil {
		return 0, err
	}

	return u.ID, nil
}

func Update(ctx context.Context, u *UserSaveReq) (int64, error) {
	err := validateSaveReq(u)
	if err != nil {
		return 0, err
	}

	if u.ID == 0 {
		return 0, errors.New("no id found for update")
	}

	hashPassword, err := security.HashPassword(u.Password)
	if err != nil {
		return 0, err
	}

	err = db.RunInTransaction(func(tx *gorm.DB) error {
		userForUpdate := User{
			ID:               u.ID,
			IsAdmin:          u.IsAdmin,
			NationalIDNumber: u.NationalIDNumber,
			Username:         u.Username,
			HashPassword:     hashPassword,
		}

		err = tx.Save(&userForUpdate).Error
		if err != nil {
			return err
		}

		u.ID = userForUpdate.ID

		return nil
	})
	if err != nil {
		return 0, err
	}

	return u.ID, nil
}

func FetchForEdit(ctx context.Context, id int64) (*UserDTO, error) {
	db := db.GetDB()

	b := UserDTO{}

	err := db.Table("users u").
		Joins("JOIN person p ON p.national_id_number = u.national_id_number").
		Select(`
			u.*, 
			p.first_name, 
			p.last_name`,
		).
		Where("id = ?", id).Scan(&b).Error
	if err != nil {
		return nil, err
	}

	return &b, nil
}

func Delete(ctx context.Context, id int64) error {
	db := db.GetDB()

	err := db.Exec("DELETE FROM users WHERE id = ?", id).Error
	if err != nil {
		return err
	}

	return nil
}
