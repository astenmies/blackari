package gqlResolver

import (
	"github.com/astenmies/lychee/server/model"
	"github.com/astenmies/lychee/server/service"
	"github.com/astenmies/lychee/server/utils"
)

type reviewResolver model.ReviewResolver

// Reviews :
// - Resolves all reviews of a post
func (r *Resolver) Reviews(args struct{ PostSlug string }) []*reviewResolver {
	var l []*reviewResolver
	for _, review := range model.Reviews[args.PostSlug] {
		l = append(l, &reviewResolver{review})
	}
	return l
}

// CreateReview :
// - Returns nil if no args are provided
// - Inserts a new review into mongodb
// - Appends the review to reviews
// - Resolves the current review if args are provided
func (r *Resolver) CreateReview(args *struct {
	PostSlug string
	Review   *model.ReviewInput
}) *reviewResolver {

	// If all fields are nil
	// "createReview": null (in graphql)
	if utils.AllNil(args.Review) {
		return nil
	}

	result := service.InsertOne(args)

	// Make a type PostResolver out of result.
	if s := result; s != nil {
		return &reviewResolver{result}
	}

	return nil
}

// func (r *reviewResolver) ID() string {
// 	return r.R.ID
// }

func (r *reviewResolver) Stars() *int32 {
	return r.R.Stars
}

func (r *reviewResolver) Commentary() *string {
	return r.R.Commentary
}
