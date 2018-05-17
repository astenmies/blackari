package gqlResolver

import (
	"context"
	"errors"

	"github.com/astenmies/lychee/server/model"
	"github.com/astenmies/lychee/server/utils"
	graphql "github.com/neelance/graphql-go"

	jwt "github.com/dgrijalva/jwt-go"
)

type viewerResolver model.ViewerResolver

func (r *Resolver) Viewer(ctx context.Context, args *struct {
	Token *string
}) (*viewerResolver, error) {

	token := ctx.Value("jwt").(*jwt.Token)
	if token == nil && args.Token == nil {
		return nil, errors.New("There needs to be a token in the Authorization header or viewer input")
	}

	if token == nil && args.Token != nil {
		viewerToken, err := utils.CheckToken(*args.Token)
		if err != nil {
			return nil, err
		}

		token = viewerToken
	}

	claims, _ := token.Claims.(jwt.MapClaims)
	id := claims["sub"].(string)

	var user model.User

	for _, u := range users {
		if id == string(u.ID) {
			user = u
		}
	}

	return &viewerResolver{
		User: user,
	}, nil
}

func (v *viewerResolver) ID() graphql.ID {
	return graphql.ID(v.User.ID)
}

func (v *viewerResolver) Username() string {
	return v.User.Username
}
