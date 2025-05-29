package education

import educationlevel "backend/internal/core/educationLevel"

type Education struct {
	ID               int64                         `gorm:"primaryKey"`
	EducationLevelID int64                         `gorm:"not null"`
	EducationLevel   educationlevel.EducationLevel `gorm:"foreignKey:EducationLevelID;references:ID"`
	Description      string                        `gorm:"type:varchar(255);not null"`
	University       string                        `gorm:"type:varchar(255);not null"`
	StartDate        int64                         `gorm:"not null"`
	EndDate          int64                         `gorm:"not null"`
	DeletedAt        int64                         `gorm:"not null"`
}
