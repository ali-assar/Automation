package person

import (
	"backend/internal/core/contactinfo"
	"backend/internal/core/familyinfo"
	"backend/internal/core/militarydetails"
	"backend/internal/core/persontype"
	"backend/internal/core/religion"
	"backend/internal/core/skills"
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
	FamilyInfo        familyinfo.FamilyInfo           `gorm:"foreignKey:FamilyInfoID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	ContactInfo       contactinfo.ContactInfo         `gorm:"foreignKey:ContactInfoID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	Skills            skills.Skills                   `gorm:"foreignKey:SkillsID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	Religion          religion.Religion               `gorm:"foreignKey:ReligionID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	PersonType        persontype.PersonType           `gorm:"foreignKey:PersonTypeID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	MilitaryDetails   militarydetails.MilitaryDetails `gorm:"foreignKey:MilitaryDetailsID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
}

func (Person) TableName() string {
	return "person"
}
