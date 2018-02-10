package blapi

import (
	"context"
	"net/url"
)

type Article struct {
	CharacterURLs []string `json:"characters"`   //
	CreatedAt     string   `json:"created"`      // ISO 8601 format.
	DirectorName  string   `json:"director"`     //
	EpisodeID     int64    `json:"episode_id"`   //
	ProducerNames string   `json:"producer"`     // Comma-separated if more than 1.
	ReleaseDate   string   `json:"release_date"` // ISO 8601 format.
	Title         string   `json:"title"`        //
	URL           string   `json:"url"`          //
	VehicleURLs   []string `json:"vehicles"`     //
}

type ArticlePage struct {
	Count    int64     `json:"count"`
	Articles []Article `json:"results"`
}

func (p ArticlePage) URLs() []string {
	urls := make([]string, len(p.Articles))
	for i, f := range p.Articles {
		urls[i] = f.URL
	}
	return urls
}

func (c *Client) Article(ctx context.Context, url string) (Article, error) {
	r, err := c.NewRequest(ctx, url)
	if err != nil {
		return Article{}, err
	}

	var f Article
	if _, err = c.Do(r, &f); err != nil {
		return Article{}, err
	}

	return f, nil
}

func (c *Client) SearchArticles(ctx context.Context, title string) (ArticlePage, error) {
	q := url.Values{"search": {title}}
	r, err := c.NewRequest(ctx, "/articles?"+q.Encode())
	if err != nil {
		return ArticlePage{}, err
	}

	var fp ArticlePage
	if _, err = c.Do(r, &fp); err != nil {
		return ArticlePage{}, err
	}

	return fp, nil
}
