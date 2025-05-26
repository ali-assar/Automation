package action

import (
	"time"

	"gorm.io/gorm"
)

type Service struct {
	repo *Repository
}

func NewService(db *gorm.DB) *Service {
	return &Service{repo: NewRepository(db)}
}

func (s *Service) LogAction(actionType int64, tableName, actionBy string) (int64, error) {
	action := Action{
		ActionType: actionType,
		Time:       time.Now().Unix(),
		TableName:  tableName,
		ActionBy:   actionBy,
	}
	if err := s.repo.Create(&action); err != nil {
		return 0, err
	}
	return action.ID, nil
}

func (s *Service) GetActionByID(id int64) (*Action, error) {
	return s.repo.GetByID(id)
}

func (s *Service) GetAllActions() ([]Action, error) {
	return s.repo.GetAll()
}

func (s *Service) UpdateAction(id int64, updates map[string]interface{}, actionBy string) error {
	if err := s.repo.Update(id, updates); err != nil {
		return err
	}
	if _, err := s.LogAction(2, "Action", actionBy); err != nil {
		// Log error
	}
	return nil
}

func (s *Service) DeleteAction(id int64, actionBy string) error {
	if err := s.repo.Delete(id); err != nil {
		return err
	}
	if _, err := s.LogAction(3, "Action", actionBy); err != nil {
		// Log error
	}
	return nil
}
