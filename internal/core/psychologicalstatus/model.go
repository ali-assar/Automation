package psychologicalstatus

type PsychologicalStatus struct {
	ID     int64  `gorm:"primaryKey;autoIncrement"`
	Status string `gorm:"type:text;not null"`
}

