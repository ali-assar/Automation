package physicalstatus

import (
	"backend/internal/core/audit"

	"gorm.io/gorm"
)

type Service struct {
	repo         *Repository
	auditService audit.ActionLogger
}

func NewService(db *gorm.DB, auditService audit.ActionLogger) *Service {
	return &Service{
		repo:         NewRepository(db),
		auditService: auditService,
	}
}

func (s *Service) CreatePhysicalStatus(status, description, actionBy string) (int64, error) {
	physicalStatus := PhysicalStatus{
		Status:      status,
		Description: description,
		DeletedAt:   0,
	}
	if err := s.repo.Create(&physicalStatus); err != nil {
		return 0, err
	}
	if _, err := s.auditService.LogAction(1, "PhysicalStatus", actionBy); err != nil {
		// Log error but don’t fail
	}
	return physicalStatus.ID, nil
}

func (s *Service) GetPhysicalStatusByID(id int64) (*PhysicalStatus, error) {
	return s.repo.GetByID(id)
}

func (s *Service) GetPhysicalStatusByName(status string) (*PhysicalStatus, error) {
	return s.repo.GetByStatus(status)
}

func (s *Service) GetAllPhysicalStatuses() ([]PhysicalStatus, error) {
	return s.repo.GetAll()
}

func (s *Service) UpdatePhysicalStatus(id int64, updates map[string]interface{}, actionBy string) error {
	physicalStatus, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}
	delete(updates, "id")
	delete(updates, "deleted_at")
	if err := s.repo.db.Model(physicalStatus).Updates(updates).Error; err != nil {
		return err
	}
	if _, err := s.auditService.LogAction(2, "PhysicalStatus", actionBy); err != nil {
		// Log error
	}
	return nil
}

func (s *Service) DeletePhysicalStatus(id int64, actionBy string) error {
	physicalStatus, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}
	if physicalStatus.DeletedAt != 0 {
		return nil // Already deleted
	}
	if err := s.repo.DeleteSoft(id); err != nil {
		return err
	}
	if _, err := s.auditService.LogAction(3, "PhysicalStatus", actionBy); err != nil {
		// Log error
	}
	return nil
}
