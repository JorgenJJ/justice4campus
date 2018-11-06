package api

import (
	"github.com/JorgenJJ/justice4campus/internal/storage"
	"github.com/gin-gonic/gin"
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
	creator := storage.MemberStruct {
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



// GetAllPublicRooms finds all rooms that are public available..
func GetAllPublicRooms(c *gin.Context) {

	rooms, err := storage.Room.FindAllPublic()
	if err != nil {
		c.JSON(200, gin.H{"status": "400", "err": err})
		return
	}

	c.JSON(200, gin.H{"status": "success", "rooms": rooms})
}


// AddMemberToRoom appends a new member to an existing room
func AddMemberToRoom(c *gin.Context) {

	member := storage.MemberStruct {
		Name: c.PostForm("nickName"),
	}

	err := storage.Room.AddMemberWithPassword(member, c.PostForm("roomName"), c.PostForm("roomPassword"))
	if err != nil {
		c.JSON(200, gin.H{"status": "400", "err": err})
		return
	}

	room, err := storage.Room.FindWithTitle(c.PostForm("roomName"))
	if err != nil {
		c.JSON(200, gin.H{"status": "400", "err": err})
		return
	}

	c.Redirect(301, "/room/" + room.ID.Hex())
	//c.JSON(200, gin.H{"status": "success", "message": "You are now added to the room"})
}