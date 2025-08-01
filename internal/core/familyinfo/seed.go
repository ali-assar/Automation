package familyinfo

import (
	"backend/internal/core/audit"

	"gorm.io/gorm"
)

func SeedFamilyInfo(db *gorm.DB, auditService audit.ActionLogger, isTest bool) error {
	if !isTest {
		return nil
	}

	repo := NewRepository(db)
	families := []struct {
		FatherDetails  string
		MotherDetails  string
		ChildsDetails  string
		HusbandDetails string
	}{
		{
			FatherDetails:  `{"name": "John Sr."}`,
			MotherDetails:  `{"name": "Mary"}`,
			ChildsDetails:  `[{"name": "Junior"}]`,
			HusbandDetails: `{"name": ""}`,
		},
		{
			FatherDetails:  `{"name": "Robert"}`,
			MotherDetails:  `{"name": "Susan"}`,
			ChildsDetails:  `[{"name": "Emma"}]`,
			HusbandDetails: `{"name": "Robert"}`,
		},
	}

	for _, f := range families {
		var existing FamilyInfo
		// Use ->> to extract the 'name' field from father_details JSON
		err := db.Where("father_details->>'name' = ? AND deleted_at = 0", f.FatherDetails).First(&existing).Error
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

		childs := f.ChildsDetails
		husbend := f.HusbandDetails

		family := &FamilyInfo{
			FatherDetails:  f.FatherDetails,
			MotherDetails:  f.MotherDetails,
			ChildsDetails:  &childs,
			HusbandDetails: &husbend,
			DeletedAt:      0,
		}
		if err := repo.Create(family); err != nil {
			return err
		}

		if _, err := auditService.LogAction(1, "FamilyInfo", "seeder"); err != nil {
			// Log error
		}
	}

	return nil
}
