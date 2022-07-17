package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/getsentry/sentry-go"
	"github.com/rs/zerolog/log"
	"github.com/suzuito/village-go/pkg/entity"
)

func (t *Usecase) InvokeFeeds(
	ctx context.Context,
) error {
	settings := []*entity.FeedSetting{}
	if err := t.StoreFeedSetting.GetFeedSettings(ctx, &settings); err != nil {
		return fmt.Errorf("GetFeedSettings is failed : %w", err)
	}
	for _, setting := range settings {
		if err := t.InvokerFeed.Invoke(ctx, setting); err != nil {
			return fmt.Errorf("Invoke is failed : %w", err)
		}
		// now := time.Now()
		// if err := t.PublishFeed(ctx, now, setting); err != nil {
		// 	sentry.CaptureException(err)
		// 	log.Error().Err(err).Msgf("PublishFeed is failed")
		// 	continue
		// }
		log.Info().Str("FeedSettingID", string(setting.ID)).Msgf("InvokePublishFeed")
	}
	return nil
}

func (t *Usecase) PublishFeed(
	ctx context.Context,
	now time.Time,
	setting *entity.FeedSetting,
) error {
	items := []*entity.FeedItem{}
	if err := t.FetchersFeed.Fetch(ctx, now, setting, &items); err != nil {
		return fmt.Errorf("Fetch is failed : %w", err)
	}
	log.Info().Str("FeedSettingID", string(setting.ID)).Int("Feeds", len(items)).Msgf("FetchFeed")
	subscribers := []*entity.FeedSubscriber{}
	if err := t.StoreFeedSetting.GetFeedSubscribers(ctx, setting.ID, &subscribers); err != nil {
		return fmt.Errorf("GetSubscribers is failed : %w", err)
	}
	log.Info().Str("FeedSettingID", string(setting.ID)).Int("Feeds", len(items)).Int("FeedSubscribers", len(subscribers)).Msgf("GetSubscribers")
	for _, subscriber := range subscribers {
		filteredItems := []*entity.FeedItem{}
		if err := t.StoreFeedHistory.FilterAlreadySent(ctx, subscriber.ID, items, &filteredItems); err != nil {
			if err != nil {
				return fmt.Errorf("FilterAlreadySent is failed : %w", err)
			}
		}
		subscriberPublisher, exists := t.AvailableFeedSubscriberPublishers[subscriber.ID]
		if !exists {
			err := fmt.Errorf("unsupported subscriber %s", subscriber.ID)
			log.Error().Err(err).Send()
			sentry.CaptureException(err)
			continue
		}
		result, err := subscriberPublisher.Publish(ctx, setting, items)
		if err != nil {
			log.Error().Err(err).Msgf("PublishEach is failed")
			sentry.CaptureException(err)
			continue
		}
		log.Info().Str("FeedSettingID", string(setting.ID)).Str("FeedSubscriberID", string(subscriber.ID)).Int("Feeds", len(filteredItems)).Int("Success", len(result.Success)).Int("Fail", len(result.Fail)).Msgf("Publish")
		if len(result.Success) <= 0 {
			return nil
		}
		if err := t.StoreFeedHistory.Put(ctx, subscriber.ID, result.Success); err != nil {
			if err != nil {
				return fmt.Errorf("Put is failed : %w", err)
			}
		}
	}
	return nil
}

type FeedSubscriberPublisherResult struct {
	Success []*entity.FeedItem
	Fail    []*entity.FeedItem
}

type FeedSubscriberPublisher interface {
	FeedSubscriberID() entity.FeedSubscriberID
	Publish(
		ctx context.Context,
		setting *entity.FeedSetting,
		items []*entity.FeedItem,
	) (*FeedSubscriberPublisherResult, error)
}

type FeedSubscriberPublishers map[entity.FeedSubscriberID]FeedSubscriberPublisher

/*
type PublishFeedEachResult struct {
	Success []*entity.FeedItem
	Fail    []*entity.FeedItem
}

func (t *Usecase) PublishFeedEach(
	ctx context.Context,
	setting *entity.FeedSetting,
	subscriber *entity.FeedSubscriber,
	items []*entity.FeedItem,
) (PublishFeedEachResult, error) {
	switch subscriber.Type {
	case entity.FeedSubscriberTypeDiscord:
		return t.PublishEachDiscord(ctx, setting, subscriber, items)
	}
	return PublishFeedEachResult{
		Fail: items,
	}, fmt.Errorf("unsupported subscriber %s", subscriber.Type)
}

func (t *Usecase) PublishEachDiscord(
	ctx context.Context,
	setting *entity.FeedSetting,
	subscriber *entity.FeedSubscriber,
	items []*entity.FeedItem,
) (PublishFeedEachResult, error) {
	result := PublishFeedEachResult{}
	var errReturned error
	for _, item := range items {
		msg := discordgo.MessageSend{}
		switch setting.Type {
		case entity.FeedSettingTypeGoBlog, entity.FeedSettingTypeTwitter:
			msg.Content = item.URL
		default:
			var author *discordgo.MessageEmbedAuthor
			if item.Source != nil {
				author = &discordgo.MessageEmbedAuthor{
					Name:    item.Source.Name,
					URL:     item.Source.URL,
					IconURL: item.Source.ImageURL,
				}
			}
			msg.Embeds = []*discordgo.MessageEmbed{
				{
					URL:       item.URL,
					Title:     item.Title,
					Timestamp: item.PublishedAt.Format("2006-01-02"),
					Author:    author,
				},
			}
		}
		if _, err := t.DiscordClient.ChannelMessageSendComplex(subscriber.DiscordChannelID, &msg); err != nil {
			result.Fail = append(result.Fail, item)
			errReturned = err
			continue
		}
		result.Success = append(result.Success, item)
	}
	return result, errReturned
}
*/
