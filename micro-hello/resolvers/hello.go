package resolvers

import (
	"fmt"

	"github.com/astenmies/lychee/micro-post/db"
	"github.com/astenmies/lychee/micro-post/models"
)

type HelloResolver struct{}

// https://github.com/graph-gophers/graphql-go/issues/96#issuecomment-307663742
func (r *HelloResolver) SayHello() string {
	return "Hi there"
}

type PostResolver struct {
	DB    *db.Services
	model models.Post
}

func (r *PostResolver) Hola() string {
	fmt.Println("COUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUCOU")
	return "Hooooola"

}
