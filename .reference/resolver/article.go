package resolver

import (
	"context"
	"strings"
	"time"

	blapi "../blapi"
	loader "../loader"
	graphql "github.com/neelance/graphql-go"
	"github.com/tonyghita/graphql-go-example/errors"
)

// ArticleResolver resolves the Article type.
type ArticleResolver struct {
	article blapi.Article
}

type NewArticlesArgs struct {
	Page blapi.ArticlePage
	URLs []string
}

type NewArticleArgs struct {
	Article blapi.Article
	URL     string
}

func NewArticle(ctx context.Context, args NewArticleArgs) (*ArticleResolver, error) {
	var article blapi.Article
	var err error

	switch {
	case args.Article.URL != "":
		article = args.Article
	case args.URL != "":
		article, err = loader.LoadArticle(ctx, args.URL)
	default:
		err = errors.UnableToResolve
	}

	if err != nil {
		return nil, err
	}

	return &ArticleResolver{article: article}, nil
}

func NewArticles(ctx context.Context, args NewArticlesArgs) (*[]*ArticleResolver, error) {
	err := loader.PrimeArticles(ctx, args.Page)
	if err != nil {
		return nil, err
	}

	results, err := loader.LoadArticles(ctx, append(args.URLs, args.Page.URLs()...))
	if err != nil {
		return nil, err
	}

	var (
		articles  = results.WithoutErrors()
		resolvers = make([]*ArticleResolver, 0, len(articles))
		errs      errors.Errors
	)

	for i, article := range articles {
		resolver, err := NewArticle(ctx, NewArticleArgs{Article: article})
		if err != nil {
			errs = append(errs, errors.WithIndex(err, i))
		}

		resolvers = append(resolvers, resolver)
	}

	return &resolvers, errs.Err()
}

// ID resolves the article's unique identifier.
func (r *ArticleResolver) ID() graphql.ID {
	return extractID(r.article.URL)
}

// Episode resolves the episode number of this article.
func (r *ArticleResolver) Episode() int32 {
	return int32(r.article.EpisodeID)
}

// DirectorName resolves the name this article's director.
func (r *ArticleResolver) DirectorName() string {
	return r.article.DirectorName
}

// ProducerNames resolves a list of names of this article's producers.
func (r *ArticleResolver) ProducerNames() []string {
	return strings.Split(r.article.ProducerNames, ", ")
}

// ReleaseDate resolves the time of the article release in the original creator country.
func (r *ArticleResolver) ReleaseDate() (graphql.Time, error) {
	t, err := time.Parse("2006-01-02", r.article.ReleaseDate)
	return graphql.Time{Time: t}, err
}

// CreatedAt resolves the RFC3339 date format of the time this resource was created.
func (r *ArticleResolver) CreatedAt(ctx context.Context) (graphql.Time, error) {
	t, err := time.Parse(time.RFC3339, r.article.CreatedAt)
	return graphql.Time{Time: t}, err
}
