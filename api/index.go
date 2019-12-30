package api

import (
	"encoding/json"
	"net/http"

	"github.com/dgrijalva/jwt-go"

	"go.uber.org/zap"
)

type JWTRequest struct {
	Token string `json:"token"`
}

// Home of application page
func Home(w http.ResponseWriter, r *http.Request) {

	logger, _ := zap.NewDevelopment()
	zap.ReplaceGlobals(logger)

	var jwtRequest JWTRequest
	err := json.NewDecoder(r.Body).Decode(&jwtRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	tokenStr := jwtRequest.Token

	// Initialize a new instance of `Claims`
	claims := &MyJWTClaims{}

	// Parse the JWT string and store the result in `claims`.
	// Note that we are passing the key in this method as well.  This method will return an error
	// if the token is invalid (if it has expired according to the expiry time we set on sign in),
	// or if the signature does not match
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("tRx10l%W7sU$1A5A%"), nil
	})

	zap.S().Infow("jwt claims", "token", token, "err", err)

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		authResponse := AuthResponse{
			Errcode: 4001,
			Errmsg:  "tokan invalid",
		}

		json.NewEncoder(w).Encode(authResponse)
		return
	}

	if !token.Valid {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	authResponse := AuthResponse{
		Errcode: 200,
		Errmsg:  "ok",
	}

	m := make(map[string]interface{})
	m["username"] = claims.UserName
	m["id"] = claims.UserID

	authResponse.Data = m

	json.NewEncoder(w).Encode(authResponse)
	return
}
