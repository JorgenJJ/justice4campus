package storage

import bson "github.com/globalsign/mgo/bson"

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

// MemberStruct is a template for any user
type MemberStruct struct {
	ID       bson.ObjectId `json:"-" bson:"_id"`
	Name     string        `json:"name" bson:"name"`
	Password string        `json:"-" bson:"-"`
}

// RoomStruct is the template for Rooms in the database
type RoomStruct struct {
	ID       bson.ObjectId  `json:"-" bson:"_id"`
	Creator  MemberStruct   `json:"creator" bson:"creator"`
	Title    string         `json:"title" bson:"title"`
	Password string         `json:"password" bson:"password"`
	Members  []MemberStruct `json:"members" bson:"members"`
}

// VoteStruct stores the voting data
type VoteStruct struct {
	Likes   int `json:"likes" bson:"likes"`
	Disikes int `json:"dislikes" bson:"dislikes"`
}

// CommentStruct to define the structure of comments
type CommentStruct struct {
	ID      bson.ObjectId `json:"-" bson:"_id"`
	Creator MemberStruct  `json:"creator" bson:"creator"`
	Text    string        `json:"text" bson:"text"`
}

// IdeaStruct is the template for ideas in the database
type IdeaStruct struct {
	ID          bson.ObjectId   `json:"-" bson:"_id"`
	Title       string          `json:"title"`
	Description string          `json:"description" bson:"description"`
	Vote        VoteStruct      `json:"votes" bson:"votes"`
	Comments    []CommentStruct `json:"comments" bson:"comments"`
}
