package hospitaldispatch

import "backend/internal/core/hospitalvisit"

type HospitalDispatch struct {
	ID            int64               `gorm:"primaryKey;autoIncrement"`
	VisitID       int64               `gorm:"not null"`
	DispatchDate  int64               `gorm:"not null"` // Unix time
	DoctorComment string              `gorm:"type:text"`
	DeletedAt     int64               `gorm:"not null"`
	Visit         hospitalvisit.Visit `gorm:"foreignKey:VisitID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
}

func (HospitalDispatch) TableName() string {
	return "hospital_dispatch"
}
