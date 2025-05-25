package bloodgroup

import (
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Create(bloodGroup *BloodGroup) error {
	return r.db.Create(bloodGroup).Error
}

func (r *Repository) GetByID(id int64) (*BloodGroup, error) {
	var bloodGroup BloodGroup
	if err := r.db.First(&bloodGroup, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &bloodGroup, nil
}

func (r *Repository) GetByGroup(group string) (*BloodGroup, error) {
	var bloodGroup BloodGroup
	if err := r.db.First(&bloodGroup, "group = ?", group).Error; err != nil {
		return nil, err
	}
	return &bloodGroup, nil
}

func (r *Repository) GetAll() ([]BloodGroup, error) {
	var bloodGroups []BloodGroup
	if err := r.db.Find(&bloodGroups).Error; err != nil {
		return nil, err
	}
	return bloodGroups, nil
}

func (r *Repository) Update(bloodGroup *BloodGroup) error {
	return r.db.Save(bloodGroup).Error
}

func (r *Repository) Delete(id int64) error {
	return r.db.Delete(&BloodGroup{}, "id = ?", id).Error
}
