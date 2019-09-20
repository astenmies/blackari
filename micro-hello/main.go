package main

import (
	"fmt"
	"net/http"

	"github.com/astenmies/lychee/core"
	resolvers "github.com/astenmies/lychee/micro-hello/resolvers"
	"github.com/astenmies/lychee/micro-hello/schema"
)

// https://github.com/graph-gophers/graphql-go/issues/106#issuecomment-350231819
// RootResolver is extended with each "microservice" resolver
type RootResolver struct {
	resolvers.HelloResolver
	resolvers.PersonResolver
	resolvers.PostResolver
}

// GetSchema returns the schema of Post
func GetSchema() string {
	// OPTION 1 - string schema
	// s := `
	// 	schema {
	// 			query: Query
	// 	}
	// 	type Query {
	// 			getGreeting: String!
	// 	}
	// `

	// OPTION 2 - Read from file
	// s, _ := core.GetSchema("schema/schema.graphql")

	// OPTION 3 - Use go-bindata
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

	fmt.Println("Starting server on http://localhost:4001")
	err := http.ListenAndServe(":4001", nil)
	if err != nil {
		fmt.Println(err.Error())
	}
}
