package routes

import (
	"encoding/json"
	"net/http"
	"strconv"

	models "../models"
	util "../utils"
)

//GET

func getUsersPaginationHandler(w http.ResponseWriter, req *http.Request) {
	page, ok := req.URL.Query()["page"]
	var usersListReturn util.UserList_JSON_Pagination
	var usersList []util.User
	if !ok || len(page[0]) < 1 {
		usersList = models.GetUsersPagination(0) //Devolvemos primera pagina
	} else {
		pageNumber, err := strconv.Atoi(page[0])
		usersListReturn.Page = pageNumber
		usersListReturn.BeforePage = pageNumber - 1
		usersListReturn.NextPage = pageNumber + 1
		if err != nil {
			usersList = models.GetUsersPagination(0) //Devolvemos primera pagina
		} else {
			usersList = models.GetUsersPagination(pageNumber)
		}
	}
	usersListReturn.UserList = usersList

	js, err := json.Marshal(usersListReturn)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)

}

//DELETE

func deleteUserHandler(w http.ResponseWriter, req *http.Request) {
	var user util.User_id_json
	json.NewDecoder(req.Body).Decode(&user)
	result := models.DeleteUser(user.Id)
	if result == true {
		util.PrintLog("Usuario con ID '" + strconv.Itoa(user.Id) + "' borrado.")
	} else {
		util.PrintLog("Error borrando usuario con ID '" + strconv.Itoa(user.Id) + "'")
	}
}
