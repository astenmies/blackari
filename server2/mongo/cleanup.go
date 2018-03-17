package mongo

import (
	"log"

	"gopkg.in/mgo.v2/bson"
)

// Cleanup will remove all mock data from the database.
func Cleanup() {
	log.Println("Cleaning up MongoDB...")
	session, collection := Get("post")
	defer session.Close()

	_, err := collection.RemoveAll(bson.M{})
	if err != nil {
		log.Fatal(err)
	}
}
