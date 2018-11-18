package api

import (
	"fmt"
	"github.com/JorgenJJ/justice4campus/internal/storage"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/url"
	"text/template"
)

// CreateRoom persists a new room
func CreateRoom(c *gin.Context) {
	
	// get user id
	cookie, err := c.Request.Cookie("uid")
    if err != nil {
		fmt.Println(err)
		c.JSON(200, gin.H{"status": "no valid user session", "message": "could not create room"})
		return
    }
	uid, _ := url.QueryUnescape(cookie.Value)
	
	// persist the room
	room, err := storage.Room.Add(storage.RoomStruct{
		CreatorID:  uid,
		Title:    c.PostForm("roomName"),
		Password: c.PostForm("roomPassword"),
		IsPublic: c.PostForm("roomPassword") == "",
	})

	if err != nil {
		c.JSON(200, gin.H{"status": "400", "err": err})
		return
	}

	// respond with the new room data
	c.SetCookie("room_id", room.ID.Hex(), 3600, "/", "", false, false)
	c.Redirect(301, "/room/" + room.ID.Hex())
}


// AddMemberToRoom appends a new member to an existing room
func AddMemberToRoom(c *gin.Context) {

	// get user and room id
	userCookie, err := c.Request.Cookie("uid")
    if err != nil {
        c.JSON(200, gin.H{"status": "err", "message": err})
	}
	uid, _ := url.QueryUnescape(userCookie.Value)
	rid := c.Param("id")

	// try to add member to room
	err = storage.Room.AddMemberID(uid, rid, c.PostForm("roomPassword"))
	if err != nil {
		c.JSON(200, gin.H{"status": "400", "message": "unauthorized", "redirect": c.GetHeader("Origin") + "/join"})
		return
	}
	c.Redirect(301, "/room/"+ rid)
}


// GetRoom get a room with id or title
func GetRoom(c *gin.Context) {
	
	id := c.Param("id")

	// fetch Room
	room, err := storage.Room.Find(id)
	if err != nil {
		c.JSON(200, gin.H{"status": "400", "err": err})
		fmt.Println(err)
		return
	}

	// fetch Full data
	creator, _ := storage.User.FindByID(room.CreatorID)
	members, _ := storage.User.FindManyByID(room.MemberIDs)
	ideas, _ := storage.Idea.FindManyByID(room.IdeaIDs)


	// Find all creator IDs
	creatorIDs := make([]string, 0)
	for _, idea := range ideas {
		creatorIDs = append(creatorIDs, idea.CreatorID)
		for _, comment := range idea.Comments {
			creatorIDs = append(creatorIDs, comment.CreatorID)
		}
	}
	creators, _ := storage.User.FindManyByID(creatorIDs)
	for _, creator := range creators {
		for i := range ideas {
			if ideas[i].CreatorID == creator.ID.Hex() {
				ideas[i].Creator = creator
			}
			for j := range ideas[i].Comments {
				if ideas[i].Comments[j].CreatorID == creator.ID.Hex() {
					ideas[i].Comments[j].Creator = creator
				}
				
			}
		}
	}

/*
	// match and append creators
	for i := range ideas {
		for j := range ideas[i].Comments {
			for _, creator := range creators {
				if ideas[i].Comments[j].CreatorID == creator.ID.Hex() {
					ideas[i].Comments[j].Creator = creator
				}
				if ideas[i].CreatorID == creator.ID.Hex() {
					ideas[i].Creator = creator
				}
			}
		}
	}
	*/
	/*
	// loop through all idea comments and find it's creator (definitely very effecient ( ͡° ͜ʖ ͡°) )
	for _, idea := range ideas {
		for i := range idea.Comments {
			idea.Comments[i].Creator, _ = storage.User.FindByID(idea.Comments[i].CreatorID)
			if err != nil {
				idea.Comments[i].Creator = storage.UserStruct{
					Name: "Not found",
				}	
			}
		}	
	}
	*/

	// set data
	room.Creator = creator
	room.Members = members
	room.Ideas = ideas

	// set cookie
	c.SetCookie("room_id", room.ID.Hex(), 3600, "/", "", false, false)
	c.HTML(http.StatusOK, "room.tmpl.html", room)
}

// GetAllRooms return a JSON object containing meta data of all public rooms
func GetAllRooms(c *gin.Context) {

	rooms, err := storage.Room.FindAll()
	if err != nil {
		c.JSON(200, gin.H{"status": "400", "err": err})
		return
	}

	for i := range rooms {
		rooms[i].Creator, err = storage.User.FindByID(rooms[i].CreatorID)
		if err != nil {
			c.JSON(200, gin.H{"status": "400", "err": err})
			return
		}
		rooms[i].Members, err = storage.User.FindManyByID(rooms[i].MemberIDs)
		if err != nil {
			c.JSON(200, gin.H{"status": "400", "err": err})
			return
		}
	}

	c.HTML(http.StatusOK, "joining.tmpl.html", gin.H{
		"rooms": rooms,
	})
}
