package controller

import (
	"encoding/json"
	"fmt"
	models "libBackend/models"
	sendRes "libBackend/utils"
	"net/http"
	"strconv"
)

var usersList = map[string]models.UserModel{}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")
	if token != "1111" {
		sendRes.Send(false, "Unauthorized", 401, w)
		return
	}

	sendRes.Send(true, usersList, 200, w)

}

func Register(w http.ResponseWriter, r *http.Request) {
	// * way 1
	// var newUser *models.UserModel
	// err := json.NewDecoder(r.Body).Decode(&newUser)
	// * way 2
	newUser := &models.UserModel{}
	err := json.NewDecoder(r.Body).Decode(newUser)

	if err != nil {
		if err.Error() == "EOF" {
			sendRes.Send(false, "invalid fields.", 400, w)
		} else {
			fmt.Println(err)
			sendRes.Send(false, err.Error(), 500, w)
		}
		return
	}

	usersList[strconv.Itoa(newUser.Id)] = *newUser
	fmt.Printf("%+v\n", usersList)

	sendRes.Send(true, "user registered !", 201, w)
}
