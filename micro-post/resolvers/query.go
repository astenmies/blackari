package resolvers

import "github.com/astenmies/lychee/micro-post/db"

type Query struct {
	DB *db.Services
}
