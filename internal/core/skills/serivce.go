package skills

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

func (s *Service) CreateSkills(educationID int64, languages, skillsDescription, certificates, actionBy string) (int64, error) {
	skills := Skills{
		EducationID:       educationID,
		Languages:         languages,
		SkillsDescription: skillsDescription,
		Certificates:      certificates,
		DeletedAt:         0,
	}
	if err := s.repo.Create(&skills); err != nil {
		return 0, err
	}
	if _, err := s.auditService.LogAction(1, "Skills", actionBy); err != nil {
		// Log error but don’t fail
	}
	return skills.ID, nil
}

func (s *Service) GetSkillsByID(id int64) (*Skills, error) {
	return s.repo.GetByID(id)
}

func (s *Service) GetSkillsByEducationID(educationID int64) ([]Skills, error) {
	return s.repo.GetByEducationID(educationID)
}

func (s *Service) GetAllSkills() ([]Skills, error) {
	return s.repo.GetAll()
}

func (s *Service) UpdateSkills(id int64, updates map[string]interface{}, actionBy string) error {
	skills, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}
	delete(updates, "id")
	delete(updates, "deleted_at")
	if err := s.repo.db.Model(skills).Updates(updates).Error; err != nil {
		return err
	}
	if _, err := s.auditService.LogAction(2, "Skills", actionBy); err != nil {
		// Log error
	}
	return nil
}

func (s *Service) DeleteSkills(id int64, actionBy string) error {
	skills, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}
	if skills.DeletedAt != 0 {
		return nil // Already deleted
	}
	if err := s.repo.DeleteSoft(id); err != nil {
		return err
	}
	if _, err := s.auditService.LogAction(3, "Skills", actionBy); err != nil {
		// Log error
	}
	return nil
}
