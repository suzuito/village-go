package usecase

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/getsentry/sentry-go"
	"github.com/rs/zerolog/log"
	"github.com/suzuito/village-go/pkg/entity"
)

type FetcherFeedGolangBlogFeed struct {
	HTTPClient *http.Client
}

func (t *FetcherFeedGolangBlogFeed) Fetch(
	ctx context.Context,
	now time.Time,
	setting *entity.FeedSetting,
	items *[]*entity.FeedItem,
) error {
	baseURL := "https://go.dev"
	blogsURL := fmt.Sprintf("%s/blog", baseURL)
	res, err := t.HTTPClient.Get(blogsURL)
	if err != nil {
		return fmt.Errorf("Get is failed : %w", err)
	}
	defer res.Body.Close()
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return fmt.Errorf("NewDocumentFromReader is failed : %w", err)
	}
	doc.Find(".blogtitle").Each(func(i int, s *goquery.Selection) {
		title := s.Find("a").Text()
		if title == "" {
			return
		}
		href, exists := s.Find("a").Attr("href")
		if !exists {
			return
		}
		dateString := s.Find(".date").Text()
		if dateString == "" {
			return
		}
		blogURL := fmt.Sprintf("%s%s", baseURL, href)
		publishedAt, err := time.Parse("2 January 2006", dateString)
		if err != nil {
			sentry.CaptureException(err)
			log.Error().Err(err).Str("dateString", dateString).Msgf("Parse is failed")
			return
		}
		item := entity.FeedItem{
			ID:          entity.FeedItemID(fmt.Sprintf("golang-blog-%s", publishedAt.Format("2006-01-02"))), // FIXME
			Title:       title,
			URL:         blogURL,
			PublishedAt: publishedAt,
		}
		item.Source = &entity.FeedItemSource{
			Name:     "The Go Blog",
			URL:      baseURL,
			ImageURL: "https://go.dev/favicon.ico",
		}

		*items = append(*items, &item)
	})
	return nil
}
