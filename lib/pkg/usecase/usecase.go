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

// InvokerFeed フィード取得処理を起動する。
// 起動されるフィード取得処理が非同期であるか、同期であるかは特に言及しない。
// フィード取得処理を同期的に実行する場合、フィード取得処理の実行が失敗した場合にはエラーを返すこと。
// フィード取得処理を非同期的に実行する場合、フィード取得処理の起動が失敗した場合にはエラーを返すこと。
type InvokerFeed interface {
	Invoke(ctx context.Context, setting *entity.FeedSetting) error
}

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

type StoreFeedSubscriber interface {
	GetSubscribers(
		ctx context.Context,
		settingID entity.FeedSettingID,
		subscribers *[]*entity.FeedSubscriber,
	) error
	PutSubscriber(
		ctx context.Context,
		subscriber *entity.FeedSubscriber,
	) error
}

type DiscordClient interface {
	ChannelMessageSendComplex(channelID string, data *discordgo.MessageSend) (st *discordgo.Message, err error)
}

type Usecase struct {
	StoreFeedHistory    StoreFeedHistory
	StoreFeedSubscriber StoreFeedSubscriber
	StoreFeedSetting    StoreFeedSetting
	FetchersFeed        FetchersFeed
	DiscordClient       DiscordClient
}
