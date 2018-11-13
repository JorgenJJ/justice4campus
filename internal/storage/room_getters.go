package storage

import (
	"errors"

	mgo "github.com/globalsign/mgo"
	bson "github.com/globalsign/mgo/bson"
)

// Find gets a RoomStruct based on its title
func (db *MongoDBRooms) Find(id string) (RoomStruct, error) {

	var room RoomStruct

	session, err := mgo.Dial(db.HOST.URI)
	if err != nil {
		return room, errors.New("error dialing the database")
	}
	defer session.Close()

	find := bson.M{"$or": []bson.M{bson.M{"_id": bson.ObjectIdHex(id)}, bson.M{"title": id}}}
	//find := bson.M{"title": id}
	err = session.DB(db.HOST.NAME).C(db.COLLECTION).Find(find).One(&room)
	if err != nil {
		return room, errors.New("error finding the document")
	}
	return room, nil
}

// FindAll gets all the public Rooms, meaning their password field are empty
func (db *MongoDBRooms) FindAll() ([]RoomStruct, error) {

	var rooms []RoomStruct

	session, err := mgo.Dial(db.HOST.URI)
	if err != nil {
		return rooms, errors.New("error dialing the database")
	}
	defer session.Close()

	//find := bson.M{"password": ""}
	err = session.DB(db.HOST.NAME).C(db.COLLECTION).Find(nil).Sort("-_id").All(&rooms)
	if err != nil {
		return rooms, errors.New("error finding the document")
	}
	return rooms, nil
}

// FindWithID gets a RoomStruct based on its id
func (db *MongoDBRooms) FindWithID(ideaID string) (RoomStruct, error) {
	var room RoomStruct

	session, err := mgo.Dial(db.HOST.URI)
	if err != nil {
		return room, errors.New("error dialing the database")
	}
	defer session.Close()

	find := bson.M{"_id": bson.ObjectIdHex(ideaID)}
	err = session.DB(db.HOST.NAME).C(db.COLLECTION).Find(find).One(&room)
	if err != nil {
		return room, errors.New("error finding the document")
	}
	return room, nil
}


// Bool to check if user is in room
func (db *MongoDBRooms) IsUserInRoom(uid string, rid string) bool {
	return false
}



// GetIdeaIDs retrives all of the idea ids for a specific room
func (db *MongoDBRooms) GetIdeaIDs(roomID string) ([]string, error) {

	var room RoomStruct
	var ideas []string

	session, err := mgo.Dial(db.HOST.URI)
	if err != nil {
		return ideas, errors.New("error dialing the database")
	}
	defer session.Close()

	find := bson.D{{"_id", bson.ObjectIdHex(roomID)}}
	err = session.DB(db.HOST.NAME).C(db.COLLECTION).Find(find).One(&room)
	if err != nil {
		return ideas, errors.New("error finding idea ids")
	}
	for _, id := range room.IdeaIDs {
		ideas = append(ideas, id)
	}

	return ideas, nil

}
