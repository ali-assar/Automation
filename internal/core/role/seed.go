package role

import (
	"backend/internal/core/audit"

	"gorm.io/gorm"
)

func SeedRole(db *gorm.DB, auditService audit.ActionLogger) error {
	repo := NewRepository(db)
	roles := []struct {
		ID   int64
		Type string
	}{
		{ID: 1, Type: "SuperAdmin"},
		{ID: 2, Type: "Admin"},
		{ID: 3, Type: "User"},
	}

	for _, r := range roles {
		existing, err := repo.GetByID(r.ID)
		if err == nil {
			if existing.DeletedAt != 0 {
				existing.DeletedAt = 0
				if err := repo.Update(existing); err != nil {
					return err
				}
			}
			if existing.Type != r.Type {
				existing.Type = r.Type
				if err := repo.Update(existing); err != nil {
					return err
				}
			}
			continue
		} else if err != gorm.ErrRecordNotFound {
			return err
		}

		role := &Role{
			ID:        r.ID,
			Type:      r.Type,
			DeletedAt: 0,
		}
		if err := repo.Create(role); err != nil {
			return err
		}
		if _, err := auditService.LogAction(1, "Role", "seeder"); err != nil {
			// Log error
		}
	}

	return nil
}
