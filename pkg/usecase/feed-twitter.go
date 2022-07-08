package usecase

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/rs/zerolog/log"
	"github.com/suzuito/village-go/pkg/entity"
)

type FetcherFeedTwitter struct {
	HTTPClient *http.Client
}

func (t *FetcherFeedTwitter) Fetch(
	ctx context.Context,
	now time.Time,
	setting *entity.FeedSetting,
	items *[]*entity.FeedItem,
) error {
	cli := twitter.NewClient(t.HTTPClient)
	tweets, _, err := cli.Timelines.UserTimeline(&twitter.UserTimelineParams{
		ScreenName: "golang",
	})
	if err != nil {
		return fmt.Errorf("UserTimeline is failed : %w", err)
	}
	threshold := now.Add(-time.Hour * 24)
	for _, tweet := range tweets {
		createdAt, err := tweet.CreatedAtTime()
		if err != nil {
			log.Error().Err(err).Msgf("CreatedAtTime is failed")
			continue
		}
		if tweet.InReplyToStatusID != 0 {
			continue
		}
		if createdAt.Before(threshold) {
			continue
		}
		item := entity.FeedItem{
			ID:          entity.FeedItemID(fmt.Sprintf("%s-%d", setting.Type, tweet.ID)),
			Title:       tweet.Text,
			URL:         fmt.Sprintf("https://twitter.com/%s/status/%d", setting.TwitterUserID, tweet.ID),
			PublishedAt: createdAt,
			Source: &entity.FeedItemSource{
				Name:     tweet.User.Name,
				URL:      fmt.Sprintf("https://twitter.com/%s", setting.TwitterUserID),
				ImageURL: tweet.User.ProfileImageURL,
			},
		}
		*items = append(*items, &item)
	}
	return nil
}
