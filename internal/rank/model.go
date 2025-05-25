package rank

type Rank struct {
    ID        int64  `gorm:"primaryKey"`
    Name      string `gorm:"type:varchar(255);not null"`
    DeletedAt int64  `gorm:"not null"`
}