package education

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

func (s *Service) CreateEducation(educationLevelID int64, description, university string, startDate, endDate int64, actionBy string) (int64, error) {
	education := Education{
		EducationLevelID: educationLevelID,
		Description:      description,
		University:       university,
		StartDate:        startDate,
		EndDate:          endDate,
		DeletedAt:        0,
	}
	if err := s.repo.Create(&education); err != nil {
		return 0, err
	}
	if _, err := s.auditService.LogAction(1, "Education", actionBy); err != nil {
		// Log error but don’t fail
	}
	return education.ID, nil
}

func (s *Service) GetEducationByID(id int64) (*Education, error) {
	return s.repo.GetByID(id)
}

func (s *Service) GetAllEducations() ([]Education, error) {
	return s.repo.GetAll()
}

func (s *Service) UpdateEducation(id int64, updates map[string]interface{}, actionBy string) error {
	education, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}
	delete(updates, "id")
	delete(updates, "deleted_at")
	if err := s.repo.db.Model(education).Updates(updates).Error; err != nil {
		return err
	}
	if _, err := s.auditService.LogAction(2, "Education", actionBy); err != nil {
		// Log error
	}
	return nil
}

func (s *Service) DeleteEducation(id int64, actionBy string) error {
	education, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}
	if education.DeletedAt != 0 {
		return nil // Already deleted
	}
	if err := s.repo.DeleteSoft(id); err != nil {
		return err
	}
	if _, err := s.auditService.LogAction(3, "Education", actionBy); err != nil {
		// Log error
	}
	return nil
}

func (s *Service) SearchEducationsByUniversity(university, actionBy string) ([]Education, error) {
	educations, err := s.repo.SearchByUniversity(university)
	if err != nil {
		return nil, err
	}
	if _, err := s.auditService.LogAction(4, "Education", actionBy); err != nil {
		// Log error
	}
	return educations, nil
}
