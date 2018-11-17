package api

import (
	"fmt"
	"github.com/JorgenJJ/justice4campus/internal/storage"
	"github.com/gin-gonic/gin"
	"github.com/liip/sheriff"
	"html/template"
	"net/http"
	"os"
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
	
	rid := c.PostForm("roomID")
	
	fmt.Println("user id", uid, "room id", rid)

	err = storage.Room.AddMemberID(uid, rid, c.PostForm("roomPassword"))
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

	c.HTML(http.StatusOK, "room.tmpl.html", room)
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

func GetAllRooms(c *gin.Context) {
	type Room struct {
		Title	string
		Author	string
	}

	rooms, err := storage.Room.FindAll()
	if err != nil {
		c.JSON(200, gin.H{"status": "400", "err": err})
		return
	}
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

	t := template.Must(template.New("joining.tmpl.html").Parse("joining.tmpl.html"))
	t.Execute(os.Stdout, rooms)
	//r := Room{"AnimeZone <.<", "Jørgen-san :3"}
	c.HTML(http.StatusOK, "joining.tmpl.html", nil)

}
