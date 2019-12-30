package api

import (
	"fmt"
	"net/http"
)

// SignIn User
func SignIn(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Please SignIn")
}

// SignUp User
func SignUp(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Please SignUp")
}
