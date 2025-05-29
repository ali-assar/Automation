package physicalinfo

import (
	"backend/internal/core/bloodgroup"
	"backend/internal/core/gender"
	"backend/internal/core/physicalstatus"
)

type PhysicalInfo struct {
	ID               int64  `gorm:"primaryKey;autoIncrement"` // Added autoIncrement
	Height           int    `gorm:"not null"`
	Weight           int    `gorm:"not null"`
	EyeColor         string `gorm:"type:varchar(255);not null"`
	BloodGroupID     int64  `gorm:"not null"`
	GenderID         int64  `gorm:"not null"`
	PhysicalStatusID int64  `gorm:"not null"`
	DeletedAt        int64  `gorm:"not null"`
	DescriptionOfHealth      string `gorm:"type:json;not null"`

	BloodGroup     bloodgroup.BloodGroup         `gorm:"foreignKey:BloodGroupID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	Gender         gender.Gender                 `gorm:"foreignKey:GenderID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	PhysicalStatus physicalstatus.PhysicalStatus `gorm:"foreignKey:PhysicalStatusID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
}
