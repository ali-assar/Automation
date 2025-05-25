package physicalinfo

import (
	"backend/internal/bloodgroup"
	"backend/internal/gender"
	"backend/internal/physicalstatus"
)

type PhysicalInfo struct {
	ID               int64                         `gorm:"primaryKey"`
	Height           int                           `gorm:"not null"` // "hight" corrected to "Height"
	Weight           int                           `gorm:"not null"`
	EyeColor         string                        `gorm:"type:varchar(255);not null"`
	BloodGroupID     int64                         `gorm:"not null"`
	GenderID         int64                         `gorm:"not null"`
	PhysicalStatusID int64                         `gorm:"not null"`
	DeletedAt        int64                         `gorm:"not null"`
	BloodGroup       bloodgroup.BloodGroup         `gorm:"foreignKey:BloodGroupID;references:ID"`
	Gender           gender.Gender                 `gorm:"foreignKey:GenderID;references:ID"`
	PhysicalStatus   physicalstatus.PhysicalStatus `gorm:"foreignKey:PhysicalStatusID;references:ID"`
}
