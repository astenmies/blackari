package resolvers

import (
	"github.com/astenmies/lychee/types"
)

type DB types.DB

type PersonResolver struct {
	name string `json:"name"`
}

func (r *PersonResolver) From() (*PersonResolver, error) {
	s := PersonResolver{
		name: "Sebastien",
	}

	return &s, nil

}

// Title resolves the title field for Post
func (p *PersonResolver) Name() *string {
	return &p.name
}
