package main

import (
	"context"

	graphql "github.com/graph-gophers/graphql-go"
	"github.com/pkg/errors"
	"gopkg.in/mgo.v2/bson"
)

type Post struct {
	ID    uint
	Title string
}

type PostResolver struct {
	db *DB
	m  Post
}

// getPost should authorize the user in ctx and return a pet or error
func (db *DB) getPost(ctx context.Context, id uint) (*Post, error) {
	result := Post{}
	err := db.DB.C("post").Find(bson.M{"id": id}).One(&result)

	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (r *Resolver) GetPost(ctx context.Context, args struct{ ID graphql.ID }) (*PostResolver, error) {
	id, err := gqlIDToUint(args.ID)
	if err != nil {
		return nil, errors.Wrap(err, "GetPet")
	}

	post, err := r.db.getPost(ctx, id)
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
