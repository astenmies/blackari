package resolvers

import (
	"github.com/astenmies/lychee/micro-post/db"
	"github.com/davecgh/go-spew/spew"
)

type GG struct {
	*db.DB
}

func (g *GG) Greet() *GG {
	spew.Dump("Greet")
	spew.Dump(g.Casual("hola"))

	VV := GG{
		DB: g.DB,
	}

	return &VV
}
