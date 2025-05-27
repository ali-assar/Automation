package prescriptions

import (
	"backend/internal/core/audit"

	"gorm.io/gorm"
)

func SeedPrescriptions(db *gorm.DB, auditService audit.ActionLogger, isTest bool) error {
	if !isTest {
		return nil
	}

	repo := NewRepository(db)
	prescriptions := []Prescription{
		{
			VisitID:    1, // Assumes Visit with ID 1 exists
			MedicineID: 1, // Assumes Medicine with ID 1 exists
			Dose:       "500mg",
			Duration:   "7 days",
			DeletedAt:  0,
		},
	}

	for _, p := range prescriptions {
		existing, err := repo.GetByID(p.ID)
		if err == nil && existing.DeletedAt != 0 {
			existing.DeletedAt = 0
			if err := repo.Update(existing); err != nil {
				return err
			}
			continue
		}
		if err != nil && err != gorm.ErrRecordNotFound {
			return err
		}

		if err := repo.Create(&p); err != nil {
			return err
		}
		if _, err := auditService.LogAction(1, "prescriptions", "seeder"); err != nil {
			// Log error
		}
	}
	return nil
}
