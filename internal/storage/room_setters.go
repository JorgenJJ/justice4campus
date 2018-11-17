package storage

import (
	"errors"
	"fmt"

	mgo "github.com/globalsign/mgo"
	bson "github.com/globalsign/mgo/bson"
)

// Add inserts a new RoomStruct in the database collection
func (db *MongoDBRooms) Add(room RoomStruct) (RoomStruct, error) {

	session, err := mgo.Dial(db.HOST.URI)
	if err != nil {
		return room, errors.New("error dialing the database")
	}
	defer session.Close()

	room.ID = bson.NewObjectId()
	room.Creator.ID = bson.NewObjectId()
	//room.IsPublic = room.Password == ""
	err = session.DB(db.HOST.NAME).C(db.COLLECTION).Insert(&room)
	if err != nil {
		return room, errors.New("error inserting the document")
	}
	return room, nil
}

// AddMember appends another member to the member list of the room
func (db *MongoDBRooms) AddMemberID(id, roomID, roomPassword string) error {

	session, err := mgo.Dial(db.HOST.URI)
	if err != nil {
		return errors.New("error dialing the database")
	}
	defer session.Close()

	find := bson.D{{"_id", bson.ObjectIdHex(roomID)}, {"password", roomPassword}}
	update := bson.M{"$push": bson.M{"members": id}}
	err = session.DB(db.HOST.NAME).C(db.COLLECTION).Update(find, update)

	if err != nil {
		return errors.New("error finding the document")
	}
	return nil
}

// DeleteWithTitle removes a specific room from the database
func (db *MongoDBRooms) DeleteWithTitle(title string) error {
	session, err := mgo.Dial(db.HOST.URI)
	if err != nil {
		return errors.New("error dialing the database")
	}
	defer session.Close()

	find := bson.M{"title": title}
	err = session.DB(db.HOST.NAME).C(db.COLLECTION).Remove(find)
	if err != nil {
		return errors.New("error finding the document")
	}
	return nil
}

// AddIdeaID appens a new id to a idea
func (db *MongoDBRooms) AddIdeaID(roomID, ideaID string) error {

	fmt.Println(roomID, ideaID)

	session, err := mgo.Dial(db.HOST.URI)
	if err != nil {
		return errors.New("error dialing the database")
	}
	defer session.Close()

	find := bson.D{{"_id", bson.ObjectIdHex(roomID)}}
	update := bson.M{"$push": bson.M{"idea_ids": ideaID}}
	err = session.DB(db.HOST.NAME).C(db.COLLECTION).Update(find, update)

	if err != nil {
		return errors.New("error finding the document")
	}
	return nil
}
