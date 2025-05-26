package bloodgroup

import (
	"backend/internal/core/audit"

	"gorm.io/gorm"
)

func SeedBloodGroup(db *gorm.DB, auditService audit.ActionLogger) error {
	repo := NewRepository(db)
	bloodGroups := []struct {
		ID   int64
		Name string
	}{
		{ID: 1, Name: "A+"},
		{ID: 2, Name: "A-"},
		{ID: 3, Name: "B+"},
		{ID: 4, Name: "B-"},
		{ID: 5, Name: "AB+"},
		{ID: 6, Name: "AB-"},
		{ID: 7, Name: "O+"},
		{ID: 8, Name: "O-"},
	}

	for _, bg := range bloodGroups {
		_, err := repo.GetByName(bg.Name)
		if err == nil {
			continue
		} else if err != gorm.ErrRecordNotFound {
			return err
		}

		bloodGroup := &BloodGroup{
			ID:   bg.ID,
			Name: bg.Name,
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
