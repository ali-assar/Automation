package skills

import (
	"backend/internal/core/audit"
	"backend/internal/core/education"

	"gorm.io/gorm"
)

func SeedSkills(db *gorm.DB, auditService audit.ActionLogger, isTest bool) error {
	if !isTest {
		return nil
	}

	repo := NewRepository(db)
	_ = education.NewRepository(db)
	skills := []struct {
		University        string // To find EducationID
		Languages         string
		SkillsDescription string
		Certificates      string
	}{
		{
			University:        "State University",
			Languages:         `["English", "Spanish"]`,
			SkillsDescription: "Programming, Leadership",
			Certificates:      "AWS Certified Developer",
		},
		{
			University:        "National Defense Academy",
			Languages:         `["English", "Arabic"]`,
			SkillsDescription: "Tactical Planning, Communication",
			Certificates:      "Military Leadership Certificate",
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
