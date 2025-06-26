package hospitaldispatch

import (
	"backend/internal/core/audit"
	"time"

	"gorm.io/gorm"
)

func SeedHospitalDispatch(db *gorm.DB, auditService audit.ActionLogger, isTest bool) error {
	if !isTest {
		return nil
	}

	repo := NewRepository(db)
	dispatches := []HospitalDispatch{
		{
			VisitID:       1, // Assumes a Visit with ID 1 exists
			DispatchDate:  time.Now().Unix(),
			DoctorComment: "Patient transferred to ICU",
			DeletedAt:     0,
		},
	}

	for _, d := range dispatches {
		existing, err := repo.GetByID(d.ID)
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

		if err := repo.Create(&d); err != nil {
			return err
		}
		if _, err := auditService.LogAction(1, "hospital_dispatch", "seeder"); err != nil {
			// Log error
		}
	}
	return nil
}
