package psychologicalstatus

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

func (s *Service) CreatePsychologicalStatus(status *PsychologicalStatus, actionBy string) (int64, error) {
	if err := s.repo.Create(status); err != nil {
		return 0, err
	}
	if _, err := s.auditService.LogAction(1, "PsychologicalStatus", actionBy); err != nil {
		// Log error but don’t fail
	}
	return status.ID, nil
}

func (s *Service) GetPsychologicalStatusByID(id int64) (*PsychologicalStatus, error) {
	return s.repo.GetByID(id)
}

func (s *Service) GetAllPsychologicalStatuses() ([]PsychologicalStatus, error) {
	return s.repo.GetAll()
}

func (s *Service) UpdatePsychologicalStatus(id int64, updates map[string]interface{}, actionBy string) error {
	delete(updates, "id")

	status, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}

	if err := s.repo.db.Model(status).Updates(updates).Error; err != nil {
		return err
	}

	if _, err := s.auditService.LogAction(2, "PsychologicalStatus", actionBy); err != nil {
		// Log error
	}
	return nil
}

func (s *Service) DeletePsychologicalStatus(id int64, actionBy string) error {
	if err := s.repo.Delete(id); err != nil {
		return err
	}

	if _, err := s.auditService.LogAction(3, "PsychologicalStatus", actionBy); err != nil {
		// Log error
	}
	return nil
}
