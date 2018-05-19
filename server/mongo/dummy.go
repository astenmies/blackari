package mongo

import (
	"log"

	"gopkg.in/mgo.v2/bson"
)

// Dummy inserts data into mongodb
func Dummy() {
	log.Println("Seeding mock data to MongoDB")
	// Call mongo Get, session and reference to the post collection
	session, collection := Get("post")
	// Close the session so its resources may be put back in the pool or collected, depending on the case.
	defer session.Close()

	// RemoveAll finds all documents matching the provided selector document
	// and removes them from the database. So we make sure the db is empty before inserting mock data.
	Cleanup()

	// The mock data that we insert.
	err := collection.Insert(
		bson.M{"ID": "1", "title": "First post", "slug": "firstpost"},
		bson.M{"ID": "2", "title": "Hello second post", "slug": "second"},
		bson.M{"ID": "3", "title": "Third post", "slug": "thirdpost"},
	)

	if err != nil {
		log.Fatal(err)
	}
	log.Println("Mock data added successfully!")
}

// DummyUser inserts data into mongodb
func DummyUser() {
	log.Println("Seeding mock user data to MongoDB")
	// Call mongo Get, session and reference to the post collection
	session, collection := Get("user")
	// Close the session so its resources may be put back in the pool or collected, depending on the case.
	defer session.Close()

	// RemoveAll finds all documents matching the provided selector document
	// and removes them from the database. So we make sure the db is empty before inserting mock data.
	Cleanup()

	// The mock data that we insert.
	err := collection.Insert(
		bson.M{"ID": "1", "title": "First post", "slug": "first-post"},
	)

	if err != nil {
		log.Fatal(err)
	}
	log.Println("Mock data added successfully!")
}
