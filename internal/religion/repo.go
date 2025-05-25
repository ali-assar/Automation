package religion

import (
    "gorm.io/gorm"
)

type Repository struct {
    db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
    return &Repository{db: db}
}

func (r *Repository) Create(religion *Religion) error {
    return r.db.Create(religion).Error
}

func (r *Repository) GetByID(id int64) (*Religion, error) {
    var religion Religion
    if err := r.db.First(&religion, "id = ?", id).Error; err != nil {
        return nil, err
    }
    return &religion, nil
}

func (r *Repository) GetByName(religionName string) (*Religion, error) {
    var religion Religion
    if err := r.db.First(&religion, "religion_name = ?", religionName).Error; err != nil {
        return nil, err
    }
    return &religion, nil
}

func (r *Repository) GetAll() ([]Religion, error) {
    var religions []Religion
    if err := r.db.Find(&religions).Error; err != nil {
        return nil, err
    }
    return religions, nil
}

func (r *Repository) Update(religion *Religion) error {
    return r.db.Save(religion).Error
}

func (r *Repository) Delete(id int64) error {
    return r.db.Delete(&Religion{}, "id = ?", id).Error
}