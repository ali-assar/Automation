package person

import (
	"backend/internal/contactinfo"
	"backend/internal/familyinfo"
	"backend/internal/militarydetails"
	"backend/internal/persontype"
	"backend/internal/physicalinfo"
	"backend/internal/religion"
	"backend/internal/skills"
	"time"
)

type Person struct {
	NationalIDNumber  string                          `gorm:"primaryKey;type:varchar(255)"`
	FirstName         string                          `gorm:"type:varchar(255);not null"`
	LastName          string                          `gorm:"type:varchar(255);not null"`
	FamilyInfoID      int64                           `gorm:"not null"`
	PhysicalInfoID    int64                           `gorm:"not null"`
	ContactInfoID     int64                           `gorm:"not null"`
	SkillsID          int64                           `gorm:"not null"`
	BirthDate         time.Time                       `gorm:"type:date;not null"`
	ReligionID        int64                           `gorm:"not null"`
	PersonTypeID      int64                           `gorm:"not null"`
	MilitaryDetailsID int64                           `gorm:"not null"`
	DeletedAt         int64                           `gorm:"not null"`
	FamilyInfo        familyinfo.FamilyInfo           `gorm:"foreignKey:FamilyInfoID;references:ID"`
	PhysicalInfo      physicalinfo.PhysicalInfo       `gorm:"foreignKey:PhysicalInfoID;references:ID"`
	ContactInfo       contactinfo.ContactInfo         `gorm:"foreignKey:ContactInfoID;references:ID"`
	Skills            skills.Skills                   `gorm:"foreignKey:SkillsID;references:ID"`
	Religion          religion.Religion               `gorm:"foreignKey:ReligionID;references:ID"`
	PersonType        persontype.PersonType           `gorm:"foreignKey:PersonTypeID;references:ID"`
	MilitaryDetails   militarydetails.MilitaryDetails `gorm:"foreignKey:MilitaryDetailsID;references:ID"`
}
