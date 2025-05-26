package militarydetails

import "backend/internal/core/rank"

type MilitaryDetails struct {
	ID                  int64     `gorm:"primaryKey"`
	RankID              int64     `gorm:"not null"`
	ServiceStartDate    int64     `gorm:"not null"`
	ServiceDispatchDate int64     `gorm:"not null"`
	ServiceUnit         int64     `gorm:"not null"`
	BattalionUnit       int64     `gorm:"not null"`
	CompanyUnit         int64     `gorm:"not null"`
	DeletedAt           int64     `gorm:"not null"`
	RankRef             rank.Rank `gorm:"foreignKey:RankID;references:ID"`
}
