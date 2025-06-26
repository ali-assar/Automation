package physicalstatus

import (
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Create(physicalStatus *PhysicalStatus) error {
	return r.db.Create(physicalStatus).Error
}

func (r *Repository) GetByID(id int64) (*PhysicalStatus, error) {
	var physicalStatus PhysicalStatus
	if err := r.db.First(&physicalStatus, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &physicalStatus, nil
}

func (r *Repository) GetByStatus(status string) (*PhysicalStatus, error) {
	var physicalStatus PhysicalStatus
	if err := r.db.First(&physicalStatus, "status = ?", status).Error; err != nil {
		return nil, err
	}
	return &physicalStatus, nil
}

func (r *Repository) GetAll() ([]PhysicalStatus, error) {
	var physicalStatuses []PhysicalStatus
	if err := r.db.Find(&physicalStatuses).Error; err != nil {
		return nil, err
	}
	return physicalStatuses, nil
}