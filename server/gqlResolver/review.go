package gqlResolver

import (
	"log"
	"reflect"

	mongo "github.com/astenmies/blackari/server/mongo"
	"github.com/astenmies/blackari/server/utils"
	"github.com/rs/xid"
)

type review struct {
	ID         string  `json:"id" bson:"_id,omitempty"`
	Stars      *int32  `json:"stars" bson:"stars,omitempty" `
	Commentary *string `json:"commentary" bson:"commentary,omitempty"`
}

var reviews = make(map[string][]*review)

// Reviews :
// - Resolves all reviews of a post
func (r *Resolver) Reviews(args struct{ PostSlug string }) []*reviewResolver {
	var l []*reviewResolver
	for _, review := range reviews[args.PostSlug] {
		l = append(l, &reviewResolver{review})
	}
	return l
}

func (m review) IsEmpty() bool {
	return reflect.DeepEqual(review{}, m)
}

// CreateReview :
// - Returns nil if no args are provided
// - Inserts a new review into mongodb
// - Appends the review to reviews
// - Resolves the current review if args are provided
func (r *Resolver) CreateReview(args *struct {
	PostSlug string
	Review   *reviewInput
}) *reviewResolver {

	// If all fields are nil
	// "createReview": null (in graphql)
	if utils.AllNil(args.Review) {
		return nil
	}

	newReview := &review{
		Stars:      args.Review.Stars,
		Commentary: args.Review.Commentary,
	}

	reviews[args.PostSlug] = append(reviews[args.PostSlug], newReview)

	userID := xid.New()
	newReview.ID = userID.String()
	session, collection := mongo.Get("review")

	defer session.Close()
	err := collection.Insert(newReview)

	if err != nil {
		log.Fatal(err)
	}

	return &reviewResolver{newReview}
}

type reviewResolver struct {
	r *review
}

func (r *reviewResolver) ID() string {
	return r.r.ID
}

func (r *reviewResolver) Stars() *int32 {
	return r.r.Stars
}

func (r *reviewResolver) Commentary() *string {
	return r.r.Commentary
}

type reviewInput struct {
	Stars      *int32
	Commentary *string
}
