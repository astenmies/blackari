package test

import (
	"testing"

	"github.com/astenmies/lychee/server/gqlResolver"
	"github.com/astenmies/lychee/server/gqlSchema"
	graphql "github.com/graph-gophers/graphql-go"
)

func TestQuery(t *testing.T) {
	rootResolver := &gqlResolver.Resolver{}
	_, err := graphql.ParseSchema(gqlSchema.GetRootSchema(), rootResolver)
	if err != nil {
		t.Error(err)

	}
}
