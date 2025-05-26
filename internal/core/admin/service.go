package admin

import (
	"backend/internal/core/audit"
	"backend/internal/core/role"
	"backend/pkg/security"
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Service struct {
	repo         *Repository
	auditService audit.ActionLogger
	roleService  *role.Service
}

func NewService(db *gorm.DB, auditService audit.ActionLogger, roleService *role.Service) *Service {
	return &Service{
		repo:         NewRepository(db),
		auditService: auditService,
		roleService:  roleService,
	}
}

func (s *Service) CreateAdmin(nationalIDNumber, userName, hashPassword string, roleID, credentialsID int64, actionBy string) (uuid.UUID, error) {
	// Validate RoleID
	if _, err := s.roleService.GetRoleByID(roleID); err != nil {
		return uuid.Nil, fmt.Errorf("invalid RoleID: %w", err)
	}
	admin := Admin{
		NationalIDNumber: nationalIDNumber,
		UserName:         userName,
		HashPassword:     hashPassword,
		RoleID:           roleID,
		CredentialsID:    credentialsID,
		DeletedAt:        0,
	}
	if err := s.repo.Create(&admin); err != nil {
		return uuid.Nil, err
	}
	if _, err := s.auditService.LogAction(1, "Admin", actionBy); err != nil {
		// Log error
	}
	return admin.ID, nil
}

func (s *Service) GetAdminByUsername(username string) (*Admin, error) {
	return s.repo.GetByUsername(username)
}

func (s *Service) GetAdminByID(id uuid.UUID) (*Admin, error) {
	return s.repo.GetByID(id)
}

func (s *Service) UpdateAdminPassword(adminID uuid.UUID, newPassword string, actionBy string) error {
	passwordHash, err := security.HashPassword(newPassword)
	if err != nil {
		return err
	}

	admin, err := s.repo.GetByID(adminID)
	if err != nil {
		return err
	}

	admin.HashPassword = passwordHash
	if err := s.repo.Update(admin); err != nil {
		return err
	}

	if _, err := s.auditService.LogAction(2, "Admin", actionBy); err != nil {
		// Log error
	}

	return nil
}

func (s *Service) DeleteAdmin(adminID uuid.UUID, actionBy string) error {
	admin, err := s.repo.GetByID(adminID)
	if err != nil {
		return err
	}

	if err := s.repo.DeleteHard(admin); err != nil {
		return err
	}

	if _, err := s.auditService.LogAction(3, "Admin", actionBy); err != nil {
		// Log error
	}

	return nil
}

func (s *Service) AuthenticateAdmin(username, password string) (bool, error) {
	admin, err := s.repo.GetByUsername(username)
	if err != nil {
		return false, err
	}

	return security.ComparePasswords(admin.HashPassword, password)
}

func (s *Service) GetAllAdmins() ([]Admin, error) {
	return s.repo.GetAll()
}

func (s *Service) UpdateAdmin(id uuid.UUID, updates map[string]interface{}, actionBy string) error {
	admin, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}
	// Validate RoleID if provided
	if roleID, ok := updates["role_id"]; ok {
		roleIDInt, ok := roleID.(float64) // JSON numbers are float64
		if !ok {
			return fmt.Errorf("invalid RoleID type")
		}
		if _, err := s.roleService.GetRoleByID(int64(roleIDInt)); err != nil {
			return fmt.Errorf("invalid RoleID: %w", err)
		}
	}
	delete(updates, "id")
	delete(updates, "deleted_at")
	if err := s.repo.db.Model(admin).Updates(updates).Error; err != nil {
		return err
	}
	if _, err := s.auditService.LogAction(2, "Admin", actionBy); err != nil {
		// Log error
	}
	return nil
}

func (s *Service) SoftDeleteAdmin(id uuid.UUID, actionBy string) error {
	admin, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}

	if admin.DeletedAt != 0 {
		return nil // Already deleted
	}

	if err := s.repo.DeleteSoft(id); err != nil {
		return err
	}

	if _, err := s.auditService.LogAction(3, "Admin", actionBy); err != nil {
		// Log error
	}

	return nil
}
