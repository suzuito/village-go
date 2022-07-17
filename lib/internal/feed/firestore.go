package feed

import (
	"time"

	"github.com/suzuito/village-go/pkg/entity"
)

var (
	colNameFeedSettings    = "FeedSettings"
	colNameFeedSubscribers = "FeedSubscribers"
	colNameFeedHistories   = "FeedHistories"
)

type docFeedHistory struct {
	ID           string
	SubscriberID entity.FeedSubscriberID
	FeedItemID   entity.FeedItemID
	CreatedAt    time.Time
}
