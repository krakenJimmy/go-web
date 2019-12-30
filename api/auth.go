package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"samples/web-basic/models"

	"github.com/dgrijalva/jwt-go"

	"samples/web-basic/tools"

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
	zap.S().Infow("signup user is", "info", user, "error", err, "ip", tools.GetIP(r.RemoteAddr))

	if user.UserName == "" || user.Password == "" {
		authResponse := AuthResponse{
			Errcode: 1000,
			Errmsg:  "username or password can't empty",
		}

		json.NewEncoder(w).Encode(authResponse)
		return
	}

	var jwtKey = []byte("tRx10l%W7sU$1A5A%")

	type MyClaims struct {
		UserID   int    `json:"user_id"`
		UserName string `json:"username"`
		jwt.StandardClaims
	}

	claims := MyClaims{
		user.ID,
		user.UserName,
		jwt.StandardClaims{
			ExpiresAt: 15000,
			Issuer:    "bude",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		zap.S().Infow("jwt result", "error", err, "token", tokenString)
	}

	authResponse := AuthResponse{
		Errcode: 200,
		Errmsg:  "ok",
		Token:   tokenString,
	}

	json.NewEncoder(w).Encode(authResponse)
	return
}
