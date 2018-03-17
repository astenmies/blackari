// Embedded in this article https://medium.com/p/c98e491015b6
package main

import (
	"fmt"
	"log"
	"net/http"

	mongo "./mongo"
	utils "./utils"
	"github.com/neelance/graphql-go"
	"github.com/neelance/graphql-go/relay"
	"github.com/rs/cors"
	"gopkg.in/mgo.v2/bson"
)

//////// MAIN ////////
func main() {
	// Create a handler for /graphql which passes cors for remote requests
	http.Handle("/graphql", cors.Default().Handler(&relay.Handler{Schema: graphqlSchema}))

	// Write a GraphiQL page to /
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)

	// ListenAndServe starts an HTTP server with a given address and handler.
	log.Fatal(http.ListenAndServe(":8080", nil))
}

//////// GRAPHQL ////////
var graphqlSchema *graphql.Schema

// Schema describes the data that we ask for
var Schema = `
    schema {
        query: Query
    }
    # The Query type represents all of the entry points.
    type Query {
        post(slug: String!): Post
    }
    type Post {
        id: ID!
        slug: String!
        title: String!
    }
    `

//////// INIT ////////
func init() {
	utils.InitViper()

	// MustParseSchema parses a GraphQL schema and attaches the given root resolver.
	// It returns an error if the Go type signature of the resolvers does not match the schema.
	graphqlSchema = graphql.MustParseSchema(Schema, &Resolver{})

	mongo.Dummy()
}

//////// RESOLVER ////////
// In order to respond to queries, a schema needs to have resolve functions for all fields.
// Go’s structs are typed collections of fields. They’re useful for grouping data together to form records.
type Resolver struct{}

type post struct {
	ID    graphql.ID
	Slug  string
	Title string
}

type postResolver struct {
	s *post
}

type searchResultResolver struct {
	result interface{}
}

// Slices can be created with the built-in make function; this is how we create dynamically-sized arrays.
var postData = make(map[string]*post)

// Post resolves the Post queries.
func (r *Resolver) Post(args struct{ Slug string }) *postResolver {

	// One result is a pointer to type post.
	oneResult := &post{}

	// Call mongo Get, session and reference to the post collection
	session, collection := mongo.Get("post")
	// Close the session so its resources may be put back in the pool or collected, depending on the case.
	defer session.Close()

	// Inside the collection, find by slug and return all fields.
	err := collection.Find(bson.M{"slug": args.Slug}).Select(bson.M{}).One(&oneResult)
	if err != nil {
		fmt.Println(err)
	}

	// Make a type postResolver out of oneResult.
	if s := oneResult; s != nil {
		return &postResolver{oneResult}
	}
	return nil
}

// Resolve each field to respond to queries.
func (r *postResolver) ID() graphql.ID {
	return r.s.ID
}

func (r *postResolver) Slug() string {
	return r.s.Slug
}

func (r *postResolver) Title() string {
	return r.s.Title
}
