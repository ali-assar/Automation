package role

type Role struct {
	ID    int64  `json:"id"`
	Title string `json:"title"`
}

func (Role) TableName() string {
	return "roles"
}
