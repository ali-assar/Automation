package role

type Role struct {
	ID    int64
	Title string
}

func (Role) TableName() string {
	return "roles"
}
