package core

import (
	"context"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/astenmies/lychee/core/static"
	"github.com/astenmies/lychee/types"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
)

func Graphql(s string, r interface{}) http.HandlerFunc {
	schema := graphql.MustParseSchema(s, r)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		relayHandler := relay.Handler{Schema: schema}
		relayHandler.ServeHTTP(w, r)
	})
}

func Playground() http.HandlerFunc {
	graphiql, _ := static.Asset("static/index.html")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write(graphiql)
	})
}

func GetSchema(path string) (string, error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}

	return string(b), nil
}

//GetClient returns a MongoDB Client
func GetClient() (*types.DB, error) {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Connect(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	return &types.DB{client}, nil
}
