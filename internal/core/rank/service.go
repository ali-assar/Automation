package rank

import (
	"backend/internal/core/audit"
	"fmt"

	"gorm.io/gorm"
)

type Service struct {
	repo  *Repository
	audit audit.ActionLogger
}

func NewService(db *gorm.DB, audit audit.ActionLogger) *Service {
	return &Service{
		repo:  NewRepository(db),
		audit: audit,
	}
}

func (s *Service) CreateRank(name string, actionBy string) error {
	rank := &Rank{
		Name:      name,
		DeletedAt: 0,
	}
	if err := s.repo.Create(rank); err != nil {
		return fmt.Errorf("failed to create rank: %v", err)
	}
	if _, err := s.audit.LogAction(1, "Rank", actionBy); err != nil {
		// Log error but don’t fail
	}
	return nil
}

func (s *Service) GetRankByID(id int64) (*Rank, error) {
	rank, err := s.repo.GetByID(id)
	if err != nil {
		return nil, fmt.Errorf("failed to get rank: %v", err)
	}
	return rank, nil
}

func (s *Service) GetRankByName(name string) (string, error) {
	rank, err := s.repo.GetByName(name)
	if err != nil {
		return "", fmt.Errorf("failed to get rank: %v", err)
	}
	return rank, nil
}

func (s *Service) GetAllRanks() ([]Rank, error) {
	ranks, err := s.repo.GetAll()
	if err != nil {
		return nil, fmt.Errorf("failed to get all ranks: %v", err)
	}
	return ranks, nil
}

func (s *Service) UpdateRank(id int64, updates map[string]interface{}, actionBy string) error {
	rank, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}
	delete(updates, "id")
	delete(updates, "deleted_at")
	if err := s.repo.db.Model(rank).Updates(updates).Error; err != nil {
		return err
	}
	if _, err := s.audit.LogAction(2, "Rank", actionBy); err != nil {
		// Log error
	}
	return nil
}

func (s *Service) DeleteRank(id int64, actionBy string) error {
	rank, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}
	if rank.DeletedAt != 0 {
		return nil // Already deleted
	}
	if err := s.repo.DeleteSoft(id); err != nil {
		return err
	}
	if _, err := s.audit.LogAction(3, "Rank", actionBy); err != nil {
		// Log error
	}
	return nil
}
