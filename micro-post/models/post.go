package models

import "github.com/graph-gophers/graphql-go"

type Post struct {
	ID    graphql.ID `json:"id"        bson:"id,omitempty"`
	Title string     `json:"title"`
}
