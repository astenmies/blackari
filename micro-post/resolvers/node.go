package resolvers

import (
	"github.com/astenmies/lychee/micro-post/models"
	"github.com/graph-gophers/graphql-go"
)

type Node interface {
	ID() graphql.ID
}

type NodeResolver struct {
	node Node
}

var users = map[string]*models.User{
	"1": {
		ID:       "1",
		LastName: "Mies",
	},
}

func (n *NodeResolver) ID() graphql.ID {
	return n.node.ID()
}

func (n *NodeResolver) ToUser() (*UserResolver, bool) {
	user, ok := n.node.(*UserResolver)
	return user, ok
}

func (r *Query) Node(args struct{ ID string }) *NodeResolver {
	user := users[args.ID]

	if user != nil {
		return &NodeResolver{
			&UserResolver{*user},
		}
	} else {
		return nil
	}
}
