package resolvers

import (
	"github.com/astenmies/lychee/micro-post/models"
	"github.com/graph-gophers/graphql-go"
)

type UserResolver struct {
	m models.User
}

func (u *UserResolver) ID() graphql.ID {
	return u.m.ID
}

func (u *UserResolver) LastName() string {
	return u.m.LastName
}
