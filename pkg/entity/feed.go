package entity

import "time"

type FeedSettingType string

const (
	FeedSettingTypeGoBlog  = "go-blog"
	FeedSettingTypeTwitter = "twitter"
)

type FeedSettingID string

type FeedSetting struct {
	ID     FeedSettingID
	Type   FeedSettingType
	Active bool

	TwitterUserID string
}

type FeedHistory struct {
	GUID       string
	RSSFeedURL string
	CreatedAt  time.Time
}

type FeedItemID string

type FeedItem struct {
	ID          FeedItemID
	Title       string
	URL         string
	PublishedAt time.Time
	Source      *FeedItemSource
}

type FeedItemSource struct {
	Name     string
	URL      string
	ImageURL string
}

type FeedSubscriberID string

type FeedSubscriberType string

const (
	FeedSubscriberTypeDiscord FeedSubscriberType = "discord"
)

type FeedSubscriber struct {
	ID               FeedSubscriberID
	Type             FeedSubscriberType
	DiscordChannelID string
}
