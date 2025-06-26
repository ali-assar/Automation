package educationlevel

import (
	"backend/internal/core/audit"

	"gorm.io/gorm"
)

func SeedEducationLevels(db *gorm.DB, auditService audit.ActionLogger) error {
	repo := NewRepository(db)
	levels := []string{"بی سواد", "سیکل", "دیپلم", "کارشناسی", "کارشناسی ارشد","دکترا"}

	for _, level := range levels {
		var existing EducationLevel
		if err := db.Where("level = ?", level).First(&existing).Error; err == nil {
			continue
		}

		newLevel := &EducationLevel{Level: level}
		if err := repo.Create(newLevel); err != nil {
			return err
		}

		if _, err := auditService.LogAction(1, "Education", "seeder"); err != nil {
			// Log error
		}
	}

	return nil
}
