package battalion

type Battalion struct {
	ID    int64  `json:"id"`
    Title string `json:"title"`
}

func (Battalion) TableName() string {
	return "battalion"
}
