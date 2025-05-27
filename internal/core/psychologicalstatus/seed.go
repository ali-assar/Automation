package psychologicalstatus

import (
	"backend/internal/core/audit"

	"gorm.io/gorm"
)

func SeedPsychologicalStatus(db *gorm.DB, auditService audit.ActionLogger) error {
	repo := NewRepository(db)
	statuses := []PsychologicalStatus{
		{Status: "Stable"},
		{Status: "Anxious"},
		{Status: "Depressed"},
	}

	for _, s := range statuses {
		existing, err := repo.GetByID(s.ID)
		if err == nil && existing != nil {
			continue
		}
		if err != nil && err != gorm.ErrRecordNotFound {
			return err
		}

		if err := repo.Create(&s); err != nil {
			return err
		}
		if _, err := auditService.LogAction(1, "psychological_status", "seeder"); err != nil {
			// Log error
		}
	}
	return nil
}
