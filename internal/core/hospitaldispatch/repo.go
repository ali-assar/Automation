package hospitaldispatch

import (
	"time"

	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Create(dispatch *HospitalDispatch) error {
	return r.db.Create(dispatch).Error
}

func (r *Repository) GetByID(id int64) (*HospitalDispatch, error) {
	var dispatch HospitalDispatch
	if err := r.db.Preload("Visit").First(&dispatch, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &dispatch, nil
}

func (r *Repository) GetAll() ([]HospitalDispatch, error) {
	var dispatches []HospitalDispatch
	if err := r.db.Preload("Visit").Where("deleted_at = 0").Find(&dispatches).Error; err != nil {
		return nil, err
	}
	return dispatches, nil
}

func (r *Repository) Update(dispatch *HospitalDispatch) error {
	return r.db.Save(dispatch).Error
}

func (r *Repository) DeleteSoft(id int64) error {
	return r.db.Model(&HospitalDispatch{}).Where("id = ?", id).
		Update("deleted_at", time.Now().Unix()).Error
}

func (r *Repository) DeleteHard(id int64) error {
	return r.db.Where("id = ?", id).Delete(&HospitalDispatch{}).Error
}
