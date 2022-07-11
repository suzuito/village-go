package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/suzuito/village-go/pkg/entity"
)

var ErrResourceNotFound = fmt.Errorf("resource not found")

// RSSFeed

type FetcherFeed interface {
	Fetch(
		ctx context.Context,
		now time.Time,
		setting *entity.FeedSetting,
		items *[]*entity.FeedItem,
	) error
}

type FetchersFeed struct {
	Fetchers map[entity.FeedSettingType]FetcherFeed
}

func (t *FetchersFeed) Fetch(
	ctx context.Context,
	now time.Time,
	setting *entity.FeedSetting,
	items *[]*entity.FeedItem,
) error {
	v, exists := t.Fetchers[setting.Type]
	if !exists {
		return fmt.Errorf("unsupported setting type %s", setting.Type)
	}
	return v.Fetch(ctx, now, setting, items)
}

type StoreFeedSetting interface {
	GetFeedSettings(
		ctx context.Context,
		settings *[]*entity.FeedSetting,
	) error
	PutFeedSetting(
		ctx context.Context,
		setting *entity.FeedSetting,
	) error
	GetFeedSubscribers(
		ctx context.Context,
		settingID entity.FeedSettingID,
		subscribers *[]*entity.FeedSubscriber,
	) error
	PutFeedSubscriber(
		ctx context.Context,
		subscriber *entity.FeedSubscriber,
	) error

	Subscribe(
		ctx context.Context,
		subscriberID entity.FeedSubscriberID,
		settingID entity.FeedSettingID,
	) error
	Unsubscribe(
		ctx context.Context,
		subscriberID entity.FeedSubscriberID,
		settingID entity.FeedSettingID,
	) error
}

type StoreFeedHistory interface {
	FilterAlreadySent(
		ctx context.Context,
		subscriberID entity.FeedSubscriberID,
		items []*entity.FeedItem,
		filtered *[]*entity.FeedItem,
	) error
	Put(
		ctx context.Context,
		subscriberID entity.FeedSubscriberID,
		items []*entity.FeedItem,
	) error
}

type DiscordClient interface {
	ChannelMessageSendComplex(channelID string, data *discordgo.MessageSend) (st *discordgo.Message, err error)
}

type Usecase struct {
	StoreFeedHistory StoreFeedHistory
	StoreFeedSetting StoreFeedSetting
	FetchersFeed     FetchersFeed
	DiscordClient    DiscordClient
}
