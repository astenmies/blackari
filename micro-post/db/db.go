package db

import (
	"context"
	"fmt"

	"github.com/astenmies/lychee/micro-post/models"
	"github.com/astenmies/lychee/types"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
)

// Services expose db methods
type Services types.DB

// Check just prints a string
func (s *Services) Check(str string) {
	fmt.Printf("db methods are working! %s", str)
}

// GetPostById retrieves a post from db based on id
func (s *Services) GetPostById(filter bson.M) (*models.Post, error) {
	var result models.Post
	collection := s.Client.Database("lychee").Collection("post")
	err := collection.FindOne(context.TODO(), filter).Decode(&result)

	if err != nil {
		// log.Errorf("%s", err)
		// Throw graphql error here!
		return nil, errors.Cause(err)
	}

	return &result, nil
}
