package test

import (
	"../resolver"
	"../schema"
	"github.com/neelance/graphql-go"
	"github.com/neelance/graphql-go/gqltesting"
	"testing"
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
