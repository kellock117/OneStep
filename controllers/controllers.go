package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"onestep/models"

	"github.com/gorilla/mux"
)

type response struct {
	ID      string `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}

func ApiCreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.UserInfo

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		log.Fatalf("Unable to decode the request body.  %v", err)
	}

	userID := models.CreateUser(user)

	res := response{
		ID:      userID,
		Message: "Stock created successfully",
	}

	// send the response
	json.NewEncoder(w).Encode(res)
}

func ApiGetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := models.GetAllUsers()

	if err != nil {
		log.Fatalf("Unable to get all users. %v", err)
	}

	// send all the users as response
	json.NewEncoder(w).Encode(users)
}

func ApiGetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID := params["id"]

	user, err := models.GetUser(userID)

	if err != nil {
		log.Fatalf("Unable to get such user. %v", err)
	}

	json.NewEncoder(w).Encode(user)
}
