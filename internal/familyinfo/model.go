package familyinfo

type FamilyInfo struct {
	ID             int64  `gorm:"primaryKey"`
	FatherDetails  string `gorm:"type:json;not null"`
	MotherDetails  string `gorm:"type:json;not null"`
	ChildsDetails  string `gorm:"type:json;not null"`
	HusbandDetails string `gorm:"type:json;not null"`
	DeletedAt      int64  `gorm:"not null"`
}
