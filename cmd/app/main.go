package main

import (
	"github.com/JorgenJJ/justice4campus/api"
	"github.com/JorgenJJ/justice4campus/internal/storage"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
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
	router.LoadHTMLGlob("templates/*")
	router.Static("/static", "static")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl.html", nil)
	})

	router.GET("/host", func(c *gin.Context) {
		c.HTML(http.StatusOK, "hosting.tmpl.html", nil)
	})

	router.GET("/join", api.GetAllRooms)

	router.POST("/join", api.AddMemberToRoom)

	router.GET("/room/all", api.GetAllRoomMetas)

	router.POST("/host", api.CreateRoom)
	router.Run(":" + port)
}
