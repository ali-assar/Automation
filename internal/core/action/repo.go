package action

import (
    "gorm.io/gorm"
)

type Repository struct {
    db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
    return &Repository{db: db}
}

func (r *Repository) Create(action *Action) error {
    return r.db.Create(action).Error
}

func (r *Repository) GetByID(id int64) (*Action, error) {
    var action Action
    if err := r.db.Preload("ActionTypeRef").First(&action, "id = ?", id).Error; err != nil {
        return nil, err
    }
    return &action, nil
}

func (r *Repository) GetAll() ([]Action, error) {
    var actions []Action
    if err := r.db.Preload("ActionTypeRef").Find(&actions).Error; err != nil {
        return nil, err
    }
    return actions, nil
}

func (r *Repository) Update(id int64, updates map[string]interface{}) error {
    return r.db.Model(&Action{}).Where("id = ?", id).Updates(updates).Error
}

func (r *Repository) Delete(id int64) error {
    return r.db.Delete(&Action{}, "id = ?", id).Error
}