package storage

import (
	"fmt"

	mgo "github.com/globalsign/mgo"
)

// RoomStruct is the template for Rooms in the database
type RoomStruct struct {
	Poster string `json:"poser"`
}

// MongoDBRooms stores the credentials to the database and collection
type MongoDBRooms struct {
	URI        string
	NAME       string
	COLLECTION string
}

// Rooms variable for other packages to interact with the storage
var Rooms RoomStorage

// RoomStorage creates interface for main application to do CRUD operations
type RoomStorage interface {
	Init()
}

// Init ensures a collection exists
func (db *MongoDBRooms) Init() {
	fmt.Println("Initializing collection", db.COLLECTION)

	// establish connection to database and close it again when method finishes
	session, err := mgo.Dial(db.URI)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	collExists := false

	// get database names
	names, err := session.DB(db.NAME).CollectionNames()
	if err != nil {
		panic(err)
	}

	// check if collection name exists in database
	for _, name := range names {
		if name == db.COLLECTION {
			collExists = true
		}
	}

	// if not then create a new empty
	if !collExists {
		fmt.Println("No", db.COLLECTION, "collection found! Creating one...")
		info := &mgo.CollectionInfo{
			Capped:         false,
			DisableIdIndex: false,
			ForceIdIndex:   false,
		}
		err = session.DB(db.NAME).C(db.COLLECTION).Create(info)
		if err != nil {
			panic(err)
		}
		fmt.Println("Created collection", db.COLLECTION)
	}
	fmt.Println("Initialized", db.COLLECTION, "collection!")
}
