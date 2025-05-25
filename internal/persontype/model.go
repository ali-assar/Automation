package persontype

type PersonType struct {
	ID   int64  `gorm:"primaryKey"`
	Type string `gorm:"type:varchar(255);not null"`
}
