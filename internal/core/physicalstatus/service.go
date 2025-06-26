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
