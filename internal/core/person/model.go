package person

import (
	"backend/internal/core/contactinfo"
	"backend/internal/core/familyinfo"
	"backend/internal/core/militarydetails"
	"backend/internal/core/persontype"
	"backend/internal/core/physicalinfo"
	"backend/internal/core/religion"
	"backend/internal/core/skills"
	"database/sql"
	"time"
)

type Person struct {
	NationalIDNumber  string                          `gorm:"primaryKey;type:varchar(255)"`
	FirstName         string                          `gorm:"type:varchar(255);not null"`
	LastName          string                          `gorm:"type:varchar(255);not null"`
	FamilyInfoID      sql.NullInt64                   `gorm:"column:family_info_id"`   // Nullable initially
	ContactInfoID     sql.NullInt64                   `gorm:"column:contact_info_id"`  // Nullable initially
	SkillsID          sql.NullInt64                   `gorm:"column:skills_id"`        // Nullable initially
	PhysicalInfoID    sql.NullInt64                   `gorm:"column:physical_info_id"` // Nullable initially
	BirthDate         time.Time                       `gorm:"type:date;not null"`
	ReligionID        sql.NullInt64                   `gorm:"column:religion_id"`         // Nullable initially
	PersonTypeID      sql.NullInt64                   `gorm:"column:person_type_id"`      // Nullable initially
	MilitaryDetailsID sql.NullInt64                   `gorm:"column:military_details_id"` // Nullable initially
	DeletedAt         int64                           `gorm:"not null"`
	FamilyInfo        familyinfo.FamilyInfo           `gorm:"foreignKey:FamilyInfoID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	ContactInfo       contactinfo.ContactInfo         `gorm:"foreignKey:ContactInfoID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	Skills            skills.Skills                   `gorm:"foreignKey:SkillsID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	Religion          religion.Religion               `gorm:"foreignKey:ReligionID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	PersonType        persontype.PersonType           `gorm:"foreignKey:PersonTypeID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	MilitaryDetails   militarydetails.MilitaryDetails `gorm:"foreignKey:MilitaryDetailsID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	PhysicalInfo      physicalinfo.PhysicalInfo       `gorm:"foreignKey:PhysicalInfoID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
}

func (Person) TableName() string {
	return "person"
}

// Helper methods to set and get values from sql.NullInt64
func (p *Person) SetFamilyInfoID(id int64) {
	p.FamilyInfoID.Valid = true
	p.FamilyInfoID.Int64 = id
}

func (p *Person) SetContactInfoID(id int64) {
	p.ContactInfoID.Valid = true
	p.ContactInfoID.Int64 = id
}

func (p *Person) SetSkillsID(id int64) {
	p.SkillsID.Valid = true
	p.SkillsID.Int64 = id
}

func (p *Person) SetPhysicalInfoID(id int64) {
	p.PhysicalInfoID.Valid = true
	p.PhysicalInfoID.Int64 = id
}

func (p *Person) SetReligionID(id int64) {
	p.ReligionID.Valid = true
	p.ReligionID.Int64 = id
}

func (p *Person) SetPersonTypeID(id int64) {
	p.PersonTypeID.Valid = true
	p.PersonTypeID.Int64 = id
}

func (p *Person) SetMilitaryDetailsID(id int64) {
	p.MilitaryDetailsID.Valid = true
	p.MilitaryDetailsID.Int64 = id
}
