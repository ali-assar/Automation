package physicalinfo

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

func (r *Repository) Create(physicalInfo *PhysicalInfo) error {
	return r.db.Create(physicalInfo).Error
}

func (r *Repository) GetByID(id int64) (*PhysicalInfo, error) {
	var physicalInfo PhysicalInfo
	if err := r.db.Preload("BloodGroup").Preload("Gender").Preload("PhysicalStatus").
		First(&physicalInfo, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &physicalInfo, nil
}

func (r *Repository) GetAll() ([]PhysicalInfo, error) {
	var physicalInfos []PhysicalInfo
	if err := r.db.Preload("BloodGroup").Preload("Gender").Preload("PhysicalStatus").
		Where("deleted_at = 0").Find(&physicalInfos).Error; err != nil {
		return nil, err
	}
	return physicalInfos, nil
}

func (r *Repository) Update(physicalInfo *PhysicalInfo) error {
	return r.db.Save(physicalInfo).Error
}

func (r *Repository) DeleteSoft(id int64) error {
	return r.db.Model(&PhysicalInfo{}).Where("id = ?", id).
		Update("deleted_at", time.Now().Unix()).Error
}

func (r *Repository) DeleteHard(id int64) error {
	return r.db.Delete(&PhysicalInfo{}, "id = ?", id).Error
}
