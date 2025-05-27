package hospitalvisit

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

func (r *Repository) Create(visit *Visit) error {
	return r.db.Create(visit).Error
}

func (r *Repository) GetByID(id int64) (*Visit, error) {
	var visit Visit
	if err := r.db.Preload("Person").First(&visit, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &visit, nil
}

func (r *Repository) GetAll() ([]Visit, error) {
	var visits []Visit
	if err := r.db.Preload("Person").Where("deleted_at = 0").Find(&visits).Error; err != nil {
		return nil, err
	}
	return visits, nil
}

func (r *Repository) Update(visit *Visit) error {
	return r.db.Save(visit).Error
}

func (r *Repository) DeleteSoft(id int64) error {
	return r.db.Model(&Visit{}).Where("id = ?", id).
		Update("deleted_at", time.Now().Unix()).Error
}

func (r *Repository) DeleteHard(id int64) error {
	return r.db.Where("id = ?", id).Delete(&Visit{}).Error
}
