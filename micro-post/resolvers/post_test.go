package resolvers

import (
	"context"
	"testing"

	"github.com/astenmies/lychee/micro-post/schema"
	gqlResolver "github.com/astenmies/lychee/server/gqlResolver"
	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/gqltesting"
)

var (
	rootSchema = graphql.MustParseSchema(schema.GetSchema())
	ctx        = context.Context
)

func TestPost(t *testing.T) {
	// connect to the database
	rootSchema := graphql.MustParseSchema(schema.GetSchema(), &gqlResolver.Resolver{})
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
