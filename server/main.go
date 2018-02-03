package main

import (
	"log"
	"net/http"

	"io/ioutil"

	"fmt"

	"github.com/neelance/graphql-go"
	"github.com/neelance/graphql-go/relay"
	"github.com/rs/cors"
)

// This will be use by our handler at /graphql
var schema *graphql.Schema

// Documentation on how to print http requests
// https://medium.com/doing-things-right/pretty-printing-http-requests-in-golang-a918d5aaa000

// This function runs at the start of the program
func init() {

	// We get the schema from the file, rather than having the schema inline here
	// I think will lead to better organizaiton of our own code
	schemaFile, err := ioutil.ReadFile("schema.graphql")
	if err != nil {
		// We will panic if we don't find the schema.graphql file in our server
		panic(err)
	}

	// We will use graphql-go library to parse our schema from "schema.graphql"
	// and the resolver is our struct that should fullfill everything in the Query
	// from our schema
	schema, err = graphql.ParseSchema(string(schemaFile), &Resolver{})
	if err != nil {
		panic(err)
	}
}

func main() {
	// We will start a small server that reads our "graphiql.html" file and
	// responds with it, so we are able to have our own graphiql
	http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write(page)
	}))

	http.Handle("/query", &relay.Handler{Schema: schema})

	// This is where our graphql server is handled, we declare "/graphql" as the route
	// where all our graphql requests will be directed to
	// http.Handle("/graphql", &relay.Handler{Schema: schema})
	http.Handle("/graphql", cors.Default().Handler(&relay.Handler{Schema: schema}))

	// We start the server by using ListenAndServe and we log if we have any error, hope not!
	fmt.Println("Listening at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// Resolver struct is the main resolver  wich we will use to fullfill
// queries and mutations that our schema.graphql defines
type Resolver struct{}

// Hello function resolves to hello: String! in the Query object in our schema.graphql
func (r *Resolver) Hello() string {
	return "world"
}

var page = []byte(`
	<!DOCTYPE html>
	<html>
		<head>
			<link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/graphiql/0.10.2/graphiql.css" />
			<script src="https://cdnjs.cloudflare.com/ajax/libs/fetch/1.1.0/fetch.min.js"></script>
			<script src="https://cdnjs.cloudflare.com/ajax/libs/react/15.5.4/react.min.js"></script>
			<script src="https://cdnjs.cloudflare.com/ajax/libs/react/15.5.4/react-dom.min.js"></script>
			<script src="https://cdnjs.cloudflare.com/ajax/libs/graphiql/0.10.2/graphiql.js"></script>
		</head>
		<body style="width: 100%; height: 100%; margin: 0; overflow: hidden;">
			<div id="graphiql" style="height: 100vh;">Loading...</div>
			<script>
				function graphQLFetcher(graphQLParams) {
					return fetch("/query", {
						method: "post",
						body: JSON.stringify(graphQLParams),
						credentials: "include",
					}).then(function (response) {
						return response.text();
					}).then(function (responseBody) {
						try {
							return JSON.parse(responseBody);
						} catch (error) {
							return responseBody;
						}
					});
				}
				ReactDOM.render(
					React.createElement(GraphiQL, {fetcher: graphQLFetcher}),
					document.getElementById("graphiql")
				);
			</script>
		</body>
	</html>
	`)
