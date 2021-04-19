package domain

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username" gorm:"unique"`
	Password []byte `json:"password"`
}
