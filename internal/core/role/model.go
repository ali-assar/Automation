package role

type Role struct {
    ID        int64  `gorm:"primaryKey"`
    Type      string `gorm:"type:varchar(255);not null"`
    DeletedAt int64  `gorm:"not null"`
}