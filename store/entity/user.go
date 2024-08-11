package entity

type User struct {
	Base
	Name     string `json:"name" db:"name"`
	Password string `json:"password" db:"password"`
	Email    string `json:"email" db:"email"`
	Token    string `json:"token" db:"token"`
}
