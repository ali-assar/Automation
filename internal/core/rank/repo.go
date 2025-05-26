package rank

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

func (r *Repository) Create(rank *Rank) error {
	return r.db.Create(rank).Error
}

func (r *Repository) GetByID(id int64) (*Rank, error) {
	var rank Rank
	if err := r.db.First(&rank, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &rank, nil
}

func (r *Repository) GetAll() ([]Rank, error) {
	var ranks []Rank
	if err := r.db.Find(&ranks).Error; err != nil {
		return nil, err
	}
	return ranks, nil
}

func (r *Repository) Update(rank *Rank) error {
	return r.db.Save(rank).Error
}



func (r *Repository) DeleteSoft(id int64) error {
	return r.db.Model(&Rank{}).Where("id = ?", id).Update("deleted_at", time.Now().Unix()).Error
}
