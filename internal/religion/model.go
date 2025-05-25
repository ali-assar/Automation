package religion

type Religion struct {
    ID           int64  `gorm:"primaryKey"`
    ReligionName string `gorm:"type:varchar(255);not null"`
    ReligionType string `gorm:"type:varchar(255);not null"`
}