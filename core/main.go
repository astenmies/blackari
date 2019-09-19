package core

import (
	"io/ioutil"
	"net/http"

	"github.com/astenmies/lychee/core/static"

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
