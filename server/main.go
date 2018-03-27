// Embedded in this article https://medium.com/p/c98e491015b6
package main

import (
	"log"
	"net/http"
	"strconv"

	gqlResolver "github.com/astenmies/blackari/server/gqlResolver"
	gqlSchema "github.com/astenmies/blackari/server/gqlSchema"
	mongo "github.com/astenmies/blackari/server/mongo"
	utils "github.com/astenmies/blackari/server/utils"
	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"github.com/rs/cors"
	"github.com/spf13/viper"
)

//////// MAIN ////////
func main() {
	// Create a handler for /graphql which passes cors for remote requests
	http.Handle("/graphql", cors.Default().Handler(&relay.Handler{Schema: gqlSchema.GraphqlSchema}))

	// Write a GraphiQL page to /
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)

	port := viper.GetInt("blackari.server.port")
	goPort := ":" + strconv.Itoa(port) // Needs ":1234" as port
	// ListenAndServe starts an HTTP server with a given address and handler.
	log.Fatal(http.ListenAndServe(goPort, nil))
}

//////// INIT ////////
func init() {
	// Init global config
	utils.InitViper()

	// MustParseSchema parses a GraphQL schema and attaches the given root resolver.
	// It returns an error if the Go type signature of the resolvers does not match the schema.
	gqlSchema.GraphqlSchema = graphql.MustParseSchema(gqlSchema.GetRootSchema(), &gqlResolver.Resolver{})

	// Insert dummy data into mongodb
	mongo.Dummy()
}
