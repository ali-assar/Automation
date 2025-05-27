package medicines

import (
	"backend/internal/core/audit"
	"gorm.io/gorm"
)

func SeedMedicines(db *gorm.DB, auditService audit.ActionLogger, isTest bool) error {
	if !isTest {
		return nil
	}

	repo := NewRepository(db)
	medicines := []Medicine{
		{
			Name:        "Paracetamol",
			Quantity:    100,
			Description: "Pain reliever",
			DeletedAt:   0,
		},
		{
			Name:        "Ibuprofen",
			Quantity:    50,
			Description: "Anti-inflammatory",
			DeletedAt:   0,
		},
	}

	for _, m := range medicines {
		existing, err := repo.GetByID(m.ID)
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

		if err := repo.Create(&m); err != nil {
			return err
		}
		if _, err := auditService.LogAction(1, "medicines", "seeder"); err != nil {
			// Log error
		}
	}
	return nil
}