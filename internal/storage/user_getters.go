package storage

import (
	"errors"

	mgo "github.com/globalsign/mgo"
	bson "github.com/globalsign/mgo/bson"
)

// FindByID an exisiting user from its username
func (db *MongoDBUsers) FindByID(id string) (UserStruct, error) {

	var user UserStruct

	session, err := mgo.Dial(db.HOST.URI)
	if err != nil {
		return user, errors.New("error dialing the database")
	}
	defer session.Close()

	find := bson.M{"_id": bson.ObjectIdHex(id)}
	err = session.DB(db.HOST.NAME).C(db.COLLECTION).Find(find).One(&user)
	if err != nil {
		return user, errors.New("error finding the document")
	}
	return user, nil
}
