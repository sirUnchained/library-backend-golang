package models

type UserModel struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Role     string `json:"role"`
}
