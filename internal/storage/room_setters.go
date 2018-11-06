package storage

import (
	"errors"

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
	err = session.DB(db.HOST.NAME).C(db.COLLECTION).Insert(&room)
	if err != nil {
		return room, errors.New("error inserting the document")
	}
	return room, nil
}

// AddMember appends another member to the member list of the room
func (db *MongoDBRooms) AddMember(member MemberStruct, roomTitle string) error {

	session, err := mgo.Dial(db.HOST.URI)
	if err != nil {
		return errors.New("error dialing the database")
	}
	defer session.Close()

	find := bson.M{"title": roomTitle}
	member.ID = bson.NewObjectId()
	update := bson.M{"$push": bson.M{"members": member}}
	err = session.DB(db.HOST.NAME).C(db.COLLECTION).Update(find, update)

	if err != nil {
		return errors.New("error finding the document")
	}
	return nil
}

func (db *MongoDBRooms) AddMemberWithPassword(member MemberStruct, roomTitle, roomPassword string) error {

	session, err := mgo.Dial(db.HOST.URI)
	if err != nil {
		return errors.New("error dialing the database")
	}
	defer session.Close()

	find := bson.D{{"title", roomTitle}, {"password", roomPassword}}
	member.ID = bson.NewObjectId()
	update := bson.M{"$push": bson.M{"members": member}}
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
