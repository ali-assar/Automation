package hospitalvisit

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

func (s *Service) CreateVisit(visit *Visit, actionBy string) (int64, error) {
	if err := s.repo.Create(visit); err != nil {
		return 0, err
	}
	if _, err := s.auditService.LogAction(1, "Visit", actionBy); err != nil {
		// Log error
	}
	return visit.ID, nil
}

func (s *Service) GetVisitByID(id int64) (*Visit, error) {
	return s.repo.GetByID(id)
}

func (s *Service) GetAllVisits() ([]Visit, error) {
	return s.repo.GetAll()
}
func (s *Service) GetVisitsByPersonID(personID string) ([]Visit, error) {
	var visits []Visit
	if err := s.repo.db.Preload("Person").Where("person_id = ? AND deleted_at = 0", personID).Find(&visits).Error; err != nil {
		return nil, err
	}
	return visits, nil
}

func (s *Service) UpdateVisit(id int64, updates map[string]interface{}, actionBy string) error {
	visit, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}
	if err := s.repo.db.Model(visit).Updates(updates).Error; err != nil {
		return err
	}
	if _, err := s.auditService.LogAction(2, "Visit", actionBy); err != nil {
		// Log error
	}
	return nil
}

func (s *Service) DeleteVisitSoft(id int64, actionBy string) error {
	if err := s.repo.DeleteSoft(id); err != nil {
		return err
	}
	if _, err := s.auditService.LogAction(3, "Visit", actionBy); err != nil {
		// Log error
	}
	return nil
}

func (s *Service) DeleteVisitHard(id int64, actionBy string) error {
	if err := s.repo.DeleteHard(id); err != nil {
		return err
	}
	if _, err := s.auditService.LogAction(3, "Visit", actionBy); err != nil {
		// Log error
	}
	return nil
}
