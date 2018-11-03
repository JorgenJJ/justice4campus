package storage

import "os"

// Setup runs the setup for each collection
func Setup() {
	IdeaSetup()
	RoomSetup()
}

// IdeaSetup configures the global connection to the database
func IdeaSetup() {

	// create credential struct for MongoDB database
	Ideas := &MongoDBIdeas{
		URI:        os.Getenv("MONGO_DB_URI"),
		NAME:       os.Getenv("MONGO_DB_NAME"),
		COLLECTION: "ideas",
	}
	defer Ideas.Init()
}

// RoomSetup configures the global connection to the database
func RoomSetup() {

	// create credential struct for MongoDB database
	Rooms := &MongoDBRooms{
		URI:        os.Getenv("MONGO_DB_URI"),
		NAME:       os.Getenv("MONGO_DB_NAME"),
		COLLECTION: "rooms",
	}
	defer Rooms.Init()
}
