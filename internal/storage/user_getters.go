package storage

import (
	"errors"

	mgo "github.com/globalsign/mgo"
	bson "github.com/globalsign/mgo/bson"
)

// FindByID finds an exisiting user
func (db *MongoDBUsers) FindByID(id string) (UserStruct, error) {
	var user UserStruct

	if !bson.IsObjectIdHex(id) {
		return user, errors.New("not a valid id")
	}

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

// FindManyByID do as it implies, finds many users from their ids
func (db *MongoDBUsers) FindManyByID(ids []string) ([]UserStruct, error) {

	var users []UserStruct

	session, err := mgo.Dial(db.HOST.URI)
	if err != nil {
		return users, errors.New("error dialing the database")
	}
	defer session.Close()

	oids := make([]bson.ObjectId, 0)
	for _, id := range ids {
		oids = append(oids, bson.ObjectIdHex(id))
	}

	find := bson.M{"_id": bson.M{"$in": oids}}
	err = session.DB(db.HOST.NAME).C(db.COLLECTION).Find(find).All(&users)
	if err != nil {
		return users, errors.New("error finding user")
	}
	return users, nil
}
