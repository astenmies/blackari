package resolvers

import (
	"github.com/astenmies/lychee/micro-post/db"
	"github.com/astenmies/lychee/micro-post/models"
	"github.com/graph-gophers/graphql-go"
	"go.mongodb.org/mongo-driver/bson"
)

type G struct {
	ID    uint32
	Title string `json:"title"`
}

type PostResolver struct {
	DB    *db.Services
	model models.Post
}

func (r *PostResolver) GetPost() (*PostResolver, error) {
	r.DB.Check("hola")

	post, err := r.DB.GetPostById(bson.M{"id": 1})
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
