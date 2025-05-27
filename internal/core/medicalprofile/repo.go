package medicalprofile

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

func (r *Repository) Create(profile *MedicalProfile) error {
	return r.db.Create(profile).Error
}

func (r *Repository) GetByID(id int64) (*MedicalProfile, error) {
	var profile MedicalProfile
	if err := r.db.Preload("Person").Preload("PhysicalInfo").Preload("PsychologicalStatus").
		First(&profile, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &profile, nil
}

func (r *Repository) GetByPersonID(personID string) (*MedicalProfile, error) {
	var profile MedicalProfile
	if err := r.db.Preload("Person").Preload("PhysicalInfo").Preload("PsychologicalStatus").
		First(&profile, "person_id = ?", personID).Error; err != nil {
		return nil, err
	}
	return &profile, nil
}

func (r *Repository) GetAll() ([]MedicalProfile, error) {
	var profiles []MedicalProfile
	if err := r.db.Preload("Person").Preload("PhysicalInfo").Preload("PsychologicalStatus").
		Where("deleted_at = 0").Find(&profiles).Error; err != nil {
		return nil, err
	}
	return profiles, nil
}

func (r *Repository) Update(profile *MedicalProfile) error {
	return r.db.Save(profile).Error
}

func (r *Repository) DeleteSoft(id int64) error {
	return r.db.Model(&MedicalProfile{}).Where("id = ?", id).
		Update("deleted_at", time.Now().Unix()).Error
}

func (r *Repository) DeleteHard(id int64) error {
	return r.db.Where("id = ?", id).Delete(&MedicalProfile{}).Error
}
