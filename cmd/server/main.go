package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	storage "github.com/JorgenJJ/justice4campus/internal/storage"
	"github.com/gorilla/mux"
)

func main() {
	// get application port from OS for app to listen on
	port := os.Getenv("PORT")
	if port == "" {
		fmt.Println("$PORT must be set")
	}

	storage.Setup()

	r := mux.NewRouter()

	fmt.Print("RUNNING")
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
