package rank

import (
	"backend/internal/core/audit"

	"gorm.io/gorm"
)

func SeedRank(db *gorm.DB, auditService audit.ActionLogger) error {
	repo := NewRepository(db)
	ranks := []struct {
		ID   int64
		Name string
	}{
		{ID: 1, Name: "Private"},
		{ID: 2, Name: "Sergeant"},
		{ID: 3, Name: "Lieutenant"},
	}

	for _, r := range ranks {
		existing, err := repo.GetByID(r.ID)
		if err == nil {
			if existing.DeletedAt != 0 {
				existing.DeletedAt = 0
				if err := repo.Update(existing); err != nil {
					return err
				}
			}
			if existing.Name != r.Name {
				existing.Name = r.Name
				if err := repo.Update(existing); err != nil {
					return err
				}
			}
			continue
		} else if err != gorm.ErrRecordNotFound {
			return err
		}

		rank := &Rank{
			ID:        r.ID,
			Name:      r.Name,
			DeletedAt: 0,
		}
		if err := repo.Create(rank); err != nil {
			return err
		}
		if _, err := auditService.LogAction(1, "rank", "seeder"); err != nil {
			// Log error
		}
	}

	return nil
}
