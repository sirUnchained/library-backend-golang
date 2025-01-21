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
	newUser := &models.UserModel{}
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		sendRes.Send(false, err.Error(), 500, w)
		return
	}

	fmt.Printf("%+v", newUser)
	usersList[strconv.Itoa(newUser.Id)] = *newUser

	sendRes.Send(true, "user registered !", 201, w)
}
