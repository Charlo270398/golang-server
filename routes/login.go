package routes

import (
	"encoding/json"
	"net/http"

	models "../models"
	util "../utils"
)

type JSON_Return struct {
	Result string
	Error  string
}

//POST

type JSON_Credentials struct {
	Password string `json:"password"`
	Email    string `json:"email"`
}

//POST
func loginUserHandler(w http.ResponseWriter, req *http.Request) {
	util.PrintLog("Intentando iniciar sesión...")
	var creds JSON_Credentials
	// Get the JSON body and decode into credentials
	err := json.NewDecoder(req.Body).Decode(&creds)
	if err != nil {
		// If the structure of the body is wrong, return an HTTP error
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	//COMPROBAMOS USER Y PASS
	jsonReturn := JSON_Return{"", ""}
	correctLogin := models.LoginUser(creds.Email, creds.Password)
	if correctLogin == true {
		jsonReturn = JSON_Return{"OK", ""}
	} else {
		jsonReturn = JSON_Return{"", "Usuario y contraseña incorrectos"}
	}
	js, err := json.Marshal(jsonReturn)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

//POST
func registerUserHandler(w http.ResponseWriter, req *http.Request) {
	var user util.User_json
	json.NewDecoder(req.Body).Decode(&user)
	inserted, err := models.InsertUser(user)
	jsonReturn := JSON_Return{"", ""}
	if inserted == true {
		jsonReturn = JSON_Return{"OK", ""}
	} else {
		jsonReturn = JSON_Return{"", "El usuario no se ha podido registrar"}
	}

	js, err := json.Marshal(jsonReturn)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
