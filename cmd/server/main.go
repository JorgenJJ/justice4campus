package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	// get application port from OS for app to listen on
	port := os.Getenv("PORT")
	if port == "" {
		fmt.Println("$PORT must be set")
	}

	r := mux.NewRouter()

	fmt.Print("RUNNING")
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
