package contactinfo

type ContactInfo struct {
	ID                   int64   `gorm:"primaryKey"`
	Address              string  `gorm:"type:varchar(255);not null"`
	PhoneNumber          string  `gorm:"not null"`
	EmergencyPhoneNumber string  `gorm:"not null"`
	LandlinePhone        string  `gorm:"not null"`
	EmailAddress         *string `gorm:"type:varchar(255)"` // Nullable
	SocialMedia          *string `gorm:"type:json"`         // Nullable
	DeletedAt            int64   `gorm:"not null"`
}
