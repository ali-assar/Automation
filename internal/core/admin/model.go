package admin

import (
	"backend/internal/core/person"
	"backend/internal/core/role"

	"github.com/google/uuid"
)

type Admin struct {
	ID               uuid.UUID     `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	NationalIDNumber string        `gorm:"type:varchar(255);not null;uniqueIndex"`
	UserName         string        `gorm:"type:varchar(255);not null"`
	HashPassword     string        `gorm:"type:varchar(255);not null"`
	RoleID           int64         `gorm:"not null"`
	DeletedAt        int64         `gorm:"not null"`
	CredentialsID    int64         `gorm:"not null"`
	Person           person.Person `gorm:"foreignKey:NationalIDNumber;references:NationalIDNumber;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;association_autocreate:false;association_autoupdate:false"`
	Role             role.Role     `gorm:"foreignKey:RoleID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
}

func (Admin) TableName() string {
	return "admin"
}
