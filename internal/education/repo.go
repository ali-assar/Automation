package education

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

func (r *Repository) Create(education *Education) error {
	return r.db.Create(education).Error
}

func (r *Repository) GetByID(id int64) (*Education, error) {
	var education Education
	if err := r.db.First(&education, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &education, nil
}

func (r *Repository) GetAll() ([]Education, error) {
	var educations []Education
	if err := r.db.Where("deleted_at = 0").Find(&educations).Error; err != nil {
		return nil, err
	}
	return educations, nil
}

func (r *Repository) Update(education *Education) error {
	return r.db.Save(education).Error
}

func (r *Repository) DeleteSoft(id int64) error {
	return r.db.Model(&Education{}).Where("id = ?", id).
		Update("deleted_at", time.Now().Unix()).Error
}

func (r *Repository) DeleteHard(id int64) error {
	return r.db.Delete(&Education{}, "id = ?", id).Error
}

func (r *Repository) SearchByUniversity(university string) ([]Education, error) {
	var educations []Education
	if err := r.db.Where("university ILIKE ? AND deleted_at = 0", "%"+university+"%").Find(&educations).Error; err != nil {
		return nil, err
	}
	return educations, nil
}
