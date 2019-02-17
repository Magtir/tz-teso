package routes

import (
	"net/http"
	"github.com/gorilla/mux"
	"project/app/Modules/Main/Controllers/MainController"
)

func InitRoute() (router *mux.Router) {
	router = mux.NewRouter()

	router.PathPrefix("/WWW/").Handler(http.StripPrefix("/WWW/", http.FileServer(http.Dir("./src/project/WWW/"))))

	router.HandleFunc("/users", MainController.UsersList)
	router.HandleFunc("/user/{id:[0-9]+}", MainController.UserGet)
	router.HandleFunc("/user/create", MainController.UserCreate)
	router.HandleFunc("/user/edit/{id:[0-9]+}", MainController.UserEdit)
	router.HandleFunc("/user/delete/{id:[0-9]+}", MainController.UserDelete)

	return
}
