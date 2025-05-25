package physicalstatus

type PhysicalStatus struct {
	ID          int64  `gorm:"primaryKey"`
	Status      string `gorm:"type:varchar(255);not null"`
	Description string `gorm:"type:json;not null"`
	DeletedAt   int64  `gorm:"not null"`
}
