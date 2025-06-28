package user

type User struct {
	ID               int64
	NationalIDNumber string
	Username         string
	HashPassword     string
	IsAdmin          bool
}

func (User) TableName() string {
	return "users"
}

type UserDTO struct {
	ID               int64  `json:"id"`
	NationalIDNumber string `json:"nationalIdNumber"`
	Username         string `json:"username"`
	IsAdmin          bool   `json:"isAdmin"`
	FirstName        string `json:"firstName"`
	LastName         string `json:"lastName"`
}

type UserSaveReq struct {
	ID               int64  `json:"id"`
	NationalIDNumber string `json:"nationalIdNumber"`
	Username         string `json:"username"`
	Password         string `json:"password"`
	IsAdmin          bool   `json:"isAdmin"`
}
