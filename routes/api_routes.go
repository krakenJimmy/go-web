package routes

import (
	"samples/web-basic/api"

	"github.com/gorilla/mux"
)

// Router for application
func Router() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/home", api.Home)
	r.HandleFunc("/signin", api.SignIn).Methods("POST")
	r.HandleFunc("/signup", api.SignUp).Methods("POST")

	return r
}
