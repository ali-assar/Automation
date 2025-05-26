package role

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

func (s *Service) CreateRole(typeName, actionBy string) (int64, error) {
	role := Role{
		Type:      typeName,
		DeletedAt: 0,
	}
	if err := s.repo.Create(&role); err != nil {
		return 0, err
	}
	if _, err := s.auditService.LogAction(1, "Role", actionBy); err != nil {
		// Log error but don’t fail
	}
	return role.ID, nil
}

func (s *Service) GetRoleByID(id int64) (*Role, error) {
	return s.repo.GetByID(id)
}

func (s *Service) GetRoleByType(typeName string) (*Role, error) {
	return s.repo.GetByType(typeName)
}

func (s *Service) GetAllRoles() ([]Role, error) {
	return s.repo.GetAll()
}

func (s *Service) UpdateRole(id int64, typeName, actionBy string) error {
	role, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}
	role.Type = typeName
	if err := s.repo.Update(role); err != nil {
		return err
	}
	if _, err := s.auditService.LogAction(2, "Role", actionBy); err != nil {
		// Log error
	}
	return nil
}

func (s *Service) DeleteRole(id int64, actionBy string) error {
	role, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}
	if role.DeletedAt != 0 {
		return nil // Already deleted
	}
	if err := s.repo.DeleteSoft(id); err != nil {
		return err
	}
	if _, err := s.auditService.LogAction(3, "Role", actionBy); err != nil {
		// Log error
	}
	return nil
}
