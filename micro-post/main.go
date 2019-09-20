package main

import (
	"log"
	"net/http"

	"github.com/astenmies/lychee/core"
	"github.com/astenmies/lychee/micro-post/resolvers"
	"github.com/astenmies/lychee/micro-post/schema"
)

// https://github.com/graph-gophers/graphql-go/issues/106#issuecomment-350231819
// RootResolver is extended with each "microservice" resolver
type RootResolver struct {
	resolvers.PostResolver
	postDB resolvers.DB
}

// GetSchema returns the schema of Post
func GetSchema() string {
	s, _ := schema.Asset("schema/schema.graphql")
	stringSchema := string(s)

	return stringSchema
}

// Can be run as a microservice or consumed by a central server
func main() {
	s := GetSchema()
	r := &RootResolver{}

	http.Handle("/graphql", core.Graphql(s, r))
	http.Handle("/", core.Playground())

	log.Fatal(http.ListenAndServe(":4444", nil))
}
