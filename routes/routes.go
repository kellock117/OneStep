package routes

import (
	"net/http"
	"onestep/controllers"

	"github.com/gorilla/mux"
)

var Routes = func(router *mux.Router) {
	fileServer := http.FileServer(http.Dir("static"))
	router.Handle("/", fileServer)
	router.HandleFunc("/GetAllUsers", controllers.ApiGetAllUsers).Methods("GET")
	router.HandleFunc("/CreateUser", controllers.ApiCreateUser).Methods("POST")
	router.HandleFunc("/GetUser/{id}", controllers.ApiGetUser).Methods("GET")
}
