package resolvers

import (
	"context"
	"fmt"

	"github.com/astenmies/lychee/micro-post/db"
	"github.com/astenmies/lychee/micro-post/models"
	"github.com/graph-gophers/graphql-go"
	"go.mongodb.org/mongo-driver/bson"
)

type PostResolver struct {
	DB    *db.Services
	model models.Post
}

func (r *PostResolver) GetPost(ctx context.Context, args struct{ ID *string }) (*PostResolver, error) {
	r.DB.Check("hola")
	id := *args.ID // dereference the pointer
	fmt.Println(id)
	post, err := r.DB.GetPostById(bson.M{"id": id})
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
