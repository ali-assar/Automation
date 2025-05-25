package role

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

func (r *Repository) Create(role *Role) error {
	return r.db.Create(role).Error
}

func (r *Repository) GetByID(id int64) (*Role, error) {
	var role Role
	if err := r.db.First(&role, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &role, nil
}

func (r *Repository) GetByType(typeName string) (*Role, error) {
	var role Role
	if err := r.db.First(&role, "type = ?", typeName).Error; err != nil {
		return nil, err
	}
	return &role, nil
}

func (r *Repository) GetAll() ([]Role, error) {
	var roles []Role
	if err := r.db.Where("deleted_at = 0").Find(&roles).Error; err != nil {
		return nil, err
	}
	return roles, nil
}

func (r *Repository) Update(role *Role) error {
	return r.db.Save(role).Error
}

func (r *Repository) DeleteSoft(id int64) error {
	return r.db.Model(&Role{}).Where("id = ?", id).
		Update("deleted_at", time.Now().Unix()).Error
}

func (r *Repository) DeleteHard(id int64) error {
	return r.db.Delete(&Role{}, "id = ?", id).Error
}
