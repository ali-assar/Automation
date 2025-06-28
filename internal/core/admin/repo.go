package admin

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

func (r *Repository) Create(admin *Admin) error {
	return r.db.Create(admin).Error
}

func (r *Repository) GetByID(id uuid.UUID) (*Admin, error) {
    var admin Admin
    if err := r.db.Preload("Person").First(&admin, "id = ?", id).Error; err != nil {
        return nil, err
    }
    return &admin, nil
}
func (r *Repository) GetByUsername(username string) (*Admin, error) {
	var admin Admin
	if err := r.db.Preload("Person").First(&admin, "user_name = ?", username).Error; err != nil {
		return nil, err
	}
	return &admin, nil
}

func (r *Repository) GetAll() ([]Admin, error) {
    var admins []Admin
    if err := r.db.Preload("Person").Preload("Role").Where("deleted_at = 0").Find(&admins).Error; err != nil {
        return nil, err
    }
    return admins, nil
}

func (r *Repository) Update(admin *Admin) error {
	return r.db.Save(admin).Error
}

func (r *Repository) DeleteSoft(id uuid.UUID) error {
	return r.db.Model(&Admin{}).Where("id = ?", id).
		Update("deleted_at", time.Now().Unix()).Error
}

func (r *Repository) DeleteHard(admin *Admin) error {
	return r.db.Delete(admin).Error
}
