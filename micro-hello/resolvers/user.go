package resolvers

import (
	"github.com/graph-gophers/graphql-go"
)

var users = map[string]*User{
	"1": {
		id:        "1",
		firstName: "Seb",
	},
}

type User struct {
	id        graphql.ID
	firstName string
}

func (q *Query) AllUsers() []*User {
	// build up a list of all the users
	userSlice := []*User{}

	for _, user := range users {
		userSlice = append(userSlice, user)
	}

	return userSlice
}

func (u *User) ID() graphql.ID {
	return u.id
}

func (u *User) Firstname() string {
	return u.firstName
}

func (n *NodeResolver) ToUser() (*User, bool) {
	user, ok := n.Node.(*User)
	return user, ok
}
