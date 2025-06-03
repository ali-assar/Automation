package physicalinfo

import (
	"backend/internal/core/audit"
	"backend/internal/core/bloodgroup"
	"backend/internal/core/gender"
	"backend/internal/core/physicalstatus"

	"gorm.io/gorm"
)

// SeedPhysicalInfo populates test data for PhysicalInfo when isTest=true
func SeedPhysicalInfo(db *gorm.DB, auditService audit.ActionLogger, isTest bool) error {
	if !isTest {
		return nil
	}

	repo := NewRepository(db)
	bgRepo := bloodgroup.NewRepository(db)
	gRepo := gender.NewRepository(db)
	psRepo := physicalstatus.NewRepository(db)

	// test records
	records := []struct {
		Height              int
		Weight              int
		EyeColor            string
		BloodGroup          string
		Gender              int64
		PhysicalStatusID    int64
		DescriptionOfHealth string
	}{
		{175, 70, "Brown", "A+", 1, 1, `"Healthy, no issues"`},
		{165, 60, "Blue", "O-", 2, 3, `"Joint issue detected"`},
	}

	for _, r := range records {
		// fetch dependencies
		bg, err := bgRepo.GetByGroup(r.BloodGroup)
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				continue
			}
			return err
		}

		// assume gender repo has GetByName or similar
		g, err := gRepo.GetByID(r.Gender)
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				continue
			}
			return err
		}

		ps, err := psRepo.GetByID(r.PhysicalStatusID)
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				continue
			}
			return err
		}

		// check existing
		var existing PhysicalInfo
		err = db.Where("blood_group_id = ? AND gender_id = ? AND deleted_at = 0", bg.ID, g.ID).
			First(&existing).Error
		if err == nil {
			// update description if changed
			if existing.DescriptionOfHealth != &r.DescriptionOfHealth {
				existing.DescriptionOfHealth = &r.DescriptionOfHealth
				if err := repo.Update(&existing); err != nil {
					return err
				}
			}
			continue
		} else if err != gorm.ErrRecordNotFound {
			return err
		}

		// create new record
		newPI := &PhysicalInfo{
			Height:              r.Height,
			Weight:              r.Weight,
			EyeColor:            r.EyeColor,
			BloodGroupID:        bg.ID,
			GenderID:            g.ID,
			PhysicalStatusID:    ps.ID,
			DescriptionOfHealth: &r.DescriptionOfHealth,
			DeletedAt:           0,
		}
		if err := repo.Create(newPI); err != nil {
			return err
		}
		if _, err := auditService.LogAction(1, "PhysicalInfo", "seeder"); err != nil {
			// ignore audit errors
		}
	}

	return nil
}
