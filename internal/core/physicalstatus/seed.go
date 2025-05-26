package physicalstatus

import (
	"backend/internal/core/audit"

	"gorm.io/gorm"
)

func SeedPhysicalStatus(db *gorm.DB, auditService audit.ActionLogger) error {
	repo := NewRepository(db)
	statuses := []struct {
		ID          int64
		Status      string
		Description string
	}{
		{ID: 1, Status: "Fit", Description: `{"details": "Fully fit for duty"}`},
		{ID: 2, Status: "Restricted", Description: `{"details": "Limited duty due to injury"}`},
		{ID: 3, Status: "Unfit", Description: `{"details": "Not fit for active duty"}`},
	}

	for _, s := range statuses {
		existing, err := repo.GetByID(s.ID)
		if err == nil {
			if existing.DeletedAt != 0 {
				existing.DeletedAt = 0
				if err := repo.Update(existing); err != nil {
					return err
				}
			}
			if existing.Status != s.Status || existing.Description != s.Description {
				existing.Status = s.Status
				existing.Description = s.Description
				if err := repo.Update(existing); err != nil {
					return err
				}
			}
			continue
		} else if err != gorm.ErrRecordNotFound {
			return err
		}

		status := &PhysicalStatus{
			ID:          s.ID,
			Status:      s.Status,
			Description: s.Description,
			DeletedAt:   0,
		}
		if err := repo.Create(status); err != nil {
			return err
		}
		if _, err := auditService.LogAction(1, "PhysicalStatus", "seeder"); err != nil {
			// Log error
		}
	}

	return nil
}
