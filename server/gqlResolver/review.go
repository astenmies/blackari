package gqlResolver

import (
	"log"
	"reflect"

	mongo "github.com/astenmies/blackari/server/mongo"
	"github.com/rs/xid"
)

type review struct {
	ID         string  `json:"id" bson:"_id,omitempty"`
	Stars      *int32  `json:"stars" bson:"stars,omitempty" `
	Commentary *string `json:"commentary" bson:"commentary,omitempty"`
}

var reviews = make(map[string][]*review)

func (r *Resolver) Reviews(args struct{ Post string }) []*reviewResolver {
	var l []*reviewResolver
	for _, review := range reviews[args.Post] {
		l = append(l, &reviewResolver{review})
	}
	return l
}

func (m review) IsEmpty() bool {
	return reflect.DeepEqual(review{}, m)
}

func (r *Resolver) CreateReview(args *struct {
	Post   string
	Review *reviewInput
}) *reviewResolver {
	review := &review{
		Stars:      args.Review.Stars,
		Commentary: args.Review.Commentary,
	}
	if review.IsEmpty() {
		return nil
	}

	// spew.Dump(args)
	reviews[args.Post] = append(reviews[args.Post], review)

	userID := xid.New()
	review.ID = userID.String()
	log.Println("Inserting review in mongo")
	// Call mongo Get, session and reference to the post collection
	session, collection := mongo.Get("review")
	// Close the session so its resources may be put back in the pool or collected, depending on the case.
	defer session.Close()

	// The mock data that we insert.
	err := collection.Insert(review)

	if err != nil {
		log.Fatal(err)
	}
	log.Println("Review inserted successfully!")

	return &reviewResolver{review}
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
