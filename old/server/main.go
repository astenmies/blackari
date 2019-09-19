package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/astenmies/lychee/server/schema"
	graphql "github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"github.com/pkg/errors"
	"github.com/rs/cors"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type DB struct {
	*mongo.Client
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

//GetClient returns a MongoDB Client
func GetClient() (*DB, error) {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Connect(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	return &DB{client}, nil
}

func init() {
	s := schema.GetRootSchema()
	c, _ := GetClient()
	err := c.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatal("Couldn't connect to the database", err)
	} else {
		log.Println("Connected!")
	}

	// MustParseSchema parses a GraphQL schema and attaches the given root resolver.
	// It returns an error if the Go type signature of the resolvers does not match the schema.
	schema.GraphqlSchema = graphql.MustParseSchema(s, &Resolver{db: c}, graphql.UseStringDescriptions())
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
