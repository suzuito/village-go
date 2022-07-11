package feed

import (
	"fmt"
	"time"

	"github.com/suzuito/village-go/pkg/entity"
)

var (
	colNameFeedSettings      = "FeedSettings"
	colNameFeedSubscribers   = "FeedSubscribers"
	colNameFeedSubscriptions = "FeedSubscriptions"
	colNameFeedHistories     = "FeedHistories"
)

type docFeedHistory struct {
	ID           string
	SubscriberID entity.FeedSubscriberID
	FeedItemID   entity.FeedItemID
	CreatedAt    time.Time
}

type docFeedSubscription struct {
	FeedSettingID    entity.FeedSettingID
	FeedSubscriberID entity.FeedSubscriberID
}

func (t *docFeedSubscription) ID() string {
	return fmt.Sprintf("%s-%s", t.FeedSettingID, t.FeedSubscriberID)
}
