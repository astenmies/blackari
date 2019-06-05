package main

import (
	"testing"

	gqlResolver "github.com/astenmies/lychee/server/gqlResolver"
	gqlSchema "github.com/astenmies/lychee/server/gqlSchema"
	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/gqltesting"
)

func TestPost(t *testing.T) {
	// connect to the database
	rootSchema := graphql.MustParseSchema(gqlSchema.GetRootSchema(), &gqlResolver.Resolver{})
	slug := "second"
	title := "Hello second post"
	t.Run("post query", func(t *testing.T) {
		gqltesting.RunTests(t, []*gqltesting.Test{
			{
				Schema: rootSchema,
				Query: `
					{
						post(slug:"` + slug + `") {
							slug
							title
						}
					}
				`,
				ExpectedResult: `
					{
						"post": {
							"slug": "` + slug + `",
							"title": "` + title + `"
						}
					}
				`,
			},
		})
	})
}
