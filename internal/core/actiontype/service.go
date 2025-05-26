package actiontype

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

func (s *Service) CreateActionType(actionName, actionBy string) (int64, error) {
	actionType := ActionType{ActionName: actionName}
	if err := s.repo.Create(&actionType); err != nil {
		return 0, err
	}
	if _, err := s.auditService.LogAction(1, "ActionType", actionBy); err != nil {
		// Log error but don’t fail
	}
	return actionType.ID, nil
}

func (s *Service) GetActionTypeByID(id int64) (*ActionType, error) {
	return s.repo.GetByID(id)
}

func (s *Service) GetAllActionTypes() ([]ActionType, error) {
	return s.repo.GetAll()
}

func (s *Service) UpdateActionType(id int64, updates map[string]interface{}, actionBy string) error {
	if err := s.repo.Update(id, updates); err != nil {
		return err
	}
	if _, err := s.auditService.LogAction(2, "ActionType", actionBy); err != nil {
		// Log error
	}
	return nil
}

func (s *Service) DeleteActionType(id int64, actionBy string) error {
	if err := s.repo.Delete(id); err != nil {
		return err
	}
	if _, err := s.auditService.LogAction(3, "ActionType", actionBy); err != nil {
		// Log error
	}
	return nil
}
