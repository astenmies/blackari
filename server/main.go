// Embedded in this article https://medium.com/p/c98e491015b6
package main

import (
	"context"
	"encoding/json"
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
// https://jacobmartins.com/2016/02/29/getting-started-with-oauth2-in-go/
type Handler struct {
	Schema *graphql.Schema
}

type test_struct struct {
	Query string
}

func parseRequest(w http.ResponseWriter, r *http.Request) {

	////////////////0//////////////////
	decoder := json.NewDecoder(r.Body)

	var t test_struct
	err := decoder.Decode(&t)

	if err != nil {
		panic(err)
	}

	fmt.Println(t.Query)

	// ////////////////1/////////////////

	// buf, bodyErr := ioutil.ReadAll(r.Body)
	// if bodyErr != nil {
	// 	log.Print("bodyErr ", bodyErr.Error())
	// 	http.Error(w, bodyErr.Error(), http.StatusInternalServerError)
	// 	return
	// }

	// rdr1 := ioutil.NopCloser(bytes.NewBuffer(buf))
	// rdr2 := ioutil.NopCloser(bytes.NewBuffer(buf))
	// log.Printf("BODY: %q", rdr1)

	// r.Body = rdr2
	// // fmt.Fprintf(w, "%q", dump)

	// ///////////////////////2///////////////////
	// // Read the content
	// var bodyBytes []byte
	// if r.Body != nil {
	// 	bodyBytes, _ = ioutil.ReadAll(r.Body)
	// }
	// // Restore the io.ReadCloser to its original state
	// r.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
	// // Use the content
	// bodyString := string(bodyBytes)
	// log.Println(bodyString)

}

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		ctx := r.Context()
		token := r.Header.Get("Authorization")

		parseRequest(w, r)

		jwt, err := utils.CheckToken(token)
		if err != nil {
			fmt.Println(err)
		} else {
			// WARNING: the token was also added within index.html to simulate
			// requests with an Authorization header
			cookie := &http.Cookie{Name: "lychee", Value: token, HttpOnly: false}
			http.SetCookie(w, cookie) // TODO: set the cookie from somewhere more useful
			// https://upgear.io/blog/golang-tip-wrapping-http-response-writer-for-middleware/
		}

		ctx = context.WithValue(ctx, "jwt", jwt)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

//////// MAIN ////////
func main() {
	// Create a handler for /graphql which passes cors for remote requests
	http.Handle("/graphql", cors.Default().Handler(authMiddleware(&relay.Handler{Schema: gqlSchema.GraphqlSchema})))

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
