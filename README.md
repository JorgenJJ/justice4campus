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
{
  "creator": "<user name>",
  "title": "<title of room>",
  "password": "<empty for public room>"
}
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
{
  "title": "<title of room>",
  "password": "<empty for public room>"
}
```

* Response body template

```
{
  "": ""
}
```
### GET /room/{id}
* What: Get room
* Response type: text/html
* Response code: 200


### POST /room/{id}/comment

* What: Post action on room, commenting/like/dislike
* Response type: application/json
* Response code: 200 if everything is OK
* Request body template

```
{
  "creator": "<user name>"
  "comment": "<text>"
}
```

* Response body template

```
{
  "": ""
}
```

