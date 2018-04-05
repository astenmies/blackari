package gqlResolver

import (
	"github.com/astenmies/lychee/server/model"
	"github.com/astenmies/lychee/server/service"
	"github.com/astenmies/lychee/server/utils"
)

type userResolver model.UserResolver

// CreateUser :
// - Returns nil if no args are provided
// - Inserts a new user into mongodb
// - Appends the user to users
// - Resolves the current user if args are provided
func (r *Resolver) CreateUser(args *struct {
	User *model.UserInput
}) *userResolver {

	// If all fields are nil
	// "createUser": null (in graphql)
	if utils.AllNil(args.User) {
		return nil
	}

	result := service.InsertUser(args)

	// Make a type PostResolver out of result.
	if s := result; s != nil {
		return &userResolver{result}
	}

	return nil
}

// func (r *userResolver) ID() string {
// 	return r.R.ID
// }

func (r *userResolver) Username() *string {
	return r.R.Username
}

func (r *userResolver) Password() *string {
	return r.R.Password
}
