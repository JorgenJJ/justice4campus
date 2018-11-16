package api

import (
	"github.com/JorgenJJ/justice4campus/internal/storage"
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
	"net/url"
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

	c.SetCookie("uid", user.ID.Hex(), 3600, "/", "", false, false)
	c.JSON(200, gin.H{"status": "success", "message": "created user", "data": user, "cookie set": user.ID.Hex()})
}


// UserLogin authenticates a new user login session
func UserLogin(c *gin.Context) {

	auth, _ := storage.User.FindByCred(storage.UserStruct{
		Name:       c.PostForm("username"),
		Password: 	c.PostForm("password"),
	})

	var form storage.UserStruct

	// This will infer what binder to use depending on the content-type header.
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	if form.Name != auth.Name || form.Password != auth.Password {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
		return
	}

	c.SetCookie("uid", auth.ID.Hex(), 3600, "/", "", false, false)
	c.JSON(http.StatusOK, gin.H{"status": "authorized", "cookie set": auth.ID.Hex()})
}


// UserLogout authenticates a new user
func UserLogout(c *gin.Context) {

	cookie, err := c.Request.Cookie("uid")
    if err != nil {
        fmt.Println(err)
    }
	uid, _ := url.QueryUnescape(cookie.Value)
	
	c.SetCookie("uid", "", 0, "/", "", false, false)
	c.JSON(200, gin.H{"status": "not implemented", "cleared cookie": uid})	
}


/*

	cookie, err := c.Request.Cookie("uid")
    if err != nil {
        fmt.Println(err)
    }
	val, _ := url.QueryUnescape(cookie.Value)
	

*/