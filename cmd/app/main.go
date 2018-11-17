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

	router.GET("/", api.Index)

	
	// USERS
	router.POST("/user/signup", api.CreateUser)
	router.POST("/user/signin", api.UserLogin)
	router.POST("/user/signout", api.UserLogout)
	
	router.GET("user/signin", func(c *gin.Context) {
		c.HTML(http.StatusOK, "signin.tmpl.html", nil)
	})
	router.GET("user/signup", func(c *gin.Context) {
		c.HTML(http.StatusOK, "signup.tmpl.html", nil)
	})


	// ROOMS
	router.GET("/host", func(c *gin.Context) {
		c.HTML(http.StatusOK, "hosting.tmpl.html", nil)
	})

	router.GET("/join", api.GetAllRooms)
	router.POST("/host", api.CreateRoom)
	router.GET("/room/:id", api.GetRoom)
	router.POST("/join", api.AddMemberToRoom)
	//router.GET("/room/all", api.GetAllRoomMetas)



	// IDEAS
	router.POST("/createIdea", api.CreateIdea)


	router.Run(":" + port)
}
