package resolver

import (
	"../mongo"
	"encoding/json"
	"fmt"

	graphql "github.com/neelance/graphql-go"
	"gopkg.in/mgo.v2/bson"
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

type PostFromDb struct {
	ID    graphql.ID
	Slug  string
	Title string
}

// Slices can be created with the built-in make function; this is how you create dynamically-sized arrays.
var postData = make(map[graphql.ID]*post)

func init() {
	for _, s := range posts {
		postData[s.ID] = s
	}
}

// Post todo clean up this mess
func (r *Resolver) Post(args struct{ ID graphql.ID }) *postResolver {

	var result []post

	session, collection := mongo.Get()
	defer session.Close()

	yep := PostFromDb{}

	// Try this query to get something... I know it's a mess for now
	// {
	// 	post(id: "super-post") {
	// 	title
	// 	slug
	//   }
	// }
	errr := collection.Find(bson.M{"slug": args.ID}).Select(bson.M{"title": 1}).One(&yep)
	if errr != nil {
		fmt.Println(errr)
	}

	fmt.Println("yep", yep)

	err := collection.Find(bson.M{"slug": args.ID}).Select(bson.M{}).All(&result)
	if err != nil {
		fmt.Println(err)
	}

	hello, _ := json.Marshal(result)
	// This prints the resulting data from mongo while it is based on the post struct
	fmt.Println(string(hello))
	// return results, nil

	// bob := string(hello)

	if s := postData[args.ID]; s != nil {
		one, _ := json.Marshal(&postResolver{s})
		fmt.Println("postData", string(one))
		fmt.Println("s ====", s)
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
	fmt.Println("Title", r.s.Title)
	return r.s.Title
}
