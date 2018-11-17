package api

import (
	"github.com/JorgenJJ/justice4campus/internal/storage"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/url"
	"fmt"
)



// Index serves the main page with user session status
func IndexPage(c *gin.Context) {

	// run "water fall" to check if user has valid user id cookie
	cookie, err := c.Request.Cookie("uid")
	if err != nil {
		c.HTML(http.StatusOK, "index.tmpl.html",  gin.H{
			"user_present" : false,
			"user_data" : nil,
		})
		fmt.Println("No cookie")
		return
	}

	uid, err := url.QueryUnescape(cookie.Value)
	if err != nil || uid == "" {
		c.HTML(http.StatusOK, "index.tmpl.html",  gin.H{
			"user_present" : false,
			"user_data" : nil,
		})
		fmt.Println("Cookie parsing error")
		return
	}

	user, err := storage.User.FindByID(uid)
	if err != nil || user.ID == "" {
		c.HTML(http.StatusOK, "index.tmpl.html",  gin.H{
			"user_present" : false,
			"user_data" : nil,
		})
		fmt.Println("User found")
		return
	}

	fmt.Println("Found user", user)
	c.HTML(http.StatusOK, "index.tmpl.html",  gin.H{
		"user_present" : true,
		"user_data" : user,
	})
}


/* HOW TO GET USER ID COOKIE
	cookie, err := c.Request.Cookie("uid")
    if err != nil {
        fmt.Println(err)
    }
	val, _ := url.QueryUnescape(cookie.Value)

*/