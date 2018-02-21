package resolver

import (
	graphql "github.com/neelance/graphql-go"
)

type article struct {
	ID    graphql.ID
	Title string
}

type articleResolver struct {
	s *article
}

var articles = []*article{
	{
		ID:    "1",
		Title: "First article",
	},
	{
		ID:    "2",
		Title: "Second article",
	},
	{
		ID:    "3",
		Title: "Third article",
	},
}

// Slices can be created with the built-in make function; this is how you create dynamically-sized arrays.
var articleData = make(map[graphql.ID]*article)

func init() {
	for _, s := range articles {
		articleData[s.ID] = s
	}
}

func (r *Resolver) Article(args struct{ ID graphql.ID }) *articleResolver {
	if s := articleData[args.ID]; s != nil {
		return &articleResolver{s}
	}
	return nil
}

func (r *articleResolver) ID() graphql.ID {
	return r.s.ID
}

func (r *articleResolver) Title() string {
	return r.s.Title
}
