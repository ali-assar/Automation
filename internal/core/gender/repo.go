package gender

import (
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Create(gender *Gender) error {
	return r.db.Create(gender).Error
}

func (r *Repository) GetByID(id int64) (*Gender, error) {
	var gender Gender
	if err := r.db.First(&gender, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &gender, nil
}

func (r *Repository) GetByGender(gender string) (*Gender, error) {
	var g Gender
	if err := r.db.First(&g, "gender = ?", gender).Error; err != nil {
		return nil, err
	}
	return &g, nil
}

func (r *Repository) GetAll() ([]Gender, error) {
	var genders []Gender
	if err := r.db.Find(&genders).Error; err != nil {
		return nil, err
	}
	return genders, nil
}

func (r *Repository) Update(gender *Gender) error {
	return r.db.Save(gender).Error
}

func (r *Repository) Delete(id int64) error {
	return r.db.Delete(&Gender{}, "id = ?", id).Error
}
