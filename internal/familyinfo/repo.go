package familyinfo

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

func (r *Repository) Create(familyInfo *FamilyInfo) error {
	return r.db.Create(familyInfo).Error
}

func (r *Repository) GetByID(id int64) (*FamilyInfo, error) {
	var familyInfo FamilyInfo
	if err := r.db.First(&familyInfo, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &familyInfo, nil
}

func (r *Repository) GetAll() ([]FamilyInfo, error) {
	var familyInfos []FamilyInfo
	if err := r.db.Where("deleted_at = 0").Find(&familyInfos).Error; err != nil {
		return nil, err
	}
	return familyInfos, nil
}

func (r *Repository) Update(familyInfo *FamilyInfo) error {
	return r.db.Save(familyInfo).Error
}

func (r *Repository) DeleteSoft(id int64) error {
	return r.db.Model(&FamilyInfo{}).Where("id = ?", id).
		Update("deleted_at", time.Now().Unix()).Error
}

func (r *Repository) DeleteHard(id int64) error {
	return r.db.Delete(&FamilyInfo{}, "id = ?", id).Error
}
