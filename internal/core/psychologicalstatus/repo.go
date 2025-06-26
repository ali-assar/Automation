package psychologicalstatus

import (
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Create(status *PsychologicalStatus) error {
	return r.db.Create(status).Error
}

func (r *Repository) GetByID(id int64) (*PsychologicalStatus, error) {
	var status PsychologicalStatus
	if err := r.db.First(&status, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &status, nil
}

func (r *Repository) GetAll() ([]PsychologicalStatus, error) {
	var statuses []PsychologicalStatus
	if err := r.db.Find(&statuses).Error; err != nil {
		return nil, err
	}
	return statuses, nil
}

func (r *Repository) Update(status *PsychologicalStatus) error {
	return r.db.Save(status).Error
}

func (r *Repository) Delete(id int64) error {
	return r.db.Where("id = ?", id).Delete(&PsychologicalStatus{}).Error
}
