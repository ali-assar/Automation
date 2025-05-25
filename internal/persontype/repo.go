package persontype

import (
    "gorm.io/gorm"
)

type Repository struct {
    db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
    return &Repository{db: db}
}

func (r *Repository) Create(personType *PersonType) error {
    return r.db.Create(personType).Error
}

func (r *Repository) GetByID(id int64) (*PersonType, error) {
    var personType PersonType
    if err := r.db.First(&personType, "id = ?", id).Error; err != nil {
        return nil, err
    }
    return &personType, nil
}

func (r *Repository) GetByType(typeName string) (*PersonType, error) {
    var personType PersonType
    if err := r.db.First(&personType, "type = ?", typeName).Error; err != nil {
        return nil, err
    }
    return &personType, nil
}

func (r *Repository) GetAll() ([]PersonType, error) {
    var personTypes []PersonType
    if err := r.db.Find(&personTypes).Error; err != nil {
        return nil, err
    }
    return personTypes, nil
}

func (r *Repository) Update(personType *PersonType) error {
    return r.db.Save(personType).Error
}

func (r *Repository) Delete(id int64) error {
    return r.db.Delete(&PersonType{}, "id = ?", id).Error
}