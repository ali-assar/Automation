package medicines

import (
	"gorm.io/gorm"
	"time"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Create(medicine *Medicine) error {
	return r.db.Create(medicine).Error
}

func (r *Repository) GetByID(id int64) (*Medicine, error) {
	var medicine Medicine
	if err := r.db.First(&medicine, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &medicine, nil
}

func (r *Repository) GetAll() ([]Medicine, error) {
	var medicines []Medicine
	if err := r.db.Where("deleted_at = 0").Find(&medicines).Error; err != nil {
		return nil, err
	}
	return medicines, nil
}

func (r *Repository) Update(medicine *Medicine) error {
	return r.db.Save(medicine).Error
}

func (r *Repository) DeleteSoft(id int64) error {
	return r.db.Model(&Medicine{}).Where("id = ?", id).
		Update("deleted_at", time.Now().Unix()).Error
}

func (r *Repository) DeleteHard(id int64) error {
	return r.db.Where("id = ?", id).Delete(&Medicine{}).Error
}