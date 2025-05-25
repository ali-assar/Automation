package skills

import "backend/internal/education"

type Skills struct {
	ID                int64               `gorm:"primaryKey"`
	EducationID       int64               `gorm:"not null"`
	Languages         string              `gorm:"type:json;not null"`
	SkillsDescription string              `gorm:"type:text;not null"`
	Certificates      string              `gorm:"type:text;not null"`
	DeletedAt         int64               `gorm:"not null"`
	Education         education.Education `gorm:"foreignKey:EducationID;references:ID"`
}
