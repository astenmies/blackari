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
	// We can then call viper.Get("string") anywhere
	viper.SetConfigName("Config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Fatal error config file: %s \n", err)
	}
}

func main() {

	// 1 Get the global config
	var (
		appName = viper.Get("app-name").(string)
	)

	// 2 Pass the Request
	// The Client will forward all headers set on the
	// initial Request except Authorization and Cookie
	c := blapi.NewClient(http.DefaultClient) // TODO: don't use the default client.

	// 3 Prepare to read
	// root an entry point for all top-level read operations.
	root, err := resolver.NewRoot(c)
	if err != nil {
		log.Fatal(err)
	}

	// 4 Create the request handler; inject dependencies.
	h := handler.GraphQL{
		// Parse and validate schema. Panic if unable to do so.
		Schema:  graphql.MustParseSchema(schema.String(), root),
		Loaders: loader.Initialize(c),
	}

	// 4 Start a small server that reads our "graphiql.html" file and
	// responds with it, so we are able to have our own graphiql
	// https://medium.com/@matryer/the-http-handler-wrapper-technique-in-golang-updated-bc7fbcffa702
	// http://echorand.me/dissecting-golangs-handlerfunc-handle-and-defaultservemux.html
	http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "graphiql.html")
	}))

	// 5 h to handle all requests to /query
	http.Handle("/query", h)

	// 6 Use h as handler for all incoming requests to /graphql
	// We wrap this in cors to attach cross-origin headers to our handler
	http.Handle("/graphql", cors.Default().Handler(h))

	// 7 Start the server by using ListenAndServe and we log if we have any error, hope not!
	fmt.Println(appName, "listening at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
