package storage

import (
	bson "github.com/globalsign/mgo/bson"
)

// MongoDBHost stores the credentials to the database and collection
type MongoDBHost struct {
	URI  string
	NAME string
}

// MongoDBRooms stores the collection name
type MongoDBRooms struct {
	HOST       MongoDBHost
	COLLECTION string
}

// MongoDBIdeas stores the collection name
type MongoDBIdeas struct {
	HOST       MongoDBHost
	COLLECTION string
}

// MongoDBUsers stores the collection name
type MongoDBUsers struct {
	HOST       MongoDBHost
	COLLECTION string
}

// UserStruct is a template for any user
type UserStruct struct {
	ID       bson.ObjectId `json:"id" bson:"_id"`
	Name     string        `json:"username" groups:"meta" bson:"name" form:"username" binding:"required"`
	Password string        `json:"-" bson:"password" form:"password" binding:"required"`
}

// RoomStruct is the template for Rooms in the database
type RoomStruct struct {
	ID        bson.ObjectId `json:"id" groups:"meta" bson:"_id"`
	CreatorID string        `json:"-" bson:"creator_id"`
	Creator   UserStruct    `json:"creator" groups:"meta" bson:"-"`
	Title     string        `json:"title" groups:"meta" bson:"title"`
	Password  string        `json:"-" bson:"password"`
	IsPublic  bool          `json:"is_public" groups:"meta" bson:"is_public"`
	MemberIDs []string      `json:"member_ids" bson:"member_ids"`
	Members   []UserStruct  `json:"members"  groups:"meta" bson:"-"`
	IdeaIDs   []string      `json:"idea_ids" bson:"idea_ids"`
	Ideas     []IdeaStruct  `json:"ideas" bson:"-" groups:"meta"`
}

// VoteStruct stores the voting data
type VoteStruct struct {
	Likes   int `json:"likes" bson:"likes"`
	Disikes int `json:"dislikes" bson:"dislikes"`
}

// CommentStruct to define the structure of comments
type CommentStruct struct {
	ID      bson.ObjectId `json:"id" bson:"_id"`
	Creator UserStruct    `json:"creator" bson:"creator"`
	Text    string        `json:"text" bson:"text"`
}

// IdeaStruct is the template for ideas in the database
type IdeaStruct struct {
	ID          bson.ObjectId   `json:"_id" bson:"_id"`
	Title       string          `json:"title"`
	Description string          `json:"description" bson:"description"`
	Vote        VoteStruct      `json:"votes" bson:"votes"`
	Comments    []CommentStruct `json:"comments" bson:"comments"`
	// RoomID      string          `json:"room_id" bson:"_id"`
}
