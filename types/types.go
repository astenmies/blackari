package types

import (
	"github.com/graph-gophers/graphql-go"
	"go.mongodb.org/mongo-driver/mongo"
)

type DB struct {
	*mongo.Client
}

type Node interface {
	ID() graphql.ID
}
