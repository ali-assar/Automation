package bloodgroup

type BloodGroup struct {
	ID    int64  `gorm:"primaryKey"`
	Group string `gorm:"type:varchar(255);not null"`
}
