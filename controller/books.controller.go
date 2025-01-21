package controller

import (
	"encoding/json"
	"fmt"
	"libBackend/models"
	"libBackend/utils"
	"net/http"
	"strconv"
)

var booksList = map[string]models.BookModel{}

func GetAllBooks(w http.ResponseWriter, r *http.Request) {
	utils.Send(true, booksList, 200, w)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	var newBook *models.BookModel
	err := json.NewDecoder(r.Body).Decode(&newBook)

	if err != nil {
		if err.Error() == "EOF" {
			utils.Send(false, "invalid fields.", 400, w)
		} else {
			fmt.Println(err)
			utils.Send(false, err.Error(), 500, w)
		}
		return
	}

	booksList[strconv.Itoa(newBook.Id)] = *newBook
	fmt.Printf("%+v\n", booksList)

	utils.Send(true, "book added !", 201, w)
}
