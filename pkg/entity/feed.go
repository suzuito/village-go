package entity

import (
	"time"
)

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

/*
type FeedSubscriberPublishResult struct {
	Success []*FeedItem
	Fail    []*FeedItem
}

func (t *FeedSubscriber) Publish(ctx context.Context, setting *FeedSetting, items []*FeedItem) (*FeedSubscriberPublishResult, error) {
	switch t.Type {
	case FeedSubscriberTypeDiscord:
		return t.publishDiscord(ctx, setting, items)
	}
	return &FeedSubscriberPublishResult{
		Fail: items,
	}, fmt.Errorf("unsupported subscriber %s", t.Type)
}

func (t *FeedSubscriber) publishDiscord(ctx context.Context, setting *FeedSetting, items []*FeedItem) (*FeedSubscriberPublishResult, error) {
	result := FeedSubscriberPublishResult{}
	var errReturned error
	for _, item := range items {
		msg := discordgo.MessageSend{}
		switch setting.Type {
		case FeedSettingTypeGoBlog, FeedSettingTypeTwitter:
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

type FeedSubscribers map[FeedSubscriberID]*FeedSubscriber
*/
