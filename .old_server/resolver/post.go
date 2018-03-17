package resolver

import (
	"../mongo"
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

func (r *Resolver) Post(args struct{ Slug string }) *postResolver {

	oneResult := &post{}

	session, collection := mongo.Get()
	defer session.Close()

	err := collection.Find(bson.M{"slug": args.Slug}).Select(bson.M{}).One(&oneResult)
	if err != nil {
		fmt.Println(err)
	}

	if s := oneResult; s != nil {
		return &postResolver{oneResult}
	}
	return nil
}

func (r *postResolver) ID() graphql.ID {
	return r.s.ID
}

func (r *postResolver) Slug() string {
	return r.s.Slug
}

func (r *postResolver) Title() string {
	return r.s.Title
}
