package storage

import (
	"errors"

	mgo "github.com/globalsign/mgo"
	bson "github.com/globalsign/mgo/bson"
)

// FindWithTitle gets a RoomStruct based on its title
func (db *MongoDBRooms) FindWithTitle(title string) (RoomStruct, error) {

	var room RoomStruct

	session, err := mgo.Dial(db.HOST.URI)
	if err != nil {
		return room, errors.New("error dialing the database")
	}
	defer session.Close()

	find := bson.M{"title": title}
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