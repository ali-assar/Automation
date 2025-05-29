package rank

import (
	"backend/internal/core/audit"

	"gorm.io/gorm"
)

func SeedRank(db *gorm.DB, auditService audit.ActionLogger) error {
	repo := NewRepository(db)
	ranks := []struct {
		ID   int64
		Name string
	}{
		{ID: 1, Name: "سرباز آموزشی"},
		{ID: 2, Name: "سرباز"},
		{ID: 3, Name: "سرباز دوم"},
		{ID: 4, Name: "سرباز یکم"},
		{ID: 5, Name: "سرجوخه"},
		{ID: 6, Name: "گروهبان سوم"},
		{ID: 7, Name: "گروهبان دوم"},
		{ID: 8, Name: "گروهبان یکم"},
		{ID: 9, Name: "استوار دوم"},
		{ID: 10, Name: "استوار یکم"},
		{ID: 11, Name: "ناویگ دومی"},
		{ID: 12, Name: "ناویگ اول"},
		{ID: 13, Name: "ستوان سوم"},
		{ID: 14, Name: "ستوان دوم"},
		{ID: 15, Name: "ستوان یکم"},
		{ID: 16, Name: "سرگرد"},
		{ID: 17, Name: "سروان"},
		{ID: 18, Name: "سرهنگ دوم"},
		{ID: 19, Name: "سرهنگ یکم"},
		{ID: 20, Name: "سرتیپ دوم"},
		{ID: 21, Name: "سرتیپ یکم"},
		{ID: 22, Name: "سپهبد"},
		{ID: 23, Name: "سرلشکر"},
	}

	for _, r := range ranks {
		existing, err := repo.GetByID(r.ID)
		if err == nil {
			if existing.DeletedAt != 0 {
				existing.DeletedAt = 0
				if err := repo.Update(existing); err != nil {
					return err
				}
			}
			if existing.Name != r.Name {
				existing.Name = r.Name
				if err := repo.Update(existing); err != nil {
					return err
				}
			}
			continue
		} else if err != gorm.ErrRecordNotFound {
			return err
		}

		rank := &Rank{
			ID:        r.ID,
			Name:      r.Name,
			DeletedAt: 0,
		}
		if err := repo.Create(rank); err != nil {
			return err
		}
		if _, err := auditService.LogAction(1, "rank", "seeder"); err != nil {
			// Log error
		}
	}

	return nil
}
