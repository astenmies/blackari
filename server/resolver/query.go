package resolver

import (
	"context"
	"fmt"

	blapi "../blapi"
	"github.com/tonyghita/graphql-go-example/errors"
)

// The QueryResolver is the entry point for all top-level read operations.
type QueryResolver struct {
	client *blapi.Client
}

func NewRoot(client *blapi.Client) (*QueryResolver, error) {
	fmt.Println("RESOLVER QUERY NewRoot ----", &QueryResolver{client: client})
	if client == nil {
		return nil, errors.UnableToResolve
	}

	return &QueryResolver{client: client}, nil
}

// ArticlesQueryArgs are the arguments for the "articles" query.
type ArticlesQueryArgs struct {
	// Title of the article. When nil, all articles are fetched.
	Title *string
}

// Articles resolves a list of articles. If no arguments are provided, all articles are fetched.
func (r QueryResolver) Articles(ctx context.Context, args ArticlesQueryArgs) (*[]*ArticleResolver, error) {
	page, err := r.client.SearchArticles(ctx, strValue(args.Title))
	if err != nil {
		return nil, err
	}

	return NewArticles(ctx, NewArticlesArgs{Page: page})
}
