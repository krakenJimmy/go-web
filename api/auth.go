package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"samples/web-basic/models"

	"go.uber.org/zap"
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

	logger, _ := zap.NewDevelopment()
	zap.ReplaceGlobals(logger)
	zap.S().Infow("signup user is", "info", user, "error", err)

	if user.UserName == "" || user.Password == "" {
		authResponse := AuthResponse{
			Errcode: 1000,
			Errmsg:  "username or password can't empty",
		}

		json.NewEncoder(w).Encode(authResponse)
		return
	}

	authResponse := AuthResponse{
		Errcode: 200,
		Errmsg:  "ok",
		Token:   "success",
	}

	json.NewEncoder(w).Encode(authResponse)
	return
}
