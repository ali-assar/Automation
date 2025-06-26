package skills

import "backend/internal/core/education"

type Skills struct {
	ID                int64               `gorm:"primaryKey"`
	EducationID       int64               `gorm:"not null"`
	Languages         *string             `gorm:"type:json"` // nullable
	SkillsDescription *string             `gorm:"type:text"` // nullable
	Certificates      *string             `gorm:"type:text"` // nullable
	DeletedAt         int64               `gorm:"not null"`
	Education         education.Education `gorm:"foreignKey:EducationID;references:ID"`
}
