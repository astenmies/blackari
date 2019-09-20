package main

import (
	"fmt"
	"net/http"

	"github.com/astenmies/lychee/core"
	"github.com/astenmies/lychee/micro-post/db"
	"github.com/astenmies/lychee/micro-post/resolvers"
	"github.com/astenmies/lychee/micro-post/schema"
)

// GetSchema returns the schema of Post
func GetSchema() string {
	s, _ := schema.Asset("schema/schema.graphql")
	stringSchema := string(s)

	return stringSchema
}

func main() {
	c, _ := core.GetClient()
	schem := GetSchema()
	database := &db.Services{c}
	res := &resolvers.PostResolver{
		DB: database,
	}

	http.Handle("/graphql", core.Graphql(schem, res))
	http.Handle("/", core.Playground())

	fmt.Println("Starting server on http://localhost:4002")
	err := http.ListenAndServe(":4002", nil)
	if err != nil {
		fmt.Println(err.Error())
	}
}
