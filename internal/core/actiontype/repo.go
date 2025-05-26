package actiontype

import (
    "gorm.io/gorm"
)

type Repository struct {
    db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
    return &Repository{db: db}
}

func (r *Repository) Create(actionType *ActionType) error {
    return r.db.Create(actionType).Error
}

func (r *Repository) GetByID(id int64) (*ActionType, error) {
    var actionType ActionType
    if err := r.db.First(&actionType, "id = ?", id).Error; err != nil {
        return nil, err
    }
    return &actionType, nil
}

func (r *Repository) GetAll() ([]ActionType, error) {
    var actionTypes []ActionType
    if err := r.db.Find(&actionTypes).Error; err != nil {
        return nil, err
    }
    return actionTypes, nil
}

func (r *Repository) Update(id int64, updates map[string]interface{}) error {
    return r.db.Model(&ActionType{}).Where("id = ?", id).Updates(updates).Error
}

func (r *Repository) Delete(id int64) error {
    return r.db.Delete(&ActionType{}, "id = ?", id).Error
}