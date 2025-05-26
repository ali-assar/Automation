package seeder

import (
	"backend/internal/core/actiontype"
	"backend/internal/core/admin"
	"backend/internal/core/audit"
	"backend/internal/core/bloodgroup"
	"backend/internal/core/contactinfo"
	"backend/internal/core/credentials"
	"backend/internal/core/education"
	"backend/internal/core/familyinfo"
	"backend/internal/core/gender"
	"backend/internal/core/militarydetails"
	"backend/internal/core/person"
	"backend/internal/core/persontype"
	"backend/internal/core/physicalinfo"
	"backend/internal/core/physicalstatus"
	"backend/internal/core/rank"
	"backend/internal/core/religion"
	"backend/internal/core/role"
	"backend/internal/core/skills"
	"backend/internal/db"
	"errors"
	"log"

	"gorm.io/gorm"
)

func Seed(isTest bool, auditService audit.ActionLogger) error {
	db := db.GetDB()
	if db == nil {
		return errors.New("database connection is nil")
	}

	// Seed immutable data
	seedFunctions := []struct {
		name string
		fn   func(*gorm.DB, audit.ActionLogger) error
	}{
		{"ActionType", actiontype.SeedActionType},
		{"BloodGroup", bloodgroup.SeedBloodGroup},
		{"Gender", gender.SeedGender},
		{"PersonType", persontype.SeedPersonType},
		{"PhysicalStatus", physicalstatus.SeedPhysicalStatus},
		{"Rank", rank.SeedRank},
		{"Religion", religion.SeedReligion},
		{"Role", role.SeedRole},
	}

	for _, sf := range seedFunctions {
		log.Printf("Seeding %s...", sf.name)
		if err := sf.fn(db, auditService); err != nil {
			log.Printf("Failed to seed %s: %v", sf.name, err)
			return err
		}
		log.Printf("Seeded %s successfully", sf.name)
	}

	// Seed mock data
	if isTest {
		mockSeedFunctions := []struct {
			name string
			fn   func(*gorm.DB, audit.ActionLogger, bool) error
		}{
			{"Education", education.SeedEducation},
			{"Skills", skills.SeedSkills},
			{"PhysicalInfo", physicalinfo.SeedPhysicalInfo},
			{"MilitaryDetails", militarydetails.SeedMilitaryDetails},
			{"FamilyInfo", familyinfo.SeedFamilyInfo},
			{"ContactInfo", contactinfo.SeedContactInfo},
			{"Person", person.SeedPerson},
			{"Admin", admin.SeedAdmin},
			{"Credentials", credentials.SeedCredentials},
		}

		for _, sf := range mockSeedFunctions {
			log.Printf("Seeding %s (test)...", sf.name)
			if err := sf.fn(db, auditService, isTest); err != nil {
				log.Printf("Failed to seed %s: %v", sf.name, err)
				return err
			}
			log.Printf("Seeded %s successfully", sf.name)
		}
	}

	return nil
}
