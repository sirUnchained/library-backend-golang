package models

type BookUserModel struct {
	Id     int `json:"id"`
	BookId int `json:"book_id"`
	UserId int `json:"user_id"`
}
