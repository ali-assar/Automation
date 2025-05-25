package skills

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

func (r *Repository) Create(skills *Skills) error {
	return r.db.Create(skills).Error
}

func (r *Repository) GetByID(id int64) (*Skills, error) {
	var skills Skills
	if err := r.db.Preload("Education").First(&skills, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &skills, nil
}

func (r *Repository) GetByEducationID(educationID int64) ([]Skills, error) {
	var skills []Skills
	if err := r.db.Preload("Education").Where("education_id = ? AND deleted_at = 0", educationID).Find(&skills).Error; err != nil {
		return nil, err
	}
	return skills, nil
}

func (r *Repository) GetAll() ([]Skills, error) {
	var skills []Skills
	if err := r.db.Preload("Education").Where("deleted_at = 0").Find(&skills).Error; err != nil {
		return nil, err
	}
	return skills, nil
}

func (r *Repository) Update(skills *Skills) error {
	return r.db.Save(skills).Error
}

func (r *Repository) DeleteSoft(id int64) error {
	return r.db.Model(&Skills{}).Where("id = ?", id).
		Update("deleted_at", time.Now().Unix()).Error
}

func (r *Repository) DeleteHard(id int64) error {
	return r.db.Delete(&Skills{}, "id = ?", id).Error
}
