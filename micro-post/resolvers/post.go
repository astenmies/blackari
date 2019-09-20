package resolvers

import (
	"context"

	"github.com/davecgh/go-spew/spew"
	"github.com/graph-gophers/graphql-go"

	"github.com/astenmies/lychee/types"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
)

type DB types.DB

type Post struct {
	ID    uint32
	Title string `json:"title"`
}

type PostResolver struct {
	DB    *DB
	model Post
}

func (db *DB) getPostById(filter bson.M) (*Post, error) {
	var result Post
	collection := db.Client.Database("lychee").Collection("post")
	err := collection.FindOne(context.TODO(), bson.M{"id": uint32(2)}).Decode(&result)

	spew.Dump(result)

	if err != nil {
		// log.Errorf("%s", err)
		// Throw graphql error here!
		return nil, errors.Cause(err)
	}

	return &result, nil
}

func (r *PostResolver) GetPost() (*PostResolver, error) {
	post, err := r.DB.getPostById(bson.M{"id": 1})
	if err != nil {
		return nil, err
	}

	s := PostResolver{
		DB:    r.DB,
		model: *post,
	}

	return &s, nil

}

// Title resolves the title field for Post
func (p *PostResolver) Title() *string {
	return &p.model.Title
}

// ID resolves the ID field for Post
func (p *PostResolver) ID() graphql.ID {
	return graphql.ID(p.model.ID)
}

/////

// type DB struct {
// 	*mongo.Client
// }

// type Resolver struct {
// 	db *DB
// }

// type Post struct {
// 	ID    uint   `json:"id"`
// 	Title string `json:"title"`
// }

// type PostResolver struct {
// 	db *DB
// 	m  Post
// }

// // getPost should authorize the user in ctx and return a post or error
// func (db *DB) getPost(ctx context.Context, filter bson.M) (*Post, error) {
// 	var result Post

// 	// https://www.mongodb.com/blog/post/mongodb-go-driver-tutorial
// 	collection := db.Client.Database("lychee").Collection("post")
// 	err := collection.FindOne(context.TODO(), filter).Decode(&result)

// 	if err != nil {
// 		// log.Errorf("%s", err)
// 		// Throw graphql error here!
// 		return nil, errors.Cause(err)
// 	}

// 	return &result, nil
// }

// func (r *Resolver) GetPost(ctx context.Context, args struct{ ID graphql.ID }) (*PostResolver, error) {
// 	id, err := gqlIDToUint(args.ID)
// 	if err != nil {
// 		return nil, errors.Wrap(err, "GetPost")
// 	}

// 	post, err := r.db.getPost(ctx, bson.M{"id": id})
// 	if err != nil {
// 		return nil, err
// 	}

// 	s := PostResolver{
// 		db: r.db,
// 		m:  *post,
// 	}

// 	return &s, nil

// }

// // ID resolves the ID field for Post
// func (p *PostResolver) ID(ctx context.Context) graphql.ID {
// 	return graphql.ID(p.m.ID)
// }

// // Title resolves the title field for Post
// func (p *PostResolver) Title(ctx context.Context) *string {
// 	return &p.m.Title
// }
