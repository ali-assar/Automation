package religion

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

func (s *Service) CreateReligion(religionName, religionType, actionBy string) (int64, error) {
	religion := Religion{
		ReligionName: religionName,
		ReligionType: religionType,
	}
	if err := s.repo.Create(&religion); err != nil {
		return 0, err
	}
	if _, err := s.auditService.LogAction(1, "Religion", actionBy); err != nil {
		// Log error but don’t fail
	}
	return religion.ID, nil
}

func (s *Service) GetReligionByID(id int64) (*Religion, error) {
	return s.repo.GetByID(id)
}

func (s *Service) GetReligionByName(religionName string) (*Religion, error) {
	return s.repo.GetByName(religionName)
}

func (s *Service) GetAllReligions() ([]Religion, error) {
	return s.repo.GetAll()
}

func (s *Service) UpdateReligion(id int64, religionName, religionType, actionBy string) error {
	religion, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}
	religion.ReligionName = religionName
	religion.ReligionType = religionType
	if err := s.repo.Update(religion); err != nil {
		return err
	}
	if _, err := s.auditService.LogAction(2, "Religion", actionBy); err != nil {
		// Log error
	}
	return nil
}

func (s *Service) DeleteReligion(id int64, actionBy string) error {
	if err := s.repo.Delete(id); err != nil {
		return err
	}
	if _, err := s.auditService.LogAction(3, "Religion", actionBy); err != nil {
		// Log error
	}
	return nil
}
