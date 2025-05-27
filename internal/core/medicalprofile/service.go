package medicalprofile

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

func (s *Service) CreateMedicalProfile(profile *MedicalProfile, actionBy string) (int64, error) {
	profile.DeletedAt = 0
	if err := s.repo.Create(profile); err != nil {
		return 0, err
	}
	if _, err := s.auditService.LogAction(1, "MedicalProfile", actionBy); err != nil {
		// Log error but don’t fail
	}
	return profile.ID, nil
}

func (s *Service) GetMedicalProfileByID(id int64) (*MedicalProfile, error) {
	return s.repo.GetByID(id)
}

func (s *Service) GetMedicalProfileByPersonID(personID string) (*MedicalProfile, error) {
	return s.repo.GetByPersonID(personID)
}

func (s *Service) GetAllMedicalProfiles() ([]MedicalProfile, error) {
	return s.repo.GetAll()
}

func (s *Service) UpdateMedicalProfile(id int64, updates map[string]interface{}, actionBy string) error {
	delete(updates, "id")
	delete(updates, "deleted_at")

	profile, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}

	if err := s.repo.db.Model(profile).Updates(updates).Error; err != nil {
		return err
	}

	if _, err := s.auditService.LogAction(2, "MedicalProfile", actionBy); err != nil {
		// Log error
	}
	return nil
}

func (s *Service) DeleteMedicalProfileSoft(id int64, actionBy string) error {
	profile, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}

	if profile.DeletedAt != 0 {
		return nil // Already deleted
	}

	if err := s.repo.DeleteSoft(id); err != nil {
		return err
	}

	if _, err := s.auditService.LogAction(3, "MedicalProfile", actionBy); err != nil {
		// Log error
	}
	return nil
}

func (s *Service) DeleteMedicalProfileHard(id int64, actionBy string) error {
	if err := s.repo.DeleteHard(id); err != nil {
		return err
	}

	if _, err := s.auditService.LogAction(3, "MedicalProfile", actionBy); err != nil {
		// Log error
	}
	return nil
}
