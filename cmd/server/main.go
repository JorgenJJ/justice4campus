package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.LoadHTMLGlob("templates/*.tmpl.html")
	router.Static("/static", "static")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl.html", nil)
	})

	router.Run(":" + port)
	/*
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
	*/
}
