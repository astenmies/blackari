package service

import (
	"log"

	"github.com/davecgh/go-spew/spew"

	"github.com/astenmies/lychee/server/utils"
	"gopkg.in/mgo.v2/bson"

	"github.com/rs/xid"
	"github.com/astenmies/lychee/server/model"
	"github.com/astenmies/lychee/server/mongo"
)

// UserFindByUsername :
// - Opens a mgo session
// - Finds the user and injects it in result
// - Closes the session so its resources may be put back in the pool or be collected (depending on the case)
// - Returns the result
func UserFindByUsername(username string) *model.User {

	result := &model.User{}

	session, collection := mongo.Get("user")
	defer session.Close()
	err := collection.Find(bson.M{"username": username}).Select(bson.M{}).One(&result)
	spew.Dump(username)

	if err != nil {
		log.Fatal(err)
	}

	return result
}

// UserLogin :
// - Defines a pointer to a User with args
// - Appends a user to users
// - Generates a userID
// - Opens a mgo session
// - Inserts a User
// - Closes the mgo session
// - Returns the user that was inserted
func UserLogin(args struct {
	Username string
	Password string
}) *model.User {

	newUser := &model.User{
		Username: args.Username,
		Password: utils.HashAndSalt(args.Password),
	}

	userID := xid.New()
	newUser.ID = userID.String()

	session, collection := mongo.Get("user")
	spew.Dump(session)

	defer session.Close()
	err := collection.Insert(newUser)

	if err != nil {
		log.Fatal(err)
	}
	return newUser
}

// UserInsert :
// - Defines a pointer to a User with args
// - Appends a user to users
// - Generates a userID
// - Opens a mgo session
// - Inserts a User
// - Closes the mgo session
// - Returns the user that was inserted
func UserInsert(args struct {
	Username string
	Password string
}) *model.User {

	newUser := &model.User{
		Username: args.Username,
		Password: utils.HashAndSalt(args.Password),
	}

	// userID := xid.New()
	// newUser.ID = userID.String()

	session, collection := mongo.Get("user")

	defer session.Close()
	err := collection.Insert(newUser)

	if err != nil {
		log.Fatal(err)
	}
	return newUser
}
