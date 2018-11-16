package storage

import (
	"errors"

	mgo "github.com/globalsign/mgo"
	bson "github.com/globalsign/mgo/bson"
)

// Add adds a new user to the database
func (db *MongoDBUsers) Add(user UserStruct) (UserStruct, error) {

	session, err := mgo.Dial(db.HOST.URI)
	if err != nil {
		return user, errors.New("error dialing the database")
	}
	defer session.Close()

	user.ID = bson.NewObjectId()
	err = session.DB(db.HOST.NAME).C(db.COLLECTION).Insert(&user)
	if err != nil {
		return user, errors.New("error inserting the document")
	}

	return user, nil
}
