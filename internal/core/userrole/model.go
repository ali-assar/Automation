package userrole

type UserRole struct {
	ID     int64
	UserID int64
	RoleID int64
}

func (UserRole) TableName() string {
	return "user_role"
}
