package prescriptions

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

func (s *Service) CreatePrescription(prescription *Prescription, actionBy string) (int64, error) {
	prescription.DeletedAt = 0
	if err := s.repo.Create(prescription); err != nil {
		return 0, err
	}
	if _, err := s.auditService.LogAction(1, "Prescription", actionBy); err != nil {
		// Log error but don’t fail
	}
	return prescription.ID, nil
}

func (s *Service) GetPrescriptionByID(id int64) (*Prescription, error) {
	return s.repo.GetByID(id)
}

func (s *Service) GetAllPrescriptions() ([]Prescription, error) {
	return s.repo.GetAll()
}

func (s *Service) UpdatePrescription(id int64, updates map[string]interface{}, actionBy string) error {
	delete(updates, "id")
	delete(updates, "deleted_at")

	prescription, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}

	if err := s.repo.db.Model(prescription).Updates(updates).Error; err != nil {
		return err
	}

	if _, err := s.auditService.LogAction(2, "Prescription", actionBy); err != nil {
		// Log error
	}
	return nil
}

func (s *Service) DeletePrescriptionSoft(id int64, actionBy string) error {
	prescription, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}

	if prescription.DeletedAt != 0 {
		return nil // Already deleted
	}

	if err := s.repo.DeleteSoft(id); err != nil {
		return err
	}

	if _, err := s.auditService.LogAction(3, "Prescription", actionBy); err != nil {
		// Log error
	}
	return nil
}

func (s *Service) DeletePrescriptionHard(id int64, actionBy string) error {
	if err := s.repo.DeleteHard(id); err != nil {
		return err
	}

	if _, err := s.auditService.LogAction(3, "Prescription", actionBy); err != nil {
		// Log error
	}
	return nil
}
