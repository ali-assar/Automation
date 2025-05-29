package gender

import (
	"backend/internal/core/audit"

	"gorm.io/gorm"
)

func SeedGender(db *gorm.DB, auditService audit.ActionLogger) error {
	repo := NewRepository(db)
	genders := []struct {
		ID     int64
		Gender string
	}{
		{ID: 1, Gender: "مرد"},
		{ID: 2, Gender: "زن"},
	}

	for _, g := range genders {
		existing, err := repo.GetByID(g.ID)
		if err == nil {
			// if existing.DeletedAt != 0 {
			//     // existing.DeletedAt = 0
			//     // if err := repo.Update(existing); err != nil {
			//     //     return err
			//     // }
			// }
			if existing.Gender != g.Gender {
				existing.Gender = g.Gender
				if err := repo.Update(existing); err != nil {
					return err
				}
			}
			continue
		} else if err != gorm.ErrRecordNotFound {
			return err
		}

		gender := &Gender{
			ID:     g.ID,
			Gender: g.Gender,
		}
		if err := repo.Create(gender); err != nil {
			return err
		}
		if _, err := auditService.LogAction(1, "gender", "seeder"); err != nil {
            // Log error
        }
	}

	return nil
}
