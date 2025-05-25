package contactinfo

import (
	"backend/internal/audit"

	"gorm.io/gorm"
)

func SeedContactInfo(db *gorm.DB, auditService audit.ActionLogger, isTest bool) error {
	if !isTest {
		return nil
	}

	repo := NewRepository(db)
	contacts := []struct {
		Address              string
		PhoneNumber          string
		EmergencyPhoneNumber string
		LandlinePhone        string
		EmailAddress         string
		SocialMedia          string
	}{
		{
			Address:              "123 Main St, City",
			PhoneNumber:          "5551234567",
			EmergencyPhoneNumber: "5559876543",
			LandlinePhone:        "5551112222",
			EmailAddress:         "john.doe@example.com",
			SocialMedia:          `{"twitter": "@johndoe"}`,
		},
		{
			Address:              "456 Oak Ave, Town",
			PhoneNumber:          "5552345678",
			EmergencyPhoneNumber: "5558765432",
			LandlinePhone:        "5553334444",
			EmailAddress:         "jane.smith@example.com",
			SocialMedia:          `{"twitter": "@janesmith"}`,
		},
	}

	for _, c := range contacts {
		existing, err := repo.GetByEmail(c.EmailAddress)
		if err == nil {
			if existing.DeletedAt != 0 {
				existing.DeletedAt = 0
				if err := repo.Update(existing); err != nil {
					return err
				}
			}
			continue
		} else if err != gorm.ErrRecordNotFound {
			return err
		}

		contact := &ContactInfo{
			Address:              c.Address,
			PhoneNumber:          c.PhoneNumber,
			EmergencyPhoneNumber: c.EmergencyPhoneNumber,
			LandlinePhone:        c.LandlinePhone,
			EmailAddress:         c.EmailAddress,
			SocialMedia:          c.SocialMedia,
			DeletedAt:            0,
		}
		if err := repo.Create(contact); err != nil {
			return err
		}
		if _, err := auditService.LogAction(1, "ContactInfo", "seeder"); err != nil {
			// Log error
		}
	}

	return nil
}
