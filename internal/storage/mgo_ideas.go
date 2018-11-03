package storage

import (
	"fmt"
	"os"

	mgo "github.com/globalsign/mgo"
)

type IdeaStruct struct {
	Poster string `json:"poser"`
}

type MongoDBIdeas struct {
	URI        string
	NAME       string
	COLLECTION string
}

// TrackSetup configures the global connection to the database
func DBSetup() {

	// create credential struct for MongoDB database
	Ideas := &MongoDBIdeas{
		URI:        os.Getenv("MONGO_DB_URI"),
		NAME:       os.Getenv("MONGO_DB_NAME"),
		COLLECTION: os.Getenv("ideas"),
	}
	Ideas.Init()
}

// Init ensures a "tracks" collection exists
func (db *MongoDBIdeas) Init() {
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
