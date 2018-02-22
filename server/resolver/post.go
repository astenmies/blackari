package resolver

import (
	"../mongo"
	"encoding/json"
	"fmt"

	graphql "github.com/neelance/graphql-go"
	"gopkg.in/mgo.v2/bson"
	"log"
)

type post struct {
	ID    graphql.ID
	Slug  string
	Title string
}

type postResolver struct {
	s *post
}

var posts = []*post{
	{
		ID:    "1",
		Slug:  "first-post",
		Title: "First post",
	},
	{
		ID:    "2",
		Slug:  "second-post",
		Title: "Second post",
	},
	{
		ID:    "3",
		Slug:  "third-post",
		Title: "Third post",
	},
}

// Slices can be created with the built-in make function; this is how you create dynamically-sized arrays.
var postData = make(map[graphql.ID]*post)

func init() {
	for _, s := range posts {
		postData[s.ID] = s
	}
}

func (r *Resolver) Post(args struct{ ID graphql.ID }) *postResolver {

	var results []post

	session, collection := mongo.Get()
	defer session.Close()

	err := collection.Find(bson.M{}).All(&results)
	if err != nil {
		log.Fatal(err)
	}

	kaka, _ := json.Marshal(results)
	// This prints the resulting data from mongo while it is based on the post struct
	fmt.Println(string(kaka))
	// return results, nil

	if s := postData[args.ID]; s != nil {
		fmt.Println("postData", &postResolver{s})
		return &postResolver{s}
	}
	return nil
}

func (r *postResolver) ID() graphql.ID {
	return r.s.ID
}

func (r *postResolver) Slug() string {
	fmt.Println("slug", r.s.Slug)
	return r.s.Slug
}

func (r *postResolver) Title() string {
	return r.s.Title
}
