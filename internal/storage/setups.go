package storage

import (
	"errors"
	"fmt"
	"os"

	mgo "github.com/globalsign/mgo"
)

// Setup configures the db connection credentials and initializes the database collections
func Setup() error {

	host := MongoDBHost{
		URI:  os.Getenv("MONGO_DB_URI"),
		NAME: os.Getenv("MONGO_DB_NAME"),
	}

	Room = &MongoDBRooms{
		HOST:       host,
		COLLECTION: "rooms",
	}

	Idea = &MongoDBIdeas{
		HOST:       host,
		COLLECTION: "ideas",
	}

	User = &MongoDBUsers{
		HOST:       host,
		COLLECTION: "users",
	}

	// initialize and handle error
	err := Room.Init()
	if err != nil {
		return err
	}

	err = Idea.Init()
	if err != nil {
		return err
	}

	err = User.Init()
	if err != nil {
		return err
	}
	return nil
}

// Init ensures a collection exists
func (db *MongoDBRooms) Init() error {
	fmt.Println("Initializing collection", db.COLLECTION)

	// establish connection to database and close it again when method finishes
	session, err := mgo.Dial(db.HOST.URI)
	if err != nil {
		return errors.New("error dialing the database")
	}
	defer session.Close()

	collExists := false

	// get database names
	names, err := session.DB(db.HOST.NAME).CollectionNames()
	if err != nil {
		return errors.New("error getting db collections")
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
		info := &mgo.CollectionInfo{}
		err = session.DB(db.HOST.NAME).C(db.COLLECTION).Create(info)
		if err != nil {
			return errors.New("error creating db collection")
		}
		fmt.Println("Created collection", db.COLLECTION)
	}
	fmt.Println("Initialized", db.COLLECTION, "collection!")
	return nil
}

// Init ensures a collection exists
func (db *MongoDBIdeas) Init() error {
	fmt.Println("Initializing collection", db.COLLECTION)

	// establish connection to database and close it again when method finishes
	session, err := mgo.Dial(db.HOST.URI)
	if err != nil {
		return errors.New("error dialing the database")
	}
	defer session.Close()

	collExists := false

	// get database names
	names, err := session.DB(db.HOST.NAME).CollectionNames()
	if err != nil {
		return errors.New("error getting db collections")
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
		info := &mgo.CollectionInfo{}
		err = session.DB(db.HOST.NAME).C(db.COLLECTION).Create(info)
		if err != nil {
			return errors.New("error creating db collection")
		}
		fmt.Println("Created collection", db.COLLECTION)
	}
	fmt.Println("Initialized", db.COLLECTION, "collection!")
	return nil
}

// Init ensures a collection exists
func (db *MongoDBUsers) Init() error {
	fmt.Println("Initializing collection", db.COLLECTION)

	// establish connection to database and close it again when method finishes
	session, err := mgo.Dial(db.HOST.URI)
	if err != nil {
		return errors.New("error dialing the database")
	}
	defer session.Close()

	collExists := false

	// get database names
	names, err := session.DB(db.HOST.NAME).CollectionNames()
	if err != nil {
		return errors.New("error getting db collections")
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
		info := &mgo.CollectionInfo{}
		err = session.DB(db.HOST.NAME).C(db.COLLECTION).Create(info)
		if err != nil {
			return errors.New("error creating db collection")
		}

		index := mgo.Index{
			Key:        []string{"name"},
			Unique:     true,
			DropDups:   false,
			Background: true,
			Sparse:     true,
		}
		err := session.DB(db.HOST.NAME).C(db.COLLECTION).EnsureIndex(index)
		if err != nil {
			return errors.New("error indexing db collection")
		}

		fmt.Println("Created collection", db.COLLECTION)
	}
	fmt.Println("Initialized", db.COLLECTION, "collection!")
	return nil
}
