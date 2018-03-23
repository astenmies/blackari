package gqlResolver

import (
	"fmt"

	mongo "github.com/astenmies/blackari/server/mongo"
	"github.com/neelance/graphql-go"
	"gopkg.in/mgo.v2/bson"
)

// In order to respond to queries, a schema needs to have resolve functions for all fields.
// Go’s structs are typed collections of fields. They’re useful for grouping data together to form records.

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

	// If nothing was returned yet, return nil
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
