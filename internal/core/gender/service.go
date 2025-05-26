package gender

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

func (s *Service) CreateGender(gender, actionBy string) (int64, error) {
	g := Gender{Gender: gender}
	if err := s.repo.Create(&g); err != nil {
		return 0, err
	}
	if _, err := s.auditService.LogAction(1, "Gender", actionBy); err != nil {
		// Log error but don’t fail
	}
	return g.ID, nil
}

func (s *Service) GetGenderByID(id int64) (*Gender, error) {
	return s.repo.GetByID(id)
}

func (s *Service) GetGenderByName(gender string) (*Gender, error) {
	return s.repo.GetByGender(gender)
}

func (s *Service) GetAllGenders() ([]Gender, error) {
	return s.repo.GetAll()
}

func (s *Service) UpdateGender(id int64, gender, actionBy string) error {
	g, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}
	g.Gender = gender
	if err := s.repo.Update(g); err != nil {
		return err
	}
	if _, err := s.auditService.LogAction(2, "Gender", actionBy); err != nil {
		// Log error
	}
	return nil
}

func (s *Service) DeleteGender(id int64, actionBy string) error {
	if err := s.repo.Delete(id); err != nil {
		return err
	}
	if _, err := s.auditService.LogAction(3, "Gender", actionBy); err != nil {
		// Log error
	}
	return nil
}
