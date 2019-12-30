package api

import (
	"fmt"
	"net/http"
)

// Home of application page
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Welcome, HomePage</h1>")
}
