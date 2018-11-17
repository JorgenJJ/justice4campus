package api

import (
	"github.com/JorgenJJ/justice4campus/internal/storage"
	"github.com/gin-gonic/gin"
	"net/http"
)



// CreateUser adds a new user to the database and create as login session by setting the user id cookie
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

	// Set cookie and redirect user to frontpage
	c.SetCookie("uid", user.ID.Hex(), 3600, "/", "", false, false)
	c.Redirect(301, c.GetHeader("Origin"))
}


// UserLogin authenticates a new user login session by setting the user id cookie
func UserLogin(c *gin.Context) {

	// Find user with parsed credentials
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

	// Set cookie and redirect user to frontpage
	c.SetCookie("uid", auth.ID.Hex(), 3600, "/", "", false, false)
	c.Redirect(301, c.GetHeader("Origin"))
}


// UserLogout clears the user id cookie and redirects to frontpage
func UserLogout(c *gin.Context) {

	// Set cookie and redirect user to frontpage
	c.SetCookie("uid", "", 0, "/", "", false, false)
	c.Redirect(301, c.GetHeader("Origin"))
}

func UserAuth(c *gin.Context) {

}



/* HOW TO GET USER ID COOKIE
	cookie, err := c.Request.Cookie("uid")
    if err != nil {
        fmt.Println(err)
    }
	val, _ := url.QueryUnescape(cookie.Value)

*/