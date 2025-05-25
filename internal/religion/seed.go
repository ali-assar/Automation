package religion

import (
	"backend/internal/audit"

	"gorm.io/gorm"
)

func SeedReligion(db *gorm.DB, auditService audit.ActionLogger) error {
	repo := NewRepository(db)
	religions := []struct {
		ID           int64
		ReligionName string
		ReligionType string
	}{
		{ID: 1, ReligionName: "Islam", ReligionType: "Monotheistic"},
		{ID: 2, ReligionName: "Christianity", ReligionType: "Monotheistic"},
		{ID: 3, ReligionName: "Judaism", ReligionType: "Monotheistic"},
	}

	for _, r := range religions {
		existing, err := repo.GetByID(r.ID)
		if err == nil {
			// if existing.DeletedAt != 0 {
			// 	existing.DeletedAt = 0
			// 	if err := repo.Update(existing); err != nil {
			// 		return err
			// 	}
			// }
			if existing.ReligionName != r.ReligionName || existing.ReligionType != r.ReligionType {
				existing.ReligionName = r.ReligionName
				existing.ReligionType = r.ReligionType
				if err := repo.Update(existing); err != nil {
					return err
				}
			}
			continue
		} else if err != gorm.ErrRecordNotFound {
			return err
		}

		religion := &Religion{
			ID:           r.ID,
			ReligionName: r.ReligionName,
			ReligionType: r.ReligionType,
		}
		if err := repo.Create(religion); err != nil {
			return err
		}
		if _, err := auditService.LogAction(1, "Religion", "seeder"); err != nil {
			// Log error
		}
	}

	return nil
}
