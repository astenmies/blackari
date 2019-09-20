package core

import (
	"context"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/astenmies/lychee/core/static"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
)

// Graphql middleware renders the handler
func Graphql(schema string, resolvers interface{}) http.HandlerFunc {
	parsedSchema := graphql.MustParseSchema(schema, resolvers)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		relayHandler := relay.Handler{Schema: parsedSchema}
		relayHandler.ServeHTTP(w, r)
	})
}

// Playground middleware renders Graphiql
func Playground() http.HandlerFunc {
	graphiql, _ := static.Asset("static/index.html")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write(graphiql)
	})
}

// GetSchema gets the schema at a provided path
func GetSchema(path string) (string, error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}

	return string(b), nil
}

// GetClient returns a MongoDB Client
func GetClient() (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Connect(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatal("Couldn't connect to the database", err)
	} else {
		log.Println("Connected!")
	}

	return client, nil
}
