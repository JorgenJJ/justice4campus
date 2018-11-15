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

	find := bson.M{"$or": []bson.M{bson.M{"_id": bson.ObjectIdHex(id)}, bson.M{"title": id}}}
	err = session.DB(db.HOST.NAME).C(db.COLLECTION).Find(find).One(&user)

	user.ID = bson.NewObjectId()
	err = session.DB(db.HOST.NAME).C(db.COLLECTION).Insert(&user)
	if err != nil {
		return user, errors.New("error inserting the document")
	}
	return user, nil
}
