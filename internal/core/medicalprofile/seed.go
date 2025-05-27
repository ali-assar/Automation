package medicalprofile

import (
	"backend/internal/core/audit"

	"gorm.io/gorm"
)

func SeedMedicalProfile(db *gorm.DB, auditService audit.ActionLogger, isTest bool) error {
	if !isTest {
		return nil
	}

	repo := NewRepository(db)
	profiles := []MedicalProfile{
		{
			PersonID:              "012345678", // Assumes this Person exists
			Allergies:             "Peanuts",
			MedicalHistory:        "Asthma",
			Vaccinations:          `{"vaccines": ["MMR", "Flu"]}`,
			BloodTypeID:           1, // Assumes BloodType with ID 1 exists
			PsychologicalStatusID: 1, // Assumes PsychologicalStatus with ID 1 exists
			DeletedAt:             0,
		},
	}

	for _, p := range profiles {
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
		if _, err := auditService.LogAction(1, "medical_profile", "seeder"); err != nil {
			// Log error
		}
	}
	return nil
}
