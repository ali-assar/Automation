package credentials

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Create(c *Credentials) error {
	return r.db.Create(c).Error
}

func (r *Repository) GetByAdminID(adminID uuid.UUID) (*Credentials, error) {
	var c Credentials
	if err := r.db.Where("admin_id = ?", adminID).First(&c).Error; err != nil {
		return nil, err
	}
	return &c, nil
}

func (r *Repository) GetAll() ([]Credentials, error) {
	var creds []Credentials
	err := r.db.Where("deleted_at IS NULL").Find(&creds).Error // Filter non-deleted
	return creds, err
}

func (r *Repository) GetSoftDeleted() ([]Credentials, error) {
	var creds []Credentials
	err := r.db.Unscoped().Where("deleted_at IS NOT NULL").Find(&creds).Error
	return creds, err
}

func (r *Repository) Update(c *Credentials) error {
	return r.db.Save(c).Error
}

func (r *Repository) GetDynamicTokenByAdminID(adminID uuid.UUID) (string, error) {
	var c Credentials
	if err := r.db.Where("admin_id = ?", adminID).First(&c).Error; err != nil {
		return "", err
	}
	return c.DynamicToken, nil
}

func (r *Repository) GetStaticTokenByAdminID(adminID uuid.UUID) (string, error) {
	var c Credentials
	if err := r.db.Where("admin_id = ?", adminID).First(&c).Error; err != nil {
		return "", err
	}
	return c.StaticToken, nil
}

func (r *Repository) UpdateDynamicTokenByAdminID(adminID uuid.UUID, token string) error {
	return r.db.Model(&Credentials{}).Where("admin_id = ?", adminID).Update("dynamic_token", token).Error
}

func (r *Repository) UpdateStaticTokenByAdminID(adminID uuid.UUID, token string) error {
	return r.db.Model(&Credentials{}).Where("admin_id = ?", adminID).Update("static_token", token).Error
}

func (r *Repository) DeleteSoft(adminID uuid.UUID) error {
	return r.db.Model(&Credentials{}).Where("admin_id = ?", adminID).
		Update("deleted_at", time.Now()).Error
}

func (r *Repository) DeleteHard(adminID uuid.UUID) error {
	return r.db.Where("admin_id = ?", adminID).Delete(&Credentials{}).Error
}
