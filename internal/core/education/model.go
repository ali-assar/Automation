package education

type Education struct {
    ID              int64  `gorm:"primaryKey"`
    EducationLevelID int64 `gorm:"not null"`
    FieldOfStudy    int64  `gorm:"not null"`
    Description     string `gorm:"type:varchar(255);not null"`
    University      string `gorm:"type:varchar(255);not null"`
    StartDate       int64  `gorm:"not null"`
    EndDate         int64  `gorm:"not null"`
    DeletedAt       int64  `gorm:"not null"`
}