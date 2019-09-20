package resolvers

import "github.com/astenmies/lychee/types"

type PP struct {
	DB string
}

func (w PP) Say() string {
	return w.DB
}

type PostDB interface {
	Casual(s interface{ types.Sayer }) string
}

func Greet(g interface{ PostDB }, i types.Sayer) string {
	return g.Casual(i)
}
