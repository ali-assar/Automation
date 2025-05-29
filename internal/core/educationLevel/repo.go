package educationlevel

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

func (r *Repository) Create(educationLevel *EducationLevel) error {
	return r.db.Create(educationLevel).Error
}

func (r *Repository) GetByID(id int64) (*EducationLevel, error) {
	var educationLevel EducationLevel
	if err := r.db.First(&educationLevel, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &educationLevel, nil
}

func (r *Repository) GetAll() ([]EducationLevel, error) {
	var educationLevels []EducationLevel
	if err := r.db.Find(&educationLevels).Error; err != nil {
		return nil, err
	}
	return educationLevels, nil
}

func (r *Repository) Update(educationLevel *EducationLevel) error {
	return r.db.Save(educationLevel).Error
}

func (r *Repository) DeleteSoft(id int64) error {
	return r.db.Model(&EducationLevel{}).Where("id = ?", id).
		Update("deleted_at", time.Now().Unix()).Error
}

func (r *Repository) DeleteHard(id int64) error {
	return r.db.Delete(&EducationLevel{}, "id = ?", id).Error
}
