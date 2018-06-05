package gqlResolver

import (
	"context"
	"errors"

	"github.com/astenmies/lychee/server/model"
	"github.com/astenmies/lychee/server/utils"
)

// IMPORTANT DOCS
// oauth2 official google example https://github.com/golang/oauth2/blob/master/google/example_test.go
// OAuth 2.0 authentication in a Golang CLI https://gist.github.com/marians/3b55318106df0e4e648158f1ffb43d38
var PublicKey = []byte("secret")

var users = []model.User{
	{
		ID:       "1",
		Username: "seb",
		Password: "seb",
	},
	{
		ID:       "2",
		Username: "bb",
		Password: "bb",
	},
}

// UserLogin :
// - Returns nil if no args are provided
// - Inserts a new user into mongodb
// - Resolves the current user if args are provided
func (r *Resolver) UserLogin(ctx context.Context, args *struct {
	Input *model.UserLoginInput
}) (string, error) {
	for _, user := range users {
		if user.Username == args.Input.Username {
			if user.Password == args.Input.Password {
				token, err := utils.GenerateToken(user)
				if err != nil {
					return "", err
				}
				return token, err
			} else {
				return "", errors.New("password is incorrect")
			}
		}
	}

	return "", errors.New("User not found")
}
