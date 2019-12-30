package routes

import (
	"samples/web-basic/api"

	"github.com/gorilla/mux"
)

func routes() {
	r := mux.NewRouter()

	r.HandleFunc("/signIn", api.SignIn).Methods("POST")
	r.HandleFunc("/signUp", api.SignUp).Methods("POST")
}
