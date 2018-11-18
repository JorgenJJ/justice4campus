package api

import (
	"github.com/JorgenJJ/justice4campus/internal/storage"
	"github.com/gin-gonic/gin"
	"net/url"
)

// CreateIdea persist a new Idea
func CreateIdea(c *gin.Context) {

	// get user and room id
	roomCookie, err := c.Request.Cookie("room_id")
    if err != nil {
        c.JSON(200, gin.H{"status": "err", "message": err})
	}
	roomID, _ := url.QueryUnescape(roomCookie.Value)

	userCookie, err := c.Request.Cookie("uid")
    if err != nil {
        c.JSON(200, gin.H{"status": "err", "message": err})
	}
	uid, _ := url.QueryUnescape(userCookie.Value)
	
	// create Idea entity
	idea := storage.IdeaStruct{
		Title:       c.PostForm("ideaTitle"),
		Description: c.PostForm("ideaDescription"),
		CreatorID: uid,
	}

	// Percist the Idea
	idea, err = storage.Idea.Add(idea)
	if err != nil {
		c.JSON(200, gin.H{"status": "400", "err": err})
		return
	}

	// Reference Idea ID in Room object
	err = storage.Room.AddIdeaID(roomID, idea.ID.Hex())
	if err != nil {
		c.JSON(200, gin.H{"status": "400", "err": err})
		return
	}

	// c.JSON(200, gin.H{"status": "success", "message": "created idea", "data": idea})
	c.Redirect(301, "/room/" + roomID)
}


// InsertComment inserts a new comment to the Idea
func InsertComment(c *gin.Context) {

	// get user id
	userCookie, err := c.Request.Cookie("uid")
    if err != nil {
        c.JSON(200, gin.H{"status": "err", "message": err})
	}
	uid, _ := url.QueryUnescape(userCookie.Value)

	// create comment entity
	comment := storage.CommentStruct {
		CreatorID: uid,
		Text: c.PostForm("commentText"),
	}

	// persist comment
	err = storage.Idea.Comment(c.Param("idea_id"), comment)
	if err != nil {
		c.JSON(200, gin.H{"status": "400", "err": err})
		return
	}

	// refresh page
	roomCookie, err := c.Request.Cookie("room_id")
    if err != nil {
        c.JSON(200, gin.H{"status": "err", "message": err})
	}
	roomID, _ := url.QueryUnescape(roomCookie.Value)
	c.Redirect(301, "/room/" + roomID)
}
