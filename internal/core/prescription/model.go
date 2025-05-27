package prescriptions

import (
	"backend/internal/core/hospitalvisit"
	"backend/internal/core/medicines"
)

type Prescription struct {
	ID         int64               `gorm:"primaryKey;autoIncrement"`
	VisitID    int64               `gorm:"not null"`
	MedicineID int64               `gorm:"not null"`
	Dose       string              `gorm:"type:text"`
	Duration   string              `gorm:"type:text"`
	DeletedAt  int64               `gorm:"not null"`
	Visit      hospitalvisit.Visit `gorm:"foreignKey:VisitID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	Medicine   medicines.Medicine  `gorm:"foreignKey:MedicineID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
}

