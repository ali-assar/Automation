package credentials

import (
	"backend/internal/admin"
	"time"

	"github.com/google/uuid"
)

type Credentials struct {
	ID           uint        `gorm:"primaryKey" json:"id"`
	AdminID      uuid.UUID   `gorm:"type:uuid;uniqueIndex" json:"admin_id"`
	DynamicToken string      `gorm:"type:varchar(255)" json:"dynamic_token,omitempty"`
	StaticToken  string      `gorm:"type:varchar(255)" json:"static_token,omitempty"`
	Admin        admin.Admin `gorm:"foreignKey:AdminID" json:"-"`
	CreatedAt    time.Time   `gorm:"not null" json:"created_at"`
	UpdatedAt    time.Time   `gorm:"not null" json:"updated_at"`
	DeletedAt    *time.Time  `gorm:"index" json:"deleted_at,omitempty"`
}
