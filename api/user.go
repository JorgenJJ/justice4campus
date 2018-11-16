package api

import (
	"github.com/JorgenJJ/justice4campus/internal/storage"
	"github.com/gin-gonic/gin"
)

// CreateUser adds a new user
func CreateUser(c *gin.Context) {

	user := storage.UserStruct{
		Name:     c.PostForm("username"),
		Password: c.PostForm("password"),
	}

	// Percist the user
	user, err := storage.User.Add(user)
	if err != nil {
		c.JSON(200, gin.H{"status": "400", "err": err})
		return
	}

	c.JSON(200, gin.H{"status": "success", "message": "created user", "data": user})
}


// UserLogin authenticates a new user
func UserLogin(c *gin.Context) {

	user := storage.UserStruct{
		Name:       c.PostForm("username"),
		Password: 	c.PostForm("password"),
	}

	// Percist the user
	authenticated, err := storage.User.Authenticate(user)
	if err != nil {
		c.JSON(200, gin.H{"authenticated": false})
		return
	}
	
	c.JSON(200, gin.H{"authenticated": authenticated})
}
