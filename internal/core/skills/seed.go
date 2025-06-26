package skills

import (
	"backend/internal/core/audit"
	"backend/internal/core/education"

	"gorm.io/gorm"
)

type skillSeed struct {
	University        string
	Languages         *string
	SkillsDescription *string
	Certificates      *string
}

func ptr(s string) *string {
	return &s
}

func SeedSkills(db *gorm.DB, auditService audit.ActionLogger, isTest bool) error {
	if !isTest {
		return nil
	}

	repo := NewRepository(db)
	_ = education.NewRepository(db)

	skills := []skillSeed{
		{
			University:        "State University",
			Languages:         ptr(`["English", "Spanish"]`),
			SkillsDescription: ptr("Programming, Leadership"),
			Certificates:      ptr("AWS Certified Developer"),
		},
		{
			University:        "Unknown Academy",
			Languages:         nil,
			SkillsDescription: nil,
			Certificates:      nil,
		},
	}

	for _, s := range skills {
		// Find EducationID
		var education education.Education
		err := db.Where("university = ? AND deleted_at = 0", s.University).First(&education).Error
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				continue
			}
			return err
		}

		// Check for existing skills
		var existing Skills
		err = db.Where("education_id = ? AND deleted_at = 0", education.ID).First(&existing).Error
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

		skill := &Skills{
			EducationID:       education.ID,
			Languages:         s.Languages,
			SkillsDescription: s.SkillsDescription,
			Certificates:      s.Certificates,
			DeletedAt:         0,
		}
		if err := repo.Create(skill); err != nil {
			return err
		}
		if _, err := auditService.LogAction(1, "Skills", "seeder"); err != nil {
			// Log error
		}
	}

	return nil
}
