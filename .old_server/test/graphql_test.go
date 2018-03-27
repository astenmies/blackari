package test

import (
	"testing"

	"../resolver"
	"../schema"
	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/gqltesting"
)

var graphqlSchema = graphql.MustParseSchema(schema.GetRootSchema(), &resolver.Resolver{})

func init() {
}

func TestHelloWorld(t *testing.T) {
	gqltesting.RunTests(t, []*gqltesting.Test{
		{
			Schema: graphqlSchema,
			Query: `
				{
					hello
				}
			`,
			ExpectedResult: `
				{
					"hello": "world"
				}
			`,
		},
	})
}

func TestFragments(t *testing.T) {
	gqltesting.RunTests(t, []*gqltesting.Test{
		{
			Schema: graphqlSchema,
			Query: `
				{
					one: article(id: 1) {
					...whatever
					}
					two: article(id: 2) {
					...whatever
					}
				}
				
				fragment whatever on Article {
					title
				}
			`,
			ExpectedResult: `
				{
					"one": {
						"title": "First article"
					},
					"two": {
						"title": "Second article"
					}
				}
			`,
		},
	})
}
