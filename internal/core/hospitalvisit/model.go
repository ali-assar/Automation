package hospitalvisit

import (
	"backend/internal/core/person"
)

type Visit struct {
	ID        int64         `gorm:"primaryKey;autoIncrement"`
	PersonID  string        `gorm:"type:varchar(255);not null"`
	Date      int64         `gorm:"not null"` // Unix time
	Reason    string        `gorm:"type:text"`
	Diagnosis string        `gorm:"type:text"`
	Treatment string        `gorm:"type:text"`
	DeletedAt int64         `gorm:"not null"`
	Person    person.Person `gorm:"foreignKey:PersonID;references:NationalIDNumber;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
}

func (Visit) TableName() string {
	return "visits" // Already consistent
}
