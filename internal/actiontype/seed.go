package actiontype

import (
	"backend/internal/audit"

	"gorm.io/gorm"
)

func SeedActionType(db *gorm.DB, auditService audit.ActionLogger) error {
	repo := NewRepository(db)
	actionTypes := []struct {
		ID         int64
		ActionName string
	}{
		{ID: 1, ActionName: "Create"},
		{ID: 2, ActionName: "Update"},
		{ID: 3, ActionName: "Delete"},
		{ID: 4, ActionName: "Search"},
	}

	for _, at := range actionTypes {
		existing, err := repo.GetByID(at.ID)
		if err == nil {
			// if existing.DeletedAt != 0 {
			// 	existing.DeletedAt = 0
			// 	if err := repo.Update(existing); err != nil {
			// 		return err
			// 	}
			// }
			if existing.ActionName != at.ActionName {
				existing.ActionName = at.ActionName
				// if err := repo.Update(existing); err != nil {
				// 	return err
				// }
			}
			continue
		} else if err != gorm.ErrRecordNotFound {
			return err
		}

		actionType := &ActionType{
			ID:         at.ID,
			ActionName: at.ActionName,
		}
		if err := repo.Create(actionType); err != nil {
			return err
		}
		if _, err := auditService.LogAction(1, "ActionType", "seeder"); err != nil {
			// Log error
		}
	}

	return nil
}
