package religion

import (
	"backend/internal/core/audit"

	"gorm.io/gorm"
)

func SeedReligion(db *gorm.DB, auditService audit.ActionLogger) error {
	repo := NewRepository(db)
	religions := []struct {
		ID           int64
		ReligionName string
		ReligionType string
	}{
		// اسلام
		{ID: 1, ReligionName: "اسلام", ReligionType: "شیعه دوازده‌امامی"},
		{ID: 2, ReligionName: "اسلام", ReligionType: "شیعه اسماعیلی"},
		{ID: 3, ReligionName: "اسلام", ReligionType: "سنی حنفی"},
		{ID: 4, ReligionName: "اسلام", ReligionType: "سنی شافعی"},
		{ID: 5, ReligionName: "اسلام", ReligionType: "زیدی"},
		// مسیحیت
		{ID: 6, ReligionName: "مسیحیت", ReligionType: "ارتدکس ارمنی"},
		{ID: 7, ReligionName: "مسیحیت", ReligionType: "کلیسای آشوری شرق"},
		{ID: 8, ReligionName: "مسیحیت", ReligionType: "کلیسای کلدانی"},
		// یهودیت
		{ID: 9, ReligionName: "یهودیت", ReligionType: "یهودیت"},
		// زرتشتی
		{ID: 10, ReligionName: "زرتشتی", ReligionType: "زرتشتی"},
		// بهائی
		{ID: 11, ReligionName: "بهائی", ReligionType: "بهائی"},
		// یارسان (اهل حق)
		{ID: 12, ReligionName: "یارسان", ReligionType: "اهل حق"},
		// مندایی
		{ID: 13, ReligionName: "مندایی", ReligionType: "مندایی"},
		// یزیدی
		{ID: 14, ReligionName: "یزیدی", ReligionType: "یزیدی"},
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
