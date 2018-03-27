package service

import (
	"log"

	"github.com/astenmies/lychee/server/model"
	"github.com/astenmies/lychee/server/mongo"
	"gopkg.in/mgo.v2/bson"
)

// FindOneBySlug :
// - Opens a mgo session
// - Finds the post and injects it in result
// - Closes the session so its resources may be put back in the pool or be collected (depending on the case)
// - Returns the result
func FindOneBySlug(slug string) *model.Post {

	result := &model.Post{}

	session, collection := mongo.Get("post")
	defer session.Close()

	err := collection.Find(bson.M{"slug": slug}).Select(bson.M{}).One(&result)
	if err != nil {
		log.Fatal(err)
	}

	return result
}
