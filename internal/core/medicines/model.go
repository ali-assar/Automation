package medicines

type Medicine struct {
	ID          int64  `gorm:"primaryKey;autoIncrement"`
	Name        string `gorm:"type:text;not null"`
	Quantity    int64  `gorm:"not null"`
	Description string `gorm:"type:text"`
	DeletedAt   int64  `gorm:"not null"`
}

func (Medicine) TableName() string {
	return "medicines"
}
