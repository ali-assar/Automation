package bloodgroup

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

func (s *Service) CreateBloodGroup(group, actionBy string) (int64, error) {
	bloodGroup := BloodGroup{Name: group}
	if err := s.repo.Create(&bloodGroup); err != nil {
		return 0, err
	}
	if _, err := s.auditService.LogAction(1, "BloodGroup", actionBy); err != nil {
		// Log error but don’t fail
	}
	return bloodGroup.ID, nil
}

func (s *Service) GetBloodGroupByID(id int64) (*BloodGroup, error) {
	return s.repo.GetByID(id)
}

func (s *Service) GetBloodGroupByGroup(group string) (*BloodGroup, error) {
	return s.repo.GetByGroup(group)
}

func (s *Service) GetAllBloodGroups() ([]BloodGroup, error) {
	return s.repo.GetAll()
}

func (s *Service) UpdateBloodGroup(id int64, group, actionBy string) error {
	bloodGroup, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}
	bloodGroup.Name = group
	if err := s.repo.Update(bloodGroup); err != nil {
		return err
	}
	if _, err := s.auditService.LogAction(2, "BloodGroup", actionBy); err != nil {
		// Log error
	}
	return nil
}

func (s *Service) DeleteBloodGroup(id int64, actionBy string) error {
	if err := s.repo.Delete(id); err != nil {
		return err
	}
	if _, err := s.auditService.LogAction(3, "BloodGroup", actionBy); err != nil {
		// Log error
	}
	return nil
}
