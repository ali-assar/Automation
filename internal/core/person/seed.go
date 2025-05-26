package person

import (
	"backend/internal/core/audit"
	"backend/internal/core/contactinfo"
	"backend/internal/core/familyinfo"
	"backend/internal/core/militarydetails"
	"backend/internal/core/persontype"
	"backend/internal/core/physicalinfo"
	"backend/internal/core/religion"
	"backend/internal/core/skills"
	"time"

	"gorm.io/gorm"
)

func SeedPerson(db *gorm.DB, auditService audit.ActionLogger, isTest bool) error {
	if !isTest {
		return nil
	}

	repo := NewRepository(db)
	_ = familyinfo.NewRepository(db)
	_ = physicalinfo.NewRepository(db)
	contactinfoRepo := contactinfo.NewRepository(db)
	_ = skills.NewRepository(db)
	_ = religion.NewRepository(db)
	_ = persontype.NewRepository(db)
	_ = militarydetails.NewRepository(db)

	persons := []struct {
		NationalIDNumber string
		FirstName        string
		LastName         string
		FamilyFather     string // To find FamilyInfoID
		EmailAddress     string // To find ContactInfoID
		BirthDate        time.Time
		ReligionName     string
		PersonType       string
	}{
		{
			NationalIDNumber: "012345678",
			FirstName:        "John",
			LastName:         "Doe",
			FamilyFather:     "John Sr.", // Just the name field
			EmailAddress:     "john.doe@example.com",
			BirthDate:        time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC),
			ReligionName:     "Islam",
			PersonType:       "Soldier",
		},
		{
			NationalIDNumber: "987654321",
			FirstName:        "Jane",
			LastName:         "Smith",
			FamilyFather:     "Robert", // Just the name field
			EmailAddress:     "jane.smith@example.com",
			BirthDate:        time.Date(1992, 6, 15, 0, 0, 0, 0, time.UTC),
			ReligionName:     "Christianity",
			PersonType:       "Officer",
		},
	}

	for _, p := range persons {
		// Find dependencies
		var familyinfo familyinfo.FamilyInfo
		err := db.Where("father_details->>'name' = ? AND deleted_at = 0", p.FamilyFather).First(&familyinfo).Error
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				continue
			}
			return err
		}

		var physicalinfo physicalinfo.PhysicalInfo
		err = db.Where("blood_group_id = ? AND deleted_at = 0", 1).First(&physicalinfo).Error // Assuming A+=1
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				continue
			}
			return err
		}

		contactinfo, err := contactinfoRepo.GetByEmail(p.EmailAddress)
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				continue
			}
			return err
		}

		var skills skills.Skills
		err = db.Where("education_id = ? AND deleted_at = 0", 1).First(&skills).Error // Assuming first education
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				continue
			}
			return err
		}

		var religion religion.Religion
		err = db.Where("religion_name = ?", p.ReligionName).First(&religion).Error
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				continue
			}
			return err
		}

		var persontype persontype.PersonType
		err = db.Where("type = ? ", p.PersonType).First(&persontype).Error
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				continue
			}
			return err
		}

		var militarydetails militarydetails.MilitaryDetails
		err = db.Where("rank_id = ? AND deleted_at = 0", 1).First(&militarydetails).Error // Assuming Private=1
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				continue
			}
			return err
		}

		existing, err := repo.GetByID(p.NationalIDNumber)
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

		person := &Person{
			NationalIDNumber:  p.NationalIDNumber,
			FirstName:         p.FirstName,
			LastName:          p.LastName,
			FamilyInfoID:      familyinfo.ID,
			PhysicalInfoID:    physicalinfo.ID,
			ContactInfoID:     contactinfo.ID,
			SkillsID:          skills.ID,
			BirthDate:         p.BirthDate,
			ReligionID:        religion.ID,
			PersonTypeID:      persontype.ID,
			MilitaryDetailsID: militarydetails.ID,
			DeletedAt:         0,
		}
		if err := repo.Create(person); err != nil {
			return err
		}
		if _, err := auditService.LogAction(1, "person", "seeder"); err != nil {
			// Log error
		}
	}

	return nil
}
