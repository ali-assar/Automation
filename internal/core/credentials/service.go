package credentials

import (
	"backend/internal/core/audit"
	"backend/internal/logger"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Service struct {
	repo         *Repository
	auditService audit.ActionLogger
	db           *gorm.DB
}

func NewService(db *gorm.DB, auditService audit.ActionLogger) *Service {
	return &Service{
		repo:         NewRepository(db),
		auditService: auditService,
		db:           db,
	}
}

func (s *Service) CreateCredentials(adminID uuid.UUID, staticToken, dynamicToken, actionBy string) (*Credentials, error) {
	if s.db == nil {
		return nil, fmt.Errorf("database connection is nil")
	}

	tx := s.db.Begin()
	if tx.Error != nil {
		return nil, fmt.Errorf("failed to begin transaction: %w", tx.Error)
	}
	defer func() {
		if r := recover(); r != nil {
			logger.Error("Panic in CreateCredentials for adminID %s: %v", adminID.String(), r)
			tx.Rollback()
		}
	}()

	now := time.Now()
	cred := Credentials{
		AdminID:      adminID,
		StaticToken:  staticToken,
		DynamicToken: dynamicToken,
		CreatedAt:    now,
		UpdatedAt:    now,
	}

	existing, err := s.repo.GetByAdminID(adminID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			if err := s.repo.Create(&cred); err != nil {
				tx.Rollback()
				logger.Error("Failed to create credentials for adminID %s: %v", adminID.String(), err)
				return nil, fmt.Errorf("failed to create credentials: %w", err)
			}
		} else {
			tx.Rollback()
			logger.Error("Failed to check existing credentials for adminID %s: %v", adminID.String(), err)
			return nil, fmt.Errorf("failed to check existing credentials: %w", err)
		}
	} else {
		updates := map[string]interface{}{
			"static_token":  staticToken,
			"dynamic_token": dynamicToken,
			"updated_at":    now,
		}
		if err := tx.Model(&Credentials{}).Where("admin_id = ?", adminID).Updates(updates).Error; err != nil {
			tx.Rollback()
			logger.Error("Failed to update credentials for adminID %s: %v", adminID.String(), err)
			return nil, fmt.Errorf("failed to update credentials: %w", err)
		}
		cred.ID = existing.ID
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		logger.Error("Failed to commit transaction for adminID %s: %v", adminID.String(), err)
		return nil, fmt.Errorf("failed to commit transaction: %w", err)
	}

	if _, err := s.auditService.LogAction(1, "Credentials", actionBy); err != nil {
		logger.Error("Failed to log audit action for %s: %v", actionBy, err)
	}
	return &cred, nil
}
func (s *Service) GetCredentialsByAdminID(adminID uuid.UUID) (*Credentials, error) {
	return s.repo.GetByAdminID(adminID)
}

func (s *Service) GetAllCredentials() ([]Credentials, error) {
	return s.repo.GetAll()
}

func (s *Service) GetSoftDeletedCredentials() ([]Credentials, error) {
	return s.repo.GetSoftDeleted()
}

func (s *Service) UpdateCredentials(adminID uuid.UUID, updates map[string]interface{}, actionBy string) error {
	cred, err := s.repo.GetByAdminID(adminID)
	if err != nil {
		return err
	}

	// Prevent updating critical fields
	delete(updates, "id")
	delete(updates, "admin_id")
	delete(updates, "deleted_at")

	if err := s.db.Model(cred).Updates(updates).Error; err != nil {
		return err
	}

	if _, err := s.auditService.LogAction(2, "Credentials", actionBy); err != nil {
		// Log error
	}
	return nil
}

func (s *Service) DeleteCredentials(adminID uuid.UUID, actionBy string) error {
	cred, err := s.repo.GetByAdminID(adminID)
	if err != nil {
		return err
	}

	if cred.DeletedAt != nil {
		return nil // Already deleted
	}

	if err := s.repo.DeleteSoft(adminID); err != nil {
		return err
	}

	if _, err := s.auditService.LogAction(3, "Credentials", actionBy); err != nil {
		// Log error
	}
	return nil
}

func (s *Service) DeleteCredentialsHard(adminID uuid.UUID, actionBy string) error {
	if err := s.repo.DeleteHard(adminID); err != nil {
		return err
	}

	if _, err := s.auditService.LogAction(3, "Credentials", actionBy); err != nil {
		// Log error
	}
	return nil
}

func (s *Service) GetStaticTokenByAdminID(adminID uuid.UUID) (string, error) {
	return s.repo.GetStaticTokenByAdminID(adminID)
}

func (s *Service) GetDynamicTokenByAdminID(adminID uuid.UUID) (string, error) {
	return s.repo.GetDynamicTokenByAdminID(adminID)
}

func (s *Service) UpdateDynamicTokenByAdminID(adminID uuid.UUID, token, actionBy string) error {
	if err := s.repo.UpdateDynamicTokenByAdminID(adminID, token); err != nil {
		return err
	}

	if _, err := s.auditService.LogAction(2, "Credentials", actionBy); err != nil {
		// Log error
	}
	return nil
}
