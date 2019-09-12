package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/astenmies/lychee/server/schema"
	graphql "github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"github.com/pkg/errors"
	"github.com/rs/cors"
	"gopkg.in/mgo.v2"
)

type DB struct {
	DB *mgo.Database
}

type Resolver struct {
	db *DB
}

func main() {
	http.Handle("/graphql", cors.Default().Handler(&relay.Handler{Schema: schema.GraphqlSchema}))

	// Write a GraphiQL page to /
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)

	port := 4444
	goPort := ":" + strconv.Itoa(port)
	log.Fatal(http.ListenAndServe(goPort, nil))
}

// https://labix.org/mgo
func newDB(path string, name string) (*DB, error) {
	session, err := mgo.Dial(path)
	if err != nil {
		panic(err)
	}
	defer session.Close()
	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)
	db := session.DB(name)
	return &DB{db}, nil
}

func init() {
	s := schema.GetRootSchema()
	db, _ := newDB("localhost", "lychee")

	// MustParseSchema parses a GraphQL schema and attaches the given root resolver.
	// It returns an error if the Go type signature of the resolvers does not match the schema.
	schema.GraphqlSchema = graphql.MustParseSchema(s, &Resolver{db: db}, graphql.UseStringDescriptions())
}

func gqlIDP(id uint) *graphql.ID {
	r := graphql.ID(fmt.Sprint(id))
	return &r
}

func gqlIDToUint(i graphql.ID) (uint, error) {
	r, err := strconv.ParseInt(string(i), 10, 32)
	if err != nil {
		return 0, errors.Wrap(err, "GqlIDToUint")
	}

	return uint(r), nil
}
