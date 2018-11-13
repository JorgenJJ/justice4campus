package api

import (
	"fmt"
	"github.com/JorgenJJ/justice4campus/internal/storage"
	"github.com/gin-gonic/gin"
	"github.com/liip/sheriff"
)

// CreateRoom persists a new room
func CreateRoom(c *gin.Context) {
	
	
	// check for existing room
	/*
	found, err := storage.Room.FindWithTitle(c.PostForm("roomName"))
	if err != nil {
		c.JSON(200, gin.H{"status": "400", "err": err})
		return
	}
	if found.ID != "" {
		c.JSON(200, gin.H{"status": "400", "err": "Room with that title already exists"})
		return
	}*/

	// build structs for a new room
	creator := storage.UserStruct {
		Name: c.PostForm("nickName"),
	}

	room := storage.RoomStruct{
		Creator:  creator,
		Title:    c.PostForm("roomName"),
		Password: c.PostForm("roomPassword"),
	}

	// persist the room
	room, err := storage.Room.Add(room)
	if err != nil {
		c.JSON(200, gin.H{"status": "400", "err": err})
		return
	}

	// respond with the new room data
	//c.JSON(200, gin.H{"status": "success", "message": "created room", "data": room})
	c.Redirect(301, "/room/" + room.ID.Hex())
}



// AddMemberToRoom appends a new member to an existing room
func AddMemberToRoom(c *gin.Context) {

	member := storage.UserStruct {
		Name: c.PostForm("nickName"),
	}

	err := storage.Room.AddMember(member, c.PostForm("roomName"), c.PostForm("roomPassword"))
	if err != nil {
		c.JSON(200, gin.H{"status": "400", "err": err})
		return
	}
	c.JSON(200, gin.H{"status": "success", "message": "You are now added to the room"})
}


// GetRoom get a room with id or title
func GetRoom(c *gin.Context) {

	id := c.Param("id")

	if id == "all" {
		GetAllRoomMetas(c)
		return
	}

	room, err := storage.Room.Find(id)
	if err != nil {
		c.JSON(200, gin.H{"status": "400", "err": err})
		fmt.Println(err)
		return
	}
	//c.Redirect(301, "/room/" + room.ID.Hex())
	c.JSON(200, gin.H{"status": "success", "room": room})
}


// GetAllRoomMetas finds all rooms that are public available..
func GetAllRoomMetas(c *gin.Context) {

	rooms, err := storage.Room.FindAll()
	if err != nil {
		c.JSON(200, gin.H{"status": "400", "err": err})
		return
	}

	o := sheriff.Options{
		Groups: []string{"meta"},
	}

	roomMetas, err := sheriff.Marshal(&o, rooms)
	if err != nil {
		panic(err)
	}

	c.JSON(200, gin.H{"status": "success", "rooms": roomMetas})
}
