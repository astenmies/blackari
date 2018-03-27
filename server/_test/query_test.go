package test

import (
	"testing"

	"github.com/astenmies/blackari/server/gqlResolver"
	"github.com/astenmies/blackari/server/gqlSchema"
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

func TestPost(t *testing.T) {
	rootSchema := graphql.MustParseSchema(gqlSchema.GetRootSchema(), &gqlResolver.Resolver{})
	slug := "second-post"
	title := "Hello second post"
	t.Run("post query", func(t *testing.T) {
		gqltesting.RunTests(t, []*gqltesting.Test{
			{
				Schema: rootSchema,
				Query: `
					{
						post(slug:"` + slug + `") {
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
							"slug": "` + slug + `",
							"title": "` + title + `"
						}
					}
				`,
			},
		})
	})
}
