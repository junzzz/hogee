package main

import (
	"log"
	"net/http"

	"github.com/junzzz/hogee/controllers"
)

func main() {
	http.HandleFunc("/", controllers.HellWorld)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
