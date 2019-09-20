package main

import (
	"log"
	"net/http"

	"github.com/astenmies/lychee/core"
	"github.com/astenmies/lychee/micro-post/db"
	"github.com/astenmies/lychee/micro-post/resolvers"
	"github.com/astenmies/lychee/micro-post/schema"
)

// https://github.com/graph-gophers/graphql-go/issues/106#issuecomment-350231819
// RootResolver is extended with each "microservice" resolver
// type RootResolver struct {
// 	*resolvers.PostResolver
// }

// GetSchema returns the schema of Post
func GetSchema() string {
	s, _ := schema.Asset("schema/schema.graphql")
	stringSchema := string(s)

	return stringSchema
}

func main() {

	c, _ := core.GetClient()

	schem := GetSchema()
	deb := &resolvers.DB{c}
	res := &resolvers.PostResolver{DB: deb}

	http.Handle("/graphql", core.Graphql(schem, res))
	http.Handle("/", core.Playground())

	database := db.DB{}
	g := resolvers.GG{
		DB: &database,
	}

	log.Println(g.Greet())

	log.Fatal(http.ListenAndServe(":4444", nil))
}
