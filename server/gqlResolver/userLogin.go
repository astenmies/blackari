package gqlResolver

import (
	"errors"

	"github.com/astenmies/lychee/server/model"
	"github.com/astenmies/lychee/server/utils"
)

var PublicKey = []byte("secret")

var users = []model.User{
	{
		Username: "seb",
		Password: "seb",
	},
	{
		Username: "bb",
		Password: "bb",
	},
}

// UserLogin :
// - Returns nil if no args are provided
// - Inserts a new user into mongodb
// - Resolves the current user if args are provided
func (r *Resolver) UserLogin(args *struct {
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
