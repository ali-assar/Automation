package physicalinfo

import (
	"backend/internal/core/audit"
	"backend/internal/core/bloodgroup"
	"backend/internal/core/gender"
	"backend/internal/core/physicalstatus"

	"gorm.io/gorm"
)

func SeedPhysicalInfo(db *gorm.DB, auditService audit.ActionLogger, isTest bool) error {
	if !isTest {
		return nil
	}

	repo := NewRepository(db)
	bloodgroupRepo := bloodgroup.NewRepository(db)
	genderRepo := gender.NewRepository(db)
	physicalstatusRepo := physicalstatus.NewRepository(db)
	physicalInfos := []struct {
		Height         int
		Weight         int
		EyeColor       string
		BloodGroup     string
		Gender         string
		PhysicalStatus string
	}{
		{
			Height:         175,
			Weight:         70,
			EyeColor:       "Brown",
			BloodGroup:     "A+",
			Gender:         "Male",
			PhysicalStatus: "Fit",
		},
		{
			Height:         165,
			Weight:         60,
			EyeColor:       "Blue",
			BloodGroup:     "O-",
			Gender:         "Female",
			PhysicalStatus: "Restricted",
		},
	}

	for _, pi := range physicalInfos {
		// Find dependencies
		bloodgroup, err := bloodgroupRepo.GetByGroup(pi.BloodGroup)
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				continue
			}
			return err
		}
		gender, err := genderRepo.GetByID(1) // Assuming Male=1, Female=2
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				continue
			}
			return err
		}
		physicalstatus, err := physicalstatusRepo.GetByID(1) // Assuming Fit=1
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				continue
			}
			return err
		}

		// Check for existing record
		var existing PhysicalInfo
		err = db.Where("blood_group_id = ? AND gender_id = ? AND deleted_at = 0", bloodgroup.ID, gender.ID).First(&existing).Error
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

		physicalInfo := &PhysicalInfo{
			Height:           pi.Height,
			Weight:           pi.Weight,
			EyeColor:         pi.EyeColor,
			BloodGroupID:     bloodgroup.ID,
			GenderID:         gender.ID,
			PhysicalStatusID: physicalstatus.ID,
			DeletedAt:        0,
		}
		if err := repo.Create(physicalInfo); err != nil {
			return err
		}
		if _, err := auditService.LogAction(1, "physicalInfo", "seeder"); err != nil {
            // Log error
        }
	}

	return nil
}
