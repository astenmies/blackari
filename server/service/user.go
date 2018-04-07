package service

import (
	"log"

	"github.com/astenmies/lychee/server/utils"

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
func InsertUser(args struct {
	Username string
	Password string
}) *model.User {

	log.Println(utils.HashAndSalt(args.Password))

	newUser := &model.User{
		Username: args.Username,
		Password: args.Password,
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
