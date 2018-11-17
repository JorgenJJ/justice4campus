package storage

import (
	"errors"

	mgo "github.com/globalsign/mgo"
	bson "github.com/globalsign/mgo/bson"
)

// Find a specific idea from id
func (db *MongoDBIdeas) Find(id string) (IdeaStruct, error) {
	var idea IdeaStruct

	session, err := mgo.Dial(db.HOST.URI)
	if err != nil {
		return idea, errors.New("error dialing the database")
	}
	defer session.Close()

	find := bson.M{"_id": bson.ObjectIdHex(id)}
	err = session.DB(db.HOST.NAME).C(db.COLLECTION).Find(find).One(&idea)
	if err != nil {
		return idea, errors.New("error finding the document")
	}
	return idea, nil
}

// FindManyByID finds all ideas with matching ids
func (db *MongoDBIdeas) FindManyByID(ids []string) ([]IdeaStruct, error) {

	var ideas []IdeaStruct
	session, err := mgo.Dial(db.HOST.URI)
	if err != nil {
		return ideas, errors.New("error dialing the database")
	}
	defer session.Close()

	// convert all strings to bson object ids
	oids := make([]bson.ObjectId, 0)
	for _, id := range ids {
		oids = append(oids, bson.ObjectIdHex(id))
	}

	find := bson.M{"_id": bson.M{"$in": oids}}
	err = session.DB(db.HOST.NAME).C(db.COLLECTION).Find(find).All(&ideas)
	if err != nil {
		return ideas, errors.New("error finding the document")
	}
	return ideas, nil
}
