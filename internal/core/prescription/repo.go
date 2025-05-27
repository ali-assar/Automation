package prescriptions

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

func (r *Repository) Create(prescription *Prescription) error {
	return r.db.Create(prescription).Error
}

func (r *Repository) GetByID(id int64) (*Prescription, error) {
	var prescription Prescription
	if err := r.db.Preload("Visit").Preload("Medicine").First(&prescription, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &prescription, nil
}

func (r *Repository) GetAll() ([]Prescription, error) {
	var prescriptions []Prescription
	if err := r.db.Preload("Visit").Preload("Medicine").Where("deleted_at = 0").Find(&prescriptions).Error; err != nil {
		return nil, err
	}
	return prescriptions, nil
}

func (r *Repository) Update(prescription *Prescription) error {
	return r.db.Save(prescription).Error
}

func (r *Repository) DeleteSoft(id int64) error {
	return r.db.Model(&Prescription{}).Where("id = ?", id).
		Update("deleted_at", time.Now().Unix()).Error
}

func (r *Repository) DeleteHard(id int64) error {
	return r.db.Where("id = ?", id).Delete(&Prescription{}).Error
}