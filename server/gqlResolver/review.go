package gqlResolver

import "github.com/neelance/graphql-go"

type review struct {
	stars      int32
	commentary *string
}

var reviews = make(map[string][]*review)

func (r *Resolver) Reviews(args struct{ Post string }) []*reviewResolver {
	var l []*reviewResolver
	for _, review := range reviews[args.Post] {
		l = append(l, &reviewResolver{review})
	}
	return l
}

func (r *Resolver) CreateReview(args *struct {
	Post   string
	Review *reviewInput
}) *reviewResolver {
	review := &review{
		stars:      args.Review.Stars,
		commentary: args.Review.Commentary,
	}
	reviews[args.Post] = append(reviews[args.Post], review)
	return &reviewResolver{review}
}

type reviewResolver struct {
	r *review
}

func (r *reviewResolver) Stars() int32 {
	return r.r.stars
}

func (r *reviewResolver) Commentary() *string {
	return r.r.commentary
}

type friendsConnectionResolver struct {
	ids  []graphql.ID
	from int
	to   int
}

type reviewInput struct {
	Stars      int32
	Commentary *string
}
