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
	}{
		{ID: 1, Status: "سالم"},
		{ID: 2, Status: "ناسالم"},
	}

	for _, s := range statuses {
		_, err := repo.GetByID(s.ID)
		if err != gorm.ErrRecordNotFound {
			return err
		}

		status := &PhysicalStatus{
			ID:          s.ID,
			Status:      s.Status,
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
