package api

import (
	"fmt"
	"github.com/JorgenJJ/justice4campus/internal/storage"
	"github.com/gin-gonic/gin"
	// "github.com/liip/sheriff"
	// "html/template"
	"net/http"
	// "os"
	"net/url"

	//"encoding/json"
)

// CreateRoom persists a new room
func CreateRoom(c *gin.Context) {
	
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
	//c.JSON(200, gin.H{"status": "success", "message": "created room", "data": room})
	c.SetCookie("room_id", room.ID.Hex(), 3600, "/", "", false, false)
	c.Redirect(301, "/room/" + room.ID.Hex())
}


// AddMemberToRoom appends a new member to an existing room
func AddMemberToRoom(c *gin.Context) {

	userCookie, err := c.Request.Cookie("uid")
    if err != nil {
        c.JSON(200, gin.H{"status": "err", "message": err})
	}
	uid, _ := url.QueryUnescape(userCookie.Value)
	rid := c.Param("id")

	err = storage.Room.AddMemberID(uid, rid, c.PostForm("roomPassword"))
	if err != nil {
		c.JSON(200, gin.H{"status": "400", "message": "unauthorized", "redirect": c.GetHeader("Origin") + "/join"})
		return
	}
	c.Redirect(301, "/room/"+ rid)
	//c.JSON(200, gin.H{"status": "success", "message": "You are now added to the room"})
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

	// set data
	room.Creator = creator
	room.Members = members
	room.Ideas = ideas

	// set cookie
	c.SetCookie("room_id", room.ID.Hex(), 3600, "/", "", false, false)
	c.HTML(http.StatusOK, "room.tmpl.html", room)
	
	// c.JSON(200, gin.H{"status": "success", "data": room})
}

// GetAllRooms return a JSON object containing meta data of all public rooms
func GetAllRooms(c *gin.Context) {

	rooms, err := storage.Room.FindAll()
	if err != nil {
		c.JSON(200, gin.H{"status": "400", "err": err})
		return
	}

	for i := 0; i < len(rooms); i++ {
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

	/*
	tpl := template.Must(template.New("joining.tmpl.html").Parse(`{{define "T"}}{{.RoomName}}{{end}}`))


	for _, room := range rooms {
		room.Title = "<div id=\"roomItem\">" +
			"<div id=\"roomItemInfo\">" +
			"<h3 id=\"roomItemName\">" + room.Title +
			"</h3>"

		tmplVars := map[string]interface{} {
			"RoomName": template.HTML("<h1>" + room.Title + "</h1>"),
		}
		tpl.ExecuteTemplate(os.Stdout, "T", tmplVars)
	}
	*/

	// t := template.Must(template.New("joining.tmpl.html").Parse("joining.tmpl.html"))
	// t.Execute(os.Stdout, data)

	//r := Room{"AnimeZone <.<", "Jørgen-san :3"}
	// c.HTML(http.StatusOK, "joining.tmpl.html", nil)
}
