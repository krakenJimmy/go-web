package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"samples/web-basic/models"
)

// AuthResponse is the result of Auth
type AuthResponse struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
	Token   string `json:"token"`
}

// SignIn User with username and password
func SignIn(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Please SignIn")
}

// SignUp User
func SignUp(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	log.Println("User: %+v", user)
	if user.UserName == "" || user.Password == "" {
		authResponse := AuthResponse{
			Errcode: 1000,
			Errmsg:  "username or password can't empty",
		}

		response, _ := json.Marshal(authResponse)
		w.Header().Set("Content-Type", "applicationi/json")
		w.Write(response)
		return
	}

	authResponse := AuthResponse{
		Errcode: 200,
		Errmsg:  "ok",
		Token:   "success",
	}

	response, _ := json.Marshal(authResponse)
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
	return
}
