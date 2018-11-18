package storage

import (
	"errors"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

// Add inserts a new IdeaStruct in the database collection
func (db *MongoDBIdeas) Add(idea IdeaStruct) (IdeaStruct, error) {

	session, err := mgo.Dial(db.HOST.URI)
	if err != nil {
		return idea, errors.New("error dialing the database")
	}
	defer session.Close()

	idea.ID = bson.NewObjectId()
	idea.HexID = idea.ID.Hex()
	err = session.DB(db.HOST.NAME).C(db.COLLECTION).Insert(&idea)
	if err != nil {
		return idea, errors.New("error inserting the document")
	}
	
	return idea, nil
}


/**
	Add a user vote in the database collection for a given Idea
	Param: ideaId - ID of the idea to be voted on
	Param: vote - 1 if positive vote -1 if negative
 */
func (db *MongoDBIdeas) Vote(ideaId string, vote int) error {

	session, err := mgo.Dial(db.HOST.URI)
	if err != nil {
		return errors.New("error dialing the database")
	}
	defer session.Close()


	find := bson.M{"_id": bson.ObjectIdHex(ideaId)}
	idea := IdeaStruct{}

	if vote == 1 {
		// add 1 to Likes
		newValue := idea.Vote.Likes + 1
		update := bson.M{"$push": bson.M{"likes": newValue}}
		err = session.DB(db.HOST.NAME).C(db.COLLECTION).Update(find, update)
		if err != nil {
			return errors.New("error updating db")
		}
	} else {
		// add 1 to Dislikes
		newValue := idea.Vote.Dislikes + 1
		update := bson.M{"$push": bson.M{"dislikes": newValue}}
		err = session.DB(db.HOST.NAME).C(db.COLLECTION).Update(find, update)
		if err != nil {
			return errors.New("error updating db")
		}
	}
	return nil
}

// Comment a user comment in the database collection for a given Idea
func (db *MongoDBIdeas) Comment(ideaId string, comment CommentStruct) error {

	session, err := mgo.Dial(db.HOST.URI)
	if err != nil {
		return errors.New("error dialing the database")
	}
	defer session.Close()

	find := bson.M{"_id": bson.ObjectIdHex(ideaId)}
	comment.ID = bson.NewObjectId()
	update := bson.M{"$push": bson.M{"comments": comment}}
	err = session.DB(db.HOST.NAME).C(db.COLLECTION).Update(find, update)

	if err != nil {
		return errors.New("error finding the document")
	}
	return nil
}