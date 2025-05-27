package hospitaldispatch

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

func (s *Service) CreateHospitalDispatch(dispatch *HospitalDispatch, actionBy string) (int64, error) {
	dispatch.DeletedAt = 0
	if err := s.repo.Create(dispatch); err != nil {
		return 0, err
	}
	if _, err := s.auditService.LogAction(1, "HospitalDispatch", actionBy); err != nil {
		// Log error but don’t fail
	}
	return dispatch.ID, nil
}

func (s *Service) GetHospitalDispatchByID(id int64) (*HospitalDispatch, error) {
	return s.repo.GetByID(id)
}

func (s *Service) GetAllHospitalDispatches() ([]HospitalDispatch, error) {
	return s.repo.GetAll()
}

func (s *Service) UpdateHospitalDispatch(id int64, updates map[string]interface{}, actionBy string) error {
	delete(updates, "id")
	delete(updates, "deleted_at")

	dispatch, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}

	if err := s.repo.db.Model(dispatch).Updates(updates).Error; err != nil {
		return err
	}

	if _, err := s.auditService.LogAction(2, "HospitalDispatch", actionBy); err != nil {
		// Log error
	}
	return nil
}

func (s *Service) DeleteHospitalDispatchSoft(id int64, actionBy string) error {
	dispatch, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}

	if dispatch.DeletedAt != 0 {
		return nil // Already deleted
	}

	if err := s.repo.DeleteSoft(id); err != nil {
		return err
	}

	if _, err := s.auditService.LogAction(3, "HospitalDispatch", actionBy); err != nil {
		// Log error
	}
	return nil
}

func (s *Service) DeleteHospitalDispatchHard(id int64, actionBy string) error {
	if err := s.repo.DeleteHard(id); err != nil {
		return err
	}

	if _, err := s.auditService.LogAction(3, "HospitalDispatch", actionBy); err != nil {
		// Log error
	}
	return nil
}
