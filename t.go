package main

import (
	"context"
	"fmt"
	"time"

	"github.com/nautilus/gateway"
	"github.com/nautilus/graphql"
)

// the schemas of each of our services
const schema1 = `
	type User {
		firstName: String!
		id: ID!
	}

	type Query {
		allUsers: [User!]!
	}
`
const schema2 = `
	type User {
		market: Market
	}

	type Market {
		name: String!
		logo: Image!
		owner: User
	}

	type Image {
		id: ID!
	}
`
const schema3 = `
	type Image {
		thumbnailURL: String!
	}
`

// the query we are going to fire
const query = `
	query {
		allUsers {
			...UserList
		}
	}

	fragment UserList on User {
		firstName
		market {
			...MarketDetails
			owner {
				market {
					owner {
						market {
							owner {
								market {
									owner {
										market {
											owner {
												market {
													owner {
														market {
															owner {
																firstName
															}
														}
													}
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}

	fragment MarketDetails on Market {
		name
		logo {
			thumbnailURL
		}
	}
`

func main() {
	// build up a list of remote services
	remoteServices := []*graphql.RemoteSchema{}
	for i, schemaString := range []string{schema1, schema2, schema3} {
		// parse the schema
		schema, err := graphql.LoadSchema(schemaString)
		if err != nil {
			panic(err)
		}

		// add the schema to the list
		remoteServices = append(remoteServices, &graphql.RemoteSchema{
			Schema: schema,
			URL:    fmt.Sprintf("url-%i", i),
		})
	}

	// in order to avoid network requests to services that don't really exist
	// we're going to use an executor that always returns a fixed value. While this
	// would produce weird results if we were to send it to a user, we don't actually
	// care about the result of the execution, just how long the planning phase took.
	executor := &gateway.MockExecutor{
		map[string]interface{}{
			"hello": "world",
		},
	}

	// instantiate the gateway we'll time with our mock executor
	gw, err := gateway.New(remoteServices, gateway.WithExecutor(executor))
	if err != nil {
		panic(err)
	}

	// start the timer
	start := time.Now()
	// execute the query
	_, err = gw.Execute(&gateway.RequestContext{
		Context: context.Background(),
		Query:   query,
	})
	if err != nil {
		panic(err)
	}

	// just print how long that took
	fmt.Println(time.Since(start))
}
