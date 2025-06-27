package battalion

type Battalion struct {
	ID    int64
	Title string
}

func (Battalion) TableName() string {
	return "battalion"
}
