package userrole

import (
	"backend/internal/core/dataviews"
	"backend/internal/core/user"
)

func GetUserListDataview() *dataviews.DataviewModel[user.UserDTO, any] {
	return &dataviews.DataviewModel[user.UserDTO, any]{
		Query:       user.AllUsersQuery,
		DataviewKey: "userrole_users",
	}
}
