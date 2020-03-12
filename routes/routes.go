package routes

import (
	"fmt"
	"net/http"

	util "../utils"
	"github.com/gorilla/mux"
)

type Page struct {
	Title string
	Body  string
}

func LoadRouter(port string) {
	router := mux.NewRouter()

	//STATIC RESOURCES
	http.Handle("/", router)

	//LOGIN
	router.HandleFunc("/login", loginUserHandler).Methods("POST")
	router.HandleFunc("/register", registerUserHandler).Methods("POST")

	//USER(GLOBAL)
	router.HandleFunc("/user/adminG/userList", getUsersPaginationHandler).Methods("GET")
	router.HandleFunc("/user/delete", deleteUserHandler).Methods("DELETE")

	if port == "" {
		port = "5001"
	}
	fmt.Println("Servidor escuchando en el puerto ", port)
	//err := http.ListenAndServeTLS(":" + port, "cert.crt", "key.key", nil)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		util.PrintErrorLog(err)
	}
}
