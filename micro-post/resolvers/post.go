package resolvers

import (
	"github.com/astenmies/lychee/types"
)

type DB types.DB

type PostResolver struct {
	title string `json:"title"`
}

func (r *PostResolver) GetPost() (*PostResolver, error) {
	s := PostResolver{
		title: "Post 1",
	}

	return &s, nil

}

// Title resolves the title field for Post
func (p *PostResolver) Title() *string {
	return &p.title
}
