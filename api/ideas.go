package api

import (
	"github.com/JorgenJJ/justice4campus/internal/storage"
	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo/bson"
)

// CreateIdea persist a new Idea
func CreateIdea(c *gin.Context) {
	// Check if user is in room
	//if (storage.Room.IsUserInRoom(getuid(), getroomid())) { }

	idea := storage.IdeaStruct{
		Title:       c.PostForm("ideaTitle"),
		Description: c.PostForm("ideaDescription"),
		RoomID:		 bson.ObjectIdHex(c.PostForm("roomID")),
	}

	// Percist the Idea
	idea, err := storage.Idea.Add(idea, "5be1af1eedaad52b0b94227a")
	if err != nil {
		c.JSON(200, gin.H{"status": "400", "err": err})
		return
	}

	c.JSON(200, gin.H{"status": "success", "message": "created idea", "data": idea})
	//c.Redirect(301, "/room/" + room.ID.Hex())

}
