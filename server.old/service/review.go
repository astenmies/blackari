package service

import (
	"log"

	"github.com/astenmies/lychee/server/model"
	"github.com/astenmies/lychee/server/mongo"
)

// PostInsert :
// - Defines a pointer to a Review with args
// - Appends a review to reviews
// - Generates a userID
// - Opens a mgo session
// - Inserts a Review
// - Closes the mgo session
// - Returns the review that was inserted
func PostInsert(args *struct {
	PostSlug string
	Review   *model.ReviewInput
}) *model.Review {

	newReview := &model.Review{
		Stars:      args.Review.Stars,
		Commentary: args.Review.Commentary,
	}

	model.Reviews[args.PostSlug] = append(model.Reviews[args.PostSlug], newReview)

	// userID := xid.New()
	// newReview.ID = userID.String()

	// [TODO]: make services independent of collections
	// Maybe wrap them all together with if statements to avoid code repetitiveness
	session, collection := mongo.Get("review")

	defer session.Close()
	err := collection.Insert(newReview)

	if err != nil {
		log.Fatal(err)
	}
	return newReview
}
