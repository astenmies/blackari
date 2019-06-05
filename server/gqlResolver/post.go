package gqlResolver

import (
	"github.com/astenmies/lychee/server/model"
	"github.com/astenmies/lychee/server/service"
)

type postResolver model.PostResolver

// Post :
// - Calls the service
// - Returns the result as a PostResolver
func (r *Resolver) Post(args struct{ Slug string }) *postResolver {

	result := service.PostFindBySlug(args.Slug)

	// Make a type PostResolver out of result.
	if s := result; s != nil {
		return &postResolver{result}
	}

	return nil
}

// A schema needs to have resolve functions for each field.

// ID :
// func (r *postResolver) ID() graphql.ID {
// 	return graphql.ID(r.S.ID)
// }

// Slug :
func (r *postResolver) Slug() string {
	return r.S.Slug
}

// Title :
func (r *postResolver) Title() string {
	return r.S.Title
}
