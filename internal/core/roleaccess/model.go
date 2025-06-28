package roleaccess

type RoleAccess struct {
	ID          int64  `json:"id"`
	RoleID      int64  `json:"role_id"`
	ResourceKey string `json:"resource_key"`
}

func (RoleAccess) TableName() string {
	return "role_access"
}

type RoleAccessDTO struct {
	RoleID  int64         `json:"roleId"`
	Details []*RoleDetail `json:"details"`
}

type RoleDetail struct {
	ResourceKey string `json:"resourceKey"`
}
