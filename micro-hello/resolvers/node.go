package resolvers

import (
	"github.com/astenmies/lychee/types"
	"github.com/graph-gophers/graphql-go"
)

type Query struct{}

type NodeResolver struct {
	types.Node
}

func (q *Query) Node(args struct{ ID string }) *NodeResolver {
	user := users[args.ID]

	if user != nil {
		return &NodeResolver{user}
	} else {
		return nil
	}
}

func (n *NodeResolver) ID() graphql.ID {
	return n.Node.ID()
}
