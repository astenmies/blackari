package main

import (
	"fmt"
	"net/http"

	"github.com/astenmies/lychee/core"
	"github.com/astenmies/lychee/micro-post/db"
	"github.com/astenmies/lychee/micro-post/resolvers"
	"github.com/astenmies/lychee/micro-post/schema"
)

func main() {
	c, _ := core.GetClient()
	schem := schema.GetSchema()
	database := &db.Services{c}

	r := &resolvers.Query{
		DB: database,
	}

	http.Handle("/graphql", core.Graphql(schem, r))
	http.Handle("/", core.Playground())

	fmt.Println("Starting server on http://localhost:4002")
	err := http.ListenAndServe(":4002", nil)
	if err != nil {
		fmt.Println(err.Error())
	}
}
