package educationlevel

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

func (s *Service) GetEducationLevelByID(id int64) (*EducationLevel, error) {
	return s.repo.GetByID(id)
}

func (s *Service) GetAllEducations() ([]EducationLevel, error) {
	return s.repo.GetAll()
}
