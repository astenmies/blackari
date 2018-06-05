package gqlResolver

import (
	"github.com/astenmies/lychee/server/model"
	"github.com/astenmies/lychee/server/service"
)

type userResolver model.UserResolver

// Post :
// - Calls the service
// - Returns the result as a PostResolver
func (r *Resolver) User(args struct{ Username string }) *userResolver {

	result := service.UserFindByUsername(args.Username)

	// Make a type PostResolver out of result.
	if s := result; s != nil {
		return &userResolver{result}
	}

	return nil
}

// UserCreate :
// - Returns nil if no args are provided
// - Inserts a new user into mongodb
// - Resolves the current user if args are provided
func (r *Resolver) UserCreate(args struct {
	Username string
	Password string
}) *userResolver {
	result := service.UserInsert(args)

	// Make a type PostResolver out of result.
	if s := result; s != nil {
		return &userResolver{result}
	}

	return nil
}

func (r *userResolver) Username() string {
	return r.R.Username
}

func (r *userResolver) Password() string {
	return r.R.Password
}
