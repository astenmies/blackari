package main

import (
	"log"
	"net/http"

	"github.com/astenmies/lychee/core"
	"github.com/astenmies/lychee/micro-hello/greeting"
)

// https://github.com/graph-gophers/graphql-go/issues/106#issuecomment-350231819
type RootResolver struct {
	greeting.GreetingResolver
}

// GetSchema returns the schema of Post
func GetSchema() string {
	s := `
		schema {
				query: Query
		}
		type Query {
				getGreeting: String!
		}
	`
	return s
}

// Can be run as a microservice or consumed by a central server
func main() {
	s, _ := core.GetSchema("schema/schema.graphql")
	// s := GetSchema()
	r := &RootResolver{}

	http.Handle("/graphql", core.Graphql(s, r))
	http.HandleFunc("/", core.Playground())

	log.Fatal(http.ListenAndServe(":4444", nil))
}
