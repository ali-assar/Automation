package credentials

import (
	"backend/internal/admin"
	"backend/internal/audit"
	"time"

	"gorm.io/gorm"
)

func SeedCredentials(db *gorm.DB, auditService audit.ActionLogger, isTest bool) error {
	if !isTest {
		return nil
	}

	repo := NewRepository(db)
	adminRepo := admin.NewRepository(db)
	credentials := []struct {
		AdminUsername string
		StaticToken   string
		DynamicToken  string
	}{
		{
			AdminUsername: "admin1",
			StaticToken:   "static_token_1",
			DynamicToken:  "dynamic_token_1",
		},
		{
			AdminUsername: "admin2",
			StaticToken:   "static_token_2",
			DynamicToken:  "dynamic_token_2",
		},
	}

	for _, c := range credentials {
		admin, err := adminRepo.GetByUsername(c.AdminUsername)
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				continue
			}
			return err
		}

		existing, err := repo.GetByAdminID(admin.ID)
		if err == nil {
			if existing.DeletedAt != nil {
				if err := repo.db.Model(&Credentials{}).Where("admin_id = ?", admin.ID).
					Update("deleted_at", nil).Error; err != nil {
					return err
				}
			}
			continue
		} else if err != gorm.ErrRecordNotFound {
			return err
		}

		now := time.Now()
		cred := &Credentials{
			AdminID:      admin.ID,
			StaticToken:  c.StaticToken,
			DynamicToken: c.DynamicToken,
			CreatedAt:    now,
			UpdatedAt:    now,
			DeletedAt:    nil,
		}
		if err := repo.Create(cred); err != nil {
			return err
		}
		if _, err := auditService.LogAction(1, "Credentials", "seeder"); err != nil {
			// Log error
		}
	}

	return nil
}
