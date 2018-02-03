package loader

import (
	"context"
	"sync"

	"github.com/nicksrandall/dataloader"

	blapi "../blapi"
	"github.com/tonyghita/graphql-go-example/errors"
)

func LoadArticle(ctx context.Context, url string) (blapi.Article, error) {
	var article blapi.Article

	ldr, err := extract(ctx, articleLoaderKey)
	if err != nil {
		return article, err
	}

	data, err := ldr.Load(ctx, dataloader.StringKey(url))()
	if err != nil {
		return article, err
	}

	article, ok := data.(blapi.Article)
	if !ok {
		return article, errors.WrongType(article, data)
	}

	return article, nil
}

func LoadArticles(ctx context.Context, urls []string) (ArticleResults, error) {
	var results []ArticleResult
	ldr, err := extract(ctx, articleLoaderKey)
	if err != nil {
		return results, err
	}

	data, errs := ldr.LoadMany(ctx, dataloader.NewKeysFromStrings(urls))()
	results = make([]ArticleResult, 0, len(urls))

	for i, d := range data {
		var e error
		if errs != nil {
			e = errs[i]
		}

		article, ok := d.(blapi.Article)
		if !ok && e == nil {
			e = errors.WrongType(article, d)
		}

		results = append(results, ArticleResult{Article: article, Error: e})
	}

	return results, nil
}

// ArticleResult is the (data, error) pair result of loading a specific key.
type ArticleResult struct {
	Article blapi.Article
	Error   error
}

// ArticleResults is a named type, so methods can be attached to []ArticleResult.
type ArticleResults []ArticleResult

// WithoutErrors filters any result pairs with non-nil errors.
func (results ArticleResults) WithoutErrors() []blapi.Article {
	var articles = make([]blapi.Article, 0, len(results))

	for _, r := range results {
		if r.Error != nil {
			continue
		}

		articles = append(articles, r.Article)
	}

	return articles
}

func PrimeArticles(ctx context.Context, page blapi.ArticlePage) error {
	ldr, err := extract(ctx, articleLoaderKey)
	if err != nil {
		return err
	}

	for _, f := range page.Articles {
		ldr.Prime(ctx, dataloader.StringKey(f.URL), f)
	}

	return nil
}

type articleGetter interface {
	Article(ctx context.Context, url string) (blapi.Article, error)
}

// ArticleLoader contains the client required to load article resources.
type articleLoader struct {
	get articleGetter
}

func newArticleLoader(client articleGetter) dataloader.BatchFunc {
	return articleLoader{get: client}.loadBatch
}

func (ldr articleLoader) loadBatch(ctx context.Context, urls dataloader.Keys) []*dataloader.Result {
	var (
		n       = len(urls)
		results = make([]*dataloader.Result, n)
		wg      sync.WaitGroup
	)

	wg.Add(n)

	for i, url := range urls {
		go func(i int, url dataloader.Key) {
			defer wg.Done()

			resp, err := ldr.get.Article(ctx, url.String())
			results[i] = &dataloader.Result{Data: resp, Error: err}
		}(i, url)
	}

	wg.Wait()

	return results
}
