package storage

import (
	"errors"

	mgo "github.com/globalsign/mgo"
	bson "github.com/globalsign/mgo/bson"
)

// FindByID finds an exisiting user
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

// FindByName finds an exisiting user
func (db *MongoDBUsers) FindByName(name string) (UserStruct, error) {

	var user UserStruct

	session, err := mgo.Dial(db.HOST.URI)
	if err != nil {
		return user, errors.New("error dialing the database")
	}
	defer session.Close()

	find := bson.M{"name": name}
	err = session.DB(db.HOST.NAME).C(db.COLLECTION).Find(find).One(&user)
	if err != nil {
		return user, errors.New("error finding the document")
	}
	return user, nil
}

// FindByCred finds an exisiting user
func (db *MongoDBUsers) FindByCred(user UserStruct) (UserStruct, error) {

	var foundUser UserStruct

	session, err := mgo.Dial(db.HOST.URI)
	if err != nil {
		return foundUser, errors.New("error dialing the database")
	}
	defer session.Close()

	find := bson.M{"$and": []bson.M{bson.M{"name": user.Name}, bson.M{"password": user.Password}}}
	err = session.DB(db.HOST.NAME).C(db.COLLECTION).Find(find).One(&foundUser)

	if err != nil {
		return foundUser, errors.New("error finding user")
	}
	return foundUser, nil
}
