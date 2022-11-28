package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"onestep/models"
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
		Message: "User created successfully",
	}

	// send the response
	json.NewEncoder(w).Encode(res)
}

func ApiLogin(w http.ResponseWriter, r *http.Request) {
	var user models.UserInfo

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		log.Fatalf("Unable to decode the request body.  %v", err)
	}

	cookie, err := models.Login(user)

	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
	} else {
		http.SetCookie(w, cookie)
		http.Redirect(w, r, "/", 302)
	}
}

func ApiLogout(w http.ResponseWriter, r *http.Request) {
	cookie := &http.Cookie{
		Name:   "session",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}

	http.SetCookie(w, cookie)
	http.Redirect(w, r, "/", 302)
}
