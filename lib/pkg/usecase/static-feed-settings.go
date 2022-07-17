package usecase

import (
	"context"
	"fmt"

	"github.com/suzuito/village-go/pkg/entity"
	"github.com/suzuito/village-go/pkg/setting"
)

func (t *Usecase) InitStaticFeedSettings(ctx context.Context) error {
	feedSettings := []entity.FeedSetting{
		{
			ID:     "go-blog-1",
			Type:   "go-blog",
			Active: true,
		},
		{
			ID:            "twitter-golang",
			Type:          "twitter",
			Active:        true,
			TwitterUserID: "golang",
		},
	}
	for _, feedSetting := range feedSettings {
		if err := t.StoreFeedSetting.PutFeedSetting(ctx, &feedSetting); err != nil {
			return fmt.Errorf("PutFeedSetting is failed : %w", err)
		}
	}
	subscribers := []entity.FeedSubscriber{}
	if setting.E.Env == "godzilla" {
		subscribers = append(subscribers, entity.FeedSubscriber{
			ID:               "go-blog1-discord-news",
			FeedSettingID:    "go-blog-1",
			Type:             "discord",
			DiscordChannelID: "995138145765044234",
		})
		subscribers = append(subscribers, entity.FeedSubscriber{
			ID:               "twitter-golang-discord-news",
			FeedSettingID:    "twitter-golang",
			Type:             "discord",
			DiscordChannelID: "995138145765044234",
		})
	} else {
		subscribers = append(subscribers, entity.FeedSubscriber{
			ID:               "go-blog1-discord-news",
			FeedSettingID:    "go-blog-1",
			Type:             "discord",
			DiscordChannelID: "995482323241947196",
		})
		subscribers = append(subscribers, entity.FeedSubscriber{
			ID:               "twitter-golang-discord-news",
			FeedSettingID:    "twitter-golang",
			Type:             "discord",
			DiscordChannelID: "995482323241947196",
		})
	}
	for _, subscriber := range subscribers {
		if err := t.StoreFeedSubscriber.PutSubscriber(ctx, &subscriber); err != nil {
			return fmt.Errorf("PutSubscriber is failed : %w", err)
		}
	}
	return nil
}
