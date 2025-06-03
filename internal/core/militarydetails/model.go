package militarydetails

import "backend/internal/core/rank"

type MilitaryDetails struct {
	ID                  int64     `gorm:"primaryKey"`
	RankID              int64     `gorm:"not null"`
	ServiceStartDate    *int64    `gorm:""`
	ServiceDispatchDate *int64    `gorm:""`
	ServiceUnit         *int64    `gorm:""`
	BattalionUnit       *int64    `gorm:""`
	CompanyUnit         *int64    `gorm:""`
	DeletedAt           int64     `gorm:"not null"`
	RankRef             rank.Rank `gorm:"foreignKey:RankID;references:ID"`
}
