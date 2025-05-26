package action

import "backend/internal/core/actiontype"

type Action struct {
	ID            int64                 `gorm:"primaryKey"`
	ActionType    int64                 `gorm:"index;not null"`
	Time          int64                 `gorm:"not null"` // Unix timestamp
	TableName     string                `gorm:"type:varchar(255);not null"`
	ActionBy      string                `gorm:"type:varchar(255);not null"`
	ActionTypeRef actiontype.ActionType `gorm:"foreignKey:ActionType;references:ID"`
}
