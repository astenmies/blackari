package service

import (
	"log"

	"github.com/astenmies/lychee/server/model"
	"github.com/astenmies/lychee/server/mongo"
)

// InsertUser :
// - Defines a pointer to a User with args
// - Appends a user to users
// - Generates a userID
// - Opens a mgo session
// - Inserts a User
// - Closes the mgo session
// - Returns the user that was inserted
func InsertUser(args *struct {
	User *model.UserInput
}) *model.User {

	newUser := &model.User{
		Username: args.User.Username,
		Password: args.User.Password,
	}

	// userID := xid.New()
	// newUser.ID = userID.String()

	// [TODO]: make services independent of collections
	// Maybe wrap them all together with if statements to avoid code repetitiveness
	session, collection := mongo.Get("user")

	defer session.Close()
	err := collection.Insert(newUser)

	if err != nil {
		log.Fatal(err)
	}
	return newUser
}
