package main

import (
	"context"
	"fmt"

	graphql "github.com/graph-gophers/graphql-go"
	"github.com/labstack/gommon/log"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
)

type Post struct {
	ID    uint   `json:"id"`
	Title string `json:"title"`
}

type PostResolver struct {
	db *DB
	m  Post
}

// getPost should authorize the user in ctx and return a post or error
func (db *DB) getPost(ctx context.Context, filter bson.M) (*Post, error) {
	var result Post

	// https://www.mongodb.com/blog/post/mongodb-go-driver-tutorial
	collection := db.Client.Database("lychee").Collection("post")
	err := collection.FindOne(context.TODO(), filter).Decode(&result)

	if err != nil {
		log.Errorf("%s", err)
		// Throw graphql error here!
		return nil, fmt.Errorf("%s", err)
	}

	return &result, nil
}

func (r *Resolver) GetPost(ctx context.Context, args struct{ ID graphql.ID }) (*PostResolver, error) {
	id, err := gqlIDToUint(args.ID)
	if err != nil {
		return nil, errors.Wrap(err, "GetPost")
	}

	post, err := r.db.getPost(ctx, bson.M{"id": id})
	if err != nil {
		return nil, err
	}

	s := PostResolver{
		db: r.db,
		m:  *post,
	}

	return &s, nil

}

// ID resolves the ID field for Post
func (p *PostResolver) ID(ctx context.Context) graphql.ID {
	return graphql.ID(p.m.ID)
}

// Title resolves the title field for Post
func (p *PostResolver) Title(ctx context.Context) *string {
	return &p.m.Title
}
