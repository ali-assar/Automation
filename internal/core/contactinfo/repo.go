package contactinfo

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

func (r *Repository) Create(contactInfo *ContactInfo) error {
	return r.db.Create(contactInfo).Error
}

func (r *Repository) GetByID(id int64) (*ContactInfo, error) {
	var contactInfo ContactInfo
	if err := r.db.First(&contactInfo, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &contactInfo, nil
}

func (r *Repository) GetByEmail(email string) (*ContactInfo, error) {
	var contactInfo ContactInfo
	if err := r.db.First(&contactInfo, "email_address = ? AND deleted_at = 0", email).Error; err != nil {
		return nil, err
	}
	return &contactInfo, nil
}

func (r *Repository) GetAll() ([]ContactInfo, error) {
	var contactInfos []ContactInfo
	if err := r.db.Where("deleted_at = 0").Find(&contactInfos).Error; err != nil {
		return nil, err
	}
	return contactInfos, nil
}

func (r *Repository) Update(contactInfo *ContactInfo) error {
	return r.db.Save(contactInfo).Error
}

func (r *Repository) DeleteSoft(id int64) error {
	return r.db.Model(&ContactInfo{}).Where("id = ?", id).
		Update("deleted_at", time.Now().Unix()).Error
}

func (r *Repository) DeleteHard(id int64) error {
	return r.db.Delete(&ContactInfo{}, "id = ?", id).Error
}
