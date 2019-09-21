package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/davecgh/go-spew/spew"
	"github.com/nautilus/gateway"
	"github.com/nautilus/graphql"
	"github.com/vektah/gqlparser/ast"
)

func main() {
	viewerField := &gateway.QueryField{
		Name: "viewer",
		Type: ast.NamedType("String", &ast.Position{}),
		Resolver: func(ctx context.Context, args map[string]interface{}) (string, error) {
			// for now just return the value in context
			spew.Dump(ctx)
			return "Hello", nil
			// return ctx.Value("user-id").(string), nil
		},
	}

	// introspect the apis
	schemas, err := graphql.IntrospectRemoteSchemas(
		"http://localhost:4001/graphql",
		"http://localhost:4002/graphql",
	)
	if err != nil {
		panic(err)
	}

	// create the gateway instance
	gw, err := gateway.New(schemas, gateway.WithQueryFields(viewerField))
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/graphql", gw.GraphQLHandler)
	http.HandleFunc("/", gw.PlaygroundHandler)

	fmt.Println("Starting server on http://localhost:4000")
	err = http.ListenAndServe(":4000", nil)
	if err != nil {
		fmt.Println(err.Error())
	}
}
