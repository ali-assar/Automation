package education

import (
	"backend/internal/core/audit"

	"gorm.io/gorm"
)

func SeedEducation(db *gorm.DB, auditService audit.ActionLogger, isTest bool) error {
	if !isTest {
		return nil
	}

	repo := NewRepository(db)
	educations := []struct {
		EducationLevelID int64
		FieldOfStudy     int64
		Description      string
		University       string
		StartDate        int64
		EndDate          int64
	}{
		{
			EducationLevelID: 1, // Placeholder
			Description:      "BSc in Computer Science",
			University:       "State University",
			StartDate:        1577836800, // 2020-01-01
			EndDate:          1654041600, // 2022-06-01
		},
		{
			EducationLevelID: 2,
			Description:      "MA in Military Strategy",
			University:       "National Defense Academy",
			StartDate:        1609459200, // 2021-01-01
			EndDate:          1685577600, // 2023-06-01
		},
	}

	for _, e := range educations {
		// Use raw query for idempotency, as no specific GetBy method
		var existing Education
		err := db.Where("university = ? AND description = ? AND deleted_at = 0", e.University, e.Description).First(&existing).Error
		if err == nil {
			if existing.DeletedAt != 0 {
				existing.DeletedAt = 0
				if err := repo.Update(&existing); err != nil {
					return err
				}
			}
			continue
		} else if err != gorm.ErrRecordNotFound {
			return err
		}
		desc := e.Description
		uni := e.University
		start := e.StartDate
		end := e.EndDate

		education := &Education{
			EducationLevelID: e.EducationLevelID,
			Description:      &desc,
			University:       &uni,
			StartDate:        &start,
			EndDate:          &end,
			DeletedAt:        0,
		}

		if err := repo.Create(education); err != nil {
			return err
		}

		if _, err := auditService.LogAction(1, "Education", "seeder"); err != nil {
			// Log error
		}
	}

	return nil
}
