package seeder

import (
	"backend/internal/actiontype"
	"backend/internal/admin"
	"backend/internal/audit"
	"backend/internal/bloodgroup"
	"backend/internal/contactinfo"
	"backend/internal/credentials"
	"backend/internal/education"
	"backend/internal/familyinfo"
	"backend/internal/gender"
	"backend/internal/militarydetails"
	"backend/internal/person"
	"backend/internal/persontype"
	"backend/internal/physicalinfo"
	"backend/internal/physicalstatus"
	"backend/internal/rank"
	"backend/internal/religion"
	"backend/internal/role"
	"backend/internal/skills"
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
