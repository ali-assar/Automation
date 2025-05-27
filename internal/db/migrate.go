package db

import (
	"backend/internal/core/action"
	"backend/internal/core/actiontype"
	"backend/internal/core/admin"
	"backend/internal/core/bloodgroup"
	"backend/internal/core/contactinfo"
	"backend/internal/core/credentials"
	"backend/internal/core/education"
	"backend/internal/core/familyinfo"
	"backend/internal/core/gender"
	"backend/internal/core/hospitaldispatch"
	"backend/internal/core/hospitalvisit"
	"backend/internal/core/medicalprofile"
	"backend/internal/core/medicines"
	"backend/internal/core/militarydetails"
	"backend/internal/core/person"
	"backend/internal/core/persontype"
	"backend/internal/core/physicalinfo"
	"backend/internal/core/physicalstatus"
	prescriptions "backend/internal/core/prescription"
	"backend/internal/core/psychologicalstatus"
	"backend/internal/core/rank"
	"backend/internal/core/religion"
	"backend/internal/core/role"
	"backend/internal/core/skills"
)

func Migrate() error {
	// Migrate tables in dependency order
	return db.AutoMigrate(
		&action.Action{},
		&actiontype.ActionType{},
		&bloodgroup.BloodGroup{},
		&gender.Gender{},
		&persontype.PersonType{},
		&physicalstatus.PhysicalStatus{},
		&rank.Rank{},
		&religion.Religion{},
		&role.Role{},
		&education.Education{},
		&skills.Skills{},
		&physicalinfo.PhysicalInfo{},
		&militarydetails.MilitaryDetails{},
		&familyinfo.FamilyInfo{},
		&contactinfo.ContactInfo{},
		&person.Person{}, // Before admin
		&admin.Admin{},
		&credentials.Credentials{},
		&medicalprofile.MedicalProfile{},
		&hospitalvisit.Visit{},
		&hospitaldispatch.HospitalDispatch{},
		&prescriptions.Prescription{},
		&medicines.Medicine{},
		&psychologicalstatus.PsychologicalStatus{},
	)
}
