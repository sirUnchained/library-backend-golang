package controller

import (
	"encoding/json"
	"fmt"
	"libBackend/models"
	"libBackend/utils"
	"net/http"
	"strconv"
)

var allRelations = map[string]models.BookUserModel{}

func GetAll(w http.ResponseWriter, r *http.Request) {
	utils.Send(true, allRelations, 200, w)
}

func Create(w http.ResponseWriter, r *http.Request) {
	// var newRelation *models.BookUserModel
	newRelation := &models.BookUserModel{}

	err := json.NewDecoder(r.Body).Decode(newRelation)
	if err != nil {
		if err.Error() == "EOF" {
			utils.Send(false, "invalid fields.", 400, w)
		} else {
			fmt.Println(err)
			utils.Send(false, err.Error(), 500, w)
		}
		return
	}

	allRelations[strconv.Itoa(newRelation.Id)] = *newRelation
	fmt.Printf("%+v\n", newRelation)

	utils.Send(true, "realation added added !", 201, w)
}
