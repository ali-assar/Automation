package militarydetails

import "backend/internal/core/rank"

type MilitaryDetails struct {
	ID                  int64     `gorm:"primaryKey"`
	RankID              int64     `gorm:"not null"`
	ServiceStartDate    *int64    `gorm:""`
	ServiceDispatchDate *int64    `gorm:""`
	ServiceUnit         *string   `gorm:""`
	BattalionUnit       *string   `gorm:""`
	CompanyUnit         *string   `gorm:""`
	DeletedAt           int64     `gorm:"not null"`
	RankRef             rank.Rank `gorm:"foreignKey:RankID;references:ID"`
}
