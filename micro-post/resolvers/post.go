package resolvers

import (
	"context"

	"github.com/astenmies/lychee/micro-post/db"
	"github.com/astenmies/lychee/micro-post/models"
	"github.com/graph-gophers/graphql-go"
	"go.mongodb.org/mongo-driver/bson"
)

type PostResolver struct {
	DB    *db.Services
	model models.Post
}

func (r *PostResolver) GetPost(ctx context.Context, args struct{ ID *int32 }) (*PostResolver, error) {
	r.DB.Check("hola")
	// id, _ := helpers.IDToUint(args.ID)
	post, err := r.DB.GetPostById(bson.M{"id": args.ID})
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
func (r *PostResolver) Title() *string {
	return &r.model.Title
}

// ID resolves the ID field for Post
func (r *PostResolver) ID() graphql.ID {
	return graphql.ID(r.model.ID)
}
