package medicalprofile

import (
	"backend/internal/core/person"
	"backend/internal/core/physicalinfo"
	"backend/internal/core/psychologicalstatus"
)

type MedicalProfile struct {
	ID             int64  `gorm:"primaryKey;autoIncrement"`
	PersonID       string `gorm:"type:varchar(255);not null"`
	PhysicalInfoID int64  `gorm:"not null"`

	Allergies             string                                  `gorm:"type:text"`
	MedicalHistory        string                                  `gorm:"type:text"`
	Vaccinations          string                                  `gorm:"type:json"`
	BloodTypeID           int64                                   `gorm:"not null"`
	PsychologicalStatusID int64                                   `gorm:"not null"`
	DeletedAt             int64                                   `gorm:"not null"`
	Person                person.Person                           `gorm:"foreignKey:PersonID;references:NationalIDNumber;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	PhysicalInfo          physicalinfo.PhysicalInfo               `gorm:"foreignKey:PhysicalInfoID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	PsychologicalStatus   psychologicalstatus.PsychologicalStatus `gorm:"foreignKey:PsychologicalStatusID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
}
