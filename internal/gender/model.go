package gender

type Gender struct {
    ID     int64  `gorm:"primaryKey"`
    Gender string `gorm:"type:varchar(255);not null"`
}