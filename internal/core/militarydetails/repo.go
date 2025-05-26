package militarydetails

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

func (r *Repository) Create(militaryDetails *MilitaryDetails) error {
	return r.db.Create(militaryDetails).Error
}

func (r *Repository) GetByID(id int64) (*MilitaryDetails, error) {
	var militaryDetails MilitaryDetails
	if err := r.db.First(&militaryDetails, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &militaryDetails, nil
}

func (r *Repository) GetAll() ([]MilitaryDetails, error) {
	var militaryDetails []MilitaryDetails
	if err := r.db.Where("deleted_at = 0").Find(&militaryDetails).Error; err != nil {
		return nil, err
	}
	return militaryDetails, nil
}

func (r *Repository) Update(militaryDetails *MilitaryDetails) error {
	return r.db.Save(militaryDetails).Error
}

func (r *Repository) DeleteSoft(id int64) error {
	return r.db.Model(&MilitaryDetails{}).Where("id = ?", id).
		Update("deleted_at", time.Now().Unix()).Error
}

func (r *Repository) DeleteHard(id int64) error {
	return r.db.Delete(&MilitaryDetails{}, "id = ?", id).Error
}
