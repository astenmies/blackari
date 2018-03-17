package mongo

import (
	"log"

	"gopkg.in/mgo.v2/bson"
)

// Cleanup will remove all mock data from the database.
func Cleanup() {
	log.Println("Cleaning up MongoDB...")

	// The collections to cleanup
	collections := []string{"post"}

	// Cleanup each collection
	for _, coll := range collections {

		session, collection := Get(coll)
		defer session.Close()

		// Remove everything inside the collection
		_, err := collection.RemoveAll(bson.M{})

		if err != nil {
			log.Fatal(err)
		}

		// Confirm
		log.Println("Cleaned collection", coll)
	}
}
