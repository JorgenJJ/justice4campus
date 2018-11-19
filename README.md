# Justice4camPUS

### Idea 
The ideas is a web-service where users can create either private or public rooms and share ideas with each other. Users inside the room will be able to write ideas and send them to the room and database, where all ideas for that room are stored. Ideas can be up- or downvoted and commented on.

### Specification:
  1) Create rooms
  - With password for private rooms and without password for public rooms
  2) Create ideas for each room
  3) Comment on each idea
  4) User login
  5) Vote on each idea

### "Behind the scenes"
All users, rooms and ideas are persisted in their respective MongoDB collection. Each room object stores `id` references to users and submitted ideas in the room. Each idea and its comments stores the reference aswell. The reference structure allows for non-dublicate storing of e.g. full data of users when they submit a comment.
The user login session is created by setting a cookie with the users id and retrieving it when needed.


## Core API Specification
___
### GET `/`
* What: Homepage
* Response type: text/html
* Response code: 200

### GET `/signup`
* What: Form for creating a room
* Response type: text/html
* Response code: 200

### POST `/user/signup`
* What: Creating a user
* Response type: REDIRECT TO /
* Response code: 200

### POST `/user/signin`
* What: Authorizing a user session
* Response type: REDIRECT TO /
* Response code: 200

### GET `/host`
* What: Form for creating a room
* Response type: text/html
* Response code: 200

### POST `/host`
* What: Creating a room
* Response type: text/html
* Response code: 200 if everything is OK
* Request body template

```
roomName: "<title of room>"
roomPassword: "<empty for public room>"
```

### GET `/join`
* What: List of rooms to join
* Response type: text/html
* Response code: 200


### POST `/join/{room id}`
* What: Joining a room
* Response type: text/html
* Response code: 200 if everything is OK
* Request body template

```
roomName: "<title of room>",
roomPassword: "<empty for public room>"
```

* Response: REDIRECT TO `/room/{room id}`


### GET `/room/{room id}`
* What: Get room
* Response type: text/html
* Response code: 200


### POST `/createIdea`
* What: Creating a Idea
* Response type: application/json
* Response code: 200 if everything is OK
* Request body template

```
Title: "<Title of Idea>"
Description: "<Description of Idea>"
```

* Response: REDIRECT TO `/room/{room id}`


### POST `/comment/{idea id}`

* What: Post action on room, commenting/like/dislike
* Response type: REDIRECT TO `/room/{room id}`
* Response code: 200 if everything is OK
* Request body template

```
commentText: "<text>"
```

