package bloodgroup

import (
	"backend/internal/core/audit"

	"gorm.io/gorm"
)

func SeedBloodGroup(db *gorm.DB, auditService audit.ActionLogger) error {
	repo := NewRepository(db)
	bloodGroups := []string{
		"A+", "A-", "B+", "B-", "AB+", "AB-", "O+", "O-",
	}

	for _, group := range bloodGroups {
		_, err := repo.GetByGroup(group)
		if err == nil {
			// if existing.DeletedAt != 0 {
			// 	existing.DeletedAt = 0
			// 	if err := repo.Update(existing); err != nil {
			// 		return err
			// 	}
			// }
			continue
		} else if err != gorm.ErrRecordNotFound {
			return err
		}

		bloodGroup := &BloodGroup{
			Group: group,
		}
		if err := repo.Create(bloodGroup); err != nil {
			return err
		}
		if _, err := auditService.LogAction(1, "BloodGroup", "seeder"); err != nil {
			// Log error
		}
	}

	return nil
}
