package medicines

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

func (s *Service) CreateMedicine(medicine *Medicine, actionBy string) (int64, error) {
	medicine.DeletedAt = 0
	if err := s.repo.Create(medicine); err != nil {
		return 0, err
	}
	if _, err := s.auditService.LogAction(1, "Medicine", actionBy); err != nil {
		// Log error but don’t fail
	}
	return medicine.ID, nil
}

func (s *Service) GetMedicineByID(id int64) (*Medicine, error) {
	return s.repo.GetByID(id)
}

func (s *Service) GetAllMedicines() ([]Medicine, error) {
	return s.repo.GetAll()
}

func (s *Service) UpdateMedicine(id int64, updates map[string]interface{}, actionBy string) error {
	delete(updates, "id")
	delete(updates, "deleted_at")

	medicine, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}

	if err := s.repo.db.Model(medicine).Updates(updates).Error; err != nil {
		return err
	}

	if _, err := s.auditService.LogAction(2, "Medicine", actionBy); err != nil {
		// Log error
	}
	return nil
}

func (s *Service) DeleteMedicineSoft(id int64, actionBy string) error {
	medicine, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}

	if medicine.DeletedAt != 0 {
		return nil // Already deleted
	}

	if err := s.repo.DeleteSoft(id); err != nil {
		return err
	}

	if _, err := s.auditService.LogAction(3, "Medicine", actionBy); err != nil {
		// Log error
	}
	return nil
}

func (s *Service) DeleteMedicineHard(id int64, actionBy string) error {
	if err := s.repo.DeleteHard(id); err != nil {
		return err
	}

	if _, err := s.auditService.LogAction(3, "Medicine", actionBy); err != nil {
		// Log error
	}
	return nil
}
