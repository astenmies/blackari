// Embedded in this article https://medium.com/p/c98e491015b6
package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"

	gqlResolver "github.com/astenmies/lychee/server/gqlResolver"
	gqlSchema "github.com/astenmies/lychee/server/gqlSchema"
	mongo "github.com/astenmies/lychee/server/mongo"
	utils "github.com/astenmies/lychee/server/utils"
	graphql "github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"github.com/rs/cors"
	"github.com/spf13/viper"
)

var PublicKey = []byte("secret")

// https://github.com/CallistoM/go-graphql-auth/blob/master/handler/login.go
// https://github.com/lpalmes/graphql-go-introduction/blob/viewer/main.go
// https://medium.com/@matryer/context-keys-in-go-5312346a868d

// type key int

// var id = key(1)

// func Set(ctx context.Context) context.Context {
// 	return context.WithValue(ctx, id, "secret")
// }
// func Get(ctx context.Context) interface{} {
// 	val := ctx.Value("jwt")
// 	return val
// }

func auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		token := r.Header.Get("Authorization")
		jwt, err := utils.CheckToken(token)
		if err != nil {
			fmt.Println(err)
		} else {
			cookie := &http.Cookie{Name: "lychee", Value: token, HttpOnly: false}

			http.SetCookie(w, cookie) // TODO: set the cookie from somewhere useful
		}

		ctx = context.WithValue(ctx, "jwt", jwt)

		// ctx = Set(ctx)
		// spew.Dump(Get(ctx))

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

//////// MAIN ////////
func main() {
	// Create a handler for /graphql which passes cors for remote requests
	http.Handle("/graphql", cors.Default().Handler(auth(&relay.Handler{Schema: gqlSchema.GraphqlSchema})))

	// Write a GraphiQL page to /
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)

	port := viper.GetInt("lychee.server.port")
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
