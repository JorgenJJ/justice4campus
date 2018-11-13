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

/*
// FindWithID gets a RoomStruct based on its title
func (db *MongoDBRooms) FindWithID(id string) (RoomStruct, error) {

	var room RoomStruct

	session, err := mgo.Dial(db.HOST.URI)
	if err != nil {
		return room, errors.New("error dialing the database")
	}
	defer session.Close()

	err = session.DB(db.HOST.NAME).C(db.COLLECTION).FindId(id).One(&room)
	if err != nil {
		return room, errors.New("error finding the document")
	}
	return room, nil
}
*/


// Bool to check if user is in room
func (db *MongoDBRooms) IsUserInRoom(uid string, rid string) bool {
	return false
}
