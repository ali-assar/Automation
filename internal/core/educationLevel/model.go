package educationlevel

type EducationLevel struct {
	ID    int64  `gorm:"primaryKey"`
	Level string `gorm:"type:varchar(100);unique;not null"`
}
