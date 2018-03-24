package mongo

import (
	"log"
	"time"

	"github.com/spf13/viper"
	mgo "gopkg.in/mgo.v2"
)

// Get returns the session and a reference to the post collection.
func Get(c string) (*mgo.Session, *mgo.Collection) {

	host := viper.GetString("blackari.mongo.host")
	wait := viper.GetDuration("blackari.mongo.maxWait")
	dbName := viper.GetString("blackari.mongo.dbName")

	maxWait := time.Duration(wait * time.Second)
	session, err := mgo.DialWithTimeout(host, maxWait)

	if err != nil {
		log.Fatal(err)
	}

	collection := session.DB(dbName).C(c)

	return session, collection
}
