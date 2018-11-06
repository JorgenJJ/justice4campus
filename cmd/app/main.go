package main

import (
	"github.com/JorgenJJ/justice4campus/api"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"

	"github.com/JorgenJJ/justice4campus/internal/storage"
)

func main() {

	err := storage.Setup()
	if err != nil {
		panic(err) // should eventually be handled gracefully
	}


	// get application port from OS for app to listen on
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

	router.GET("/host", func(c *gin.Context) {
		c.HTML(http.StatusOK, "hosting.tmpl.html", nil)
	})

	router.GET("/join", func(c *gin.Context) {
		c.HTML(http.StatusOK, "joining.tmpl.html", nil)
	})
	router.POST("/join", api.AddMemberToRoom)


	router.GET("/room/public", api.GetAllPublicRooms)

	router.POST("/host", api.CreateRoom)
	

	router.Run(":" + port)
}