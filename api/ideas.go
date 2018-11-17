package api

import (
	"github.com/JorgenJJ/justice4campus/internal/storage"
	"github.com/gin-gonic/gin"
	"net/url"
)

// CreateIdea persist a new Idea
func CreateIdea(c *gin.Context) {
	// Check if user is in room
	//if (storage.Room.IsUserInRoom(getuid(), getroomid())) { }

	roomCookie, err := c.Request.Cookie("room_id")
    if err != nil {
        c.JSON(200, gin.H{"status": "err", "message": err})
	}
	roomID, _ := url.QueryUnescape(roomCookie.Value)

	idea := storage.IdeaStruct{
		Title:       c.PostForm("ideaTitle"),
		Description: c.PostForm("ideaDescription"),
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

	c.JSON(200, gin.H{"status": "success", "message": "created idea", "data": idea})
	//c.Redirect(301, "/room/" + room.ID.Hex())
}
