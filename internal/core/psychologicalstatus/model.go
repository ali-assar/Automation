package psychologicalstatus

type PsychologicalStatus struct {
	ID     int64  `gorm:"primaryKey;autoIncrement"`
	Status string `gorm:"type:text;not null"`
}

func (PsychologicalStatus) TableName() string {
	return "psychological_status"
}
