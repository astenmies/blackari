package main

import (
	"fmt"
	"log"
	"net/http"

	blapi "./blapi"
	handler "./handler"
	loader "./loader"
	resolver "./resolver"
	schema "./schema"
	graphql "github.com/neelance/graphql-go"
	"github.com/rs/cors"
	"github.com/spf13/viper"
)

// Documentation on how to print http requests
// https://medium.com/doing-things-right/pretty-printing-http-requests-in-golang-a918d5aaa000

// This function runs at the start of the program
func init() {

	// Initialize viper
	// After this, we can call viper.Get("string") anywhere
	viper.SetConfigName("Config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Fatal error config file: %s \n", err)
	}
}

func main() {

	var (
		appName = viper.Get("app-name").(string)
	)

	c := blapi.NewClient(http.DefaultClient) // TODO: don't use the default client.

	root, err := resolver.NewRoot(c)
	if err != nil {
		log.Fatal(err)
	}

	// Create the request handler; inject dependencies.
	h := handler.GraphQL{
		// Parse and validate schema. Panic if unable to do so.
		Schema:  graphql.MustParseSchema(schema.String(), root),
		Loaders: loader.Initialize(c),
	}

	// We will start a small server that reads our "graphiql.html" file and
	// responds with it, so we are able to have our own graphiql
	http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "graphiql.html")
	}))

	http.Handle("/query", h)

	// This is where our graphql server is handled, we declare "/graphql" as the route
	// where all our graphql requests will be directed to
	http.Handle("/graphql", cors.Default().Handler(h))
	// We start the server by using ListenAndServe and we log if we have any error, hope not!
	fmt.Println(appName, "listening at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
