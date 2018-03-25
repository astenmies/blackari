package test

import (
	"context"
	"testing"

	"github.com/astenmies/blackari/server/gqlResolver"
	"github.com/astenmies/blackari/server/gqlSchema"
	"github.com/astenmies/blackari/server/mongo"
	graphql "github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/gqltesting"
)

func TestQuery(t *testing.T) {
	rootResolver := &gqlResolver.Resolver{}
	_, err := graphql.ParseSchema(gqlSchema.GetRootSchema(), rootResolver)
	if err != nil {
		t.Error(err)
	}
}

var rootSchema = graphql.MustParseSchema(gqlSchema.GetRootSchema(), &gqlResolver.Resolver{})

func TestPost(t *testing.T) {
	ctx := context.Background()
	// Call mongo Get, session and reference to the post collection
	session, collection := mongo.Get("post")
	// Close the session so its resources may be put back in the pool or collected, depending on the case.
	defer session.Close()
	con := context.WithValue(ctx, "collection", collection)
	gqltesting.RunTests(t, []*gqltesting.Test{
		{
			Context: con,
			Schema:  rootSchema,
			Query: `
				{
					post(slug:"second-post") {
						id
						slug
						title
					}
				}
			`,
			ExpectedResult: `
				{
					"post": {
						"id": "",
						"slug": "second-post",
						"title": "Hello second post"
					}
				}
			`,
		},
	})
}
