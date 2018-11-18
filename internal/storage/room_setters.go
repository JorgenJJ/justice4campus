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
	room.HexID = room.ID.Hex()
	err = session.DB(db.HOST.NAME).C(db.COLLECTION).Insert(&room)
	if err != nil {
		return room, errors.New("error inserting the document")
	}
	return room, nil
}

// AddMemberID appends another member to the member list of the room
func (db *MongoDBRooms) AddMemberID(id, roomID, roomPassword string) error {

	session, err := mgo.Dial(db.HOST.URI)
	if err != nil {
		return errors.New("error dialing the database")
	}
	defer session.Close()

	hasMember, err := Room.HasMember(id, roomID)
	if err != nil {
		return errors.New("error finding the document")
	}

	if !hasMember {
		// update room with new member id
		find := bson.M{"_id": bson.ObjectIdHex(roomID), "password": roomPassword}
		update := bson.M{"$push": bson.M{"member_ids": id}}

		err = session.DB(db.HOST.NAME).C(db.COLLECTION).Update(find, update)
		if err != nil {
			return errors.New("error finding the document")
		}
	}
	return nil
}

// HasMember checks if room has member
func (db *MongoDBRooms) HasMember(memberID, roomID string) (bool, error) {

	hasMember := false

	session, err := mgo.Dial(db.HOST.URI)
	if err != nil {
		return hasMember, errors.New("error dialing the database")
	}
	defer session.Close()

	var room RoomStruct

	// search query for room
	find := bson.M{"_id": bson.ObjectIdHex(roomID)}
	err = session.DB(db.HOST.NAME).C(db.COLLECTION).Find(find).One(&room)

	for _, _memberID := range room.MemberIDs {
		if _memberID == memberID {
			hasMember = true
			break
		}
	}

	if err != nil {
		return hasMember, errors.New("error finding the document")
	}
	return hasMember, nil
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

	session, err := mgo.Dial(db.HOST.URI)
	if err != nil {
		return errors.New("error dialing the database")
	}
	defer session.Close()

	find := bson.M{"_id": bson.ObjectIdHex(roomID)}
	update := bson.M{"$push": bson.M{"idea_ids": ideaID}}
	err = session.DB(db.HOST.NAME).C(db.COLLECTION).Update(find, update)

	if err != nil {
		return errors.New("error finding the document")
	}
	return nil
}
