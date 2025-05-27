package hospitalvisit

import (
	"backend/internal/core/audit"
	"time"

	"gorm.io/gorm"
)

func SeedVisits(db *gorm.DB, auditService audit.ActionLogger, isTest bool) error {
	if !isTest {
		return nil
	}

	repo := NewRepository(db)
	visits := []Visit{
		{
			PersonID:  "012345678", // Assumes this person exists
			Date:      time.Now().Unix(),
			Reason:    "Routine checkup",
			Diagnosis: "Healthy",
			Treatment: "None",
			DeletedAt: 0,
		},
	}

	for _, v := range visits {
		if err := repo.Create(&v); err != nil {
			return err
		}
		if _, err := auditService.LogAction(1, "visits", "seeder"); err != nil {
			// Log error
		}
	}
	return nil
}
