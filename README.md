# justice4campus  
This is me contributing.  


### Specification:
  1) Create rooms
  - With password for private rooms and without password for public rooms
  2) Create ideas for each room
  3) Vote on each idea
  4) Comment on each idea
  5) User login
  
  
  
## Core API Specification

### GET /

* What: Homepage
* Response type: text/html
* Response code: 200


### GET /host
* What: Form for creating a room
* Response type: text/html
* Response code: 200

### POST /host

* What: Creating a room
* Response type: application/json
* Response code: 200 if everything is OK
* Request body template

```
nickName: "<user name>"
roomName: "<title of room>"
roomPassword: "<empty for public room>"
```

* Response body template

```
{
  "": ""
}
```
### GET /join
* What: Form for joining a room
* Response type: text/html
* Response code: 200

### POST /join

* What: Joining a room
* Response type: application/json
* Response code: 200 if everything is OK
* Request body template

```
nickName: "<user name>"
roomName: "<title of room>",
roomPassword: "<empty for public room>"
```

* Response body template

```
{
  "": ""
}
```
### GET /room/public
* What: Gets all of the public rooms
* Response: application/json
* Response code: 200
```
{
  "rooms": [
    {
      "id": "<id>",
      "creator": {
        "id": "<id>",
        "name": "<user name>"
      },
      "title": "<room title>",
      "password": "",
      "members": [
        {
          "id": "<id>",
          "name": "<user name>"
        },
        {
          "id": "<id>",
          "name": "<user name>"
        }
      ]
    }
  ]
}


```



### GET /room/{id}
* What: Get room
* Response type: text/html
* Response code: 200

### POST /createIdea
* What: Creating a Idea
* Response type: application/json
* Response code: 200 if everything is OK
* Request body template

```
Title: "<Title of Idea>"
Description: "<Description of Idea>"
```

* Response body template

```
{
  "": ""
}
```

## ???? 
### POST /room/{id}/comment

* What: Post action on room, commenting/like/dislike
* Response type: application/json
* Response code: 200 if everything is OK
* Request body template

```
nickName: "<user name>"
comment: "<text>"
```

* Response body template

```
{
  "": ""
}
```

