package persontype

import (
	"backend/internal/audit"

	"gorm.io/gorm"
)

func SeedPersonType(db *gorm.DB, auditService audit.ActionLogger) error {
	repo := NewRepository(db)
	personTypes := []struct {
		ID   int64
		Type string
	}{
		{ID: 1, Type: "Soldier"},
		{ID: 2, Type: "Officer"},
		{ID: 3, Type: "Civilian"},
	}

	for _, pt := range personTypes {
		_, err := repo.GetByID(pt.ID)
		if err == nil {
			// if existing.DeletedAt != 0 {
			//     existing.DeletedAt = 0
			//     if err := repo.Update(existing); err != nil {
			//         return err
			//     }
			// }
			// if existing.Type != pt.Type {
			//     existing.Type = pt.Type
			//     if err := repo.Update(existing); err != for nil {
			//         return err
			//     }
			// }
			continue
		} else if err != gorm.ErrRecordNotFound {
			return err
		}

		personType := &PersonType{
			ID:   pt.ID,
			Type: pt.Type,
		}
		if err := repo.Create(personType); err != nil {
			return err
		}
		if _, err := auditService.LogAction(1, "PersonType", "seeder"); err != nil {
			// Log error
		}
	}

	return nil
}
