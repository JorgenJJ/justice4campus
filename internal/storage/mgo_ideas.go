package storage

import (
	"fmt"
	"os"

	mgo "github.com/globalsign/mgo"
)

// IdeaStruct is the template for ideas in the database
type IdeaStruct struct {
	Poster string `json:"poser"`
}

// MongoDBIdeas stores the credentials to the database and collection
type MongoDBIdeas struct {
	URI        string
	NAME       string
	COLLECTION string
}

var Ideas IdeaStorage

// IdeaStorage creates interface for main application to do CRUD operations
type IdeaStorage interface {
	Init()
}

// IdeaSetup configures the global connection to the database
func IdeaSetup() {

	// create credential struct for MongoDB database
	Ideas := &MongoDBIdeas{
		URI:        os.Getenv("MONGO_DB_URI"),
		NAME:       os.Getenv("MONGO_DB_NAME"),
		COLLECTION: "ideas",
	}
	Ideas.Init()
}

// Init ensures a collection exists
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
