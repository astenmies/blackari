// Embedded in this article https://medium.com/p/c98e491015b6
package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"

	gqlResolver "github.com/astenmies/lychee/server/gqlResolver"
	gqlSchema "github.com/astenmies/lychee/server/gqlSchema"
	mongo "github.com/astenmies/lychee/server/mongo"
	utils "github.com/astenmies/lychee/server/utils"
	"github.com/dghubble/gologin/google"
	"github.com/fatih/color"
	"github.com/gorilla/sessions"
	graphql "github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"github.com/rs/cors"
	"github.com/spf13/viper"
	"golang.org/x/oauth2"
	googleOAuth2 "golang.org/x/oauth2/google"
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

		// Read the content
		buf, bodyErr := ioutil.ReadAll(r.Body)
		if bodyErr != nil {
			log.Print("bodyErr ", bodyErr.Error())
			http.Error(w, bodyErr.Error(), http.StatusInternalServerError)
			return
		}

		rdr1 := ioutil.NopCloser(bytes.NewBuffer(buf))
		rdr2 := ioutil.NopCloser(bytes.NewBuffer(buf))
		// Restore the to its original state
		r.Body = rdr2
		// manipulate rd1 only
		// log.Printf("BODY: %q", rdr1)
		decoder := json.NewDecoder(rdr1)

		var t test_struct
		err := decoder.Decode(&t)

		if err != nil {
			panic(err)
		}
		log.Println("json quergoogleOAuth2 ---", t.Query)
		///////////////
		ctx := r.Context()
		token := r.Header.Get("Authorization")

		// parseRequest(w, r)

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

const (
	sessionName    = "example-google-app"
	sessionSecret  = "example cookie signing secret"
	sessionUserKey = "googleID"
)

type Config struct {
	ClientID     string
	ClientSecret string
}

var sessionStore = sessions.NewCookieStore([]byte(sessionSecret), nil)

// issueSession issues a cookie session after successful Google login
func issueSession() http.Handler {
	fn := func(w http.ResponseWriter, req *http.Request) {
		ctx := req.Context()
		googleUser, err := google.UserFromContext(ctx)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// 2. Implement a success handler to issue some form of session
		session, _ := sessionStore.New(req, sessionName)
		session.Values[sessionUserKey] = googleUser.Id
		session.Save(req, w)
		http.Redirect(w, req, "/profile", http.StatusFound)
	}
	return http.HandlerFunc(fn)
}

//////// MAIN ////////
func main() {

	// Create a handler for /graphql which passes cors for remote requests
	http.Handle("/graphql", cors.Default().Handler(authMiddleware(&relay.Handler{Schema: gqlSchema.GraphqlSchema})))

	http.HandleFunc("/google/callback", callbackHandler)

	// Write a GraphiQL page to /
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)

	port := viper.GetInt("lychee.server.port")
	goPort := ":" + strconv.Itoa(port) // Needs ":1234" as port
	// ListenAndServe starts an HTTP server with a given address and handler.
	log.Fatal(http.ListenAndServe(goPort, nil))
}

var (
	conf *oauth2.Config
)

func callbackHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	queryParts, _ := url.ParseQuery(r.URL.RawQuery)
	// Use the authorization code that is pushed to the redirect
	// URL.
	code := queryParts["code"][0]
	log.Printf("code: %s\n", code)

	oconf := &oauth2.Config{
		ClientID:     viper.GetString("lychee.oauth.google_client_id"),
		ClientSecret: viper.GetString("lychee.oauth.google_client_secret"),
		RedirectURL:  "http://localhost:8080/google/callback",
		Scopes:       []string{"profile", "email"},
		Endpoint:     googleOAuth2.Endpoint,
	}

	// Exchange will do the handshake to retrieve the initial access token.
	tok, err := oconf.Exchange(ctx, code)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("--------------------------------------------------")
	log.Printf("Token: %s", tok)
	// The HTTP Client returned by conf.Client will refresh the token as necessary.
	client := conf.Client(ctx, tok)

	resp, err := client.Get("http://localhost:8080/")
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println(color.CyanString("Authentication successful"))
	}
	defer resp.Body.Close()

	// show succes page
	msg := "<p><strong>Success!</strong></p>"
	msg = msg + "<p>You are authenticated and can now return to the CLI.</p>"
	fmt.Fprintf(w, msg)
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
