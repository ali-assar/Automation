package actiontype

type ActionType struct {
	ID         int64  `gorm:"primaryKey"`
	ActionName string `gorm:"type:varchar(255);not null"`
}


const (
    ActionCreate int64 = 1
    ActionUpdate int64 = 2
    ActionDelete int64 = 3
    ActionSearch int64 = 4
)