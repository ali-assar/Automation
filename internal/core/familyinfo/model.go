package familyinfo

type FamilyInfo struct {
	ID             int64   `gorm:"primaryKey"`
	FatherDetails  string  `gorm:"type:json;not null"`
	MotherDetails  string  `gorm:"type:json;not null"`
	ChildsDetails  *string `gorm:"type:json"` // nullable
	HusbandDetails *string `gorm:"type:json"` // nullable
	DeletedAt      int64   `gorm:"not null"`
}
