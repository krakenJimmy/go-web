package main

import (
	"log"
	"net/http"
	"samples/web-basic/routes"
)

func main() {
	r := routes.Router()
	log.Fatal(http.ListenAndServe(":8881", r))
}
