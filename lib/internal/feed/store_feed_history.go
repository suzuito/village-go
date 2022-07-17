package feed

import (
	"context"
	"fmt"

	"cloud.google.com/go/firestore"
	"github.com/suzuito/village-go/pkg/entity"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func newDocFeedHistoryID(
	subscriberID entity.FeedSubscriberID,
	feedItemID entity.FeedItemID,
) string {
	return fmt.Sprintf("%s-%s", subscriberID, feedItemID)
}

func newDocFeedHistory(
	subscriberID entity.FeedSubscriberID,
	feedItemID entity.FeedItemID,
) *docFeedHistory {
	return &docFeedHistory{
		ID:           newDocFeedHistoryID(subscriberID, feedItemID),
		SubscriberID: subscriberID,
		FeedItemID:   feedItemID,
	}
}

type StoreFeedHistory struct {
	FirestoreClient *firestore.Client
}

func (t *StoreFeedHistory) FilterAlreadySent(
	ctx context.Context,
	subscriberID entity.FeedSubscriberID,
	items []*entity.FeedItem,
	filtered *[]*entity.FeedItem,
) error {
	col := t.FirestoreClient.
		Collection(colNameFeedHistories)
	for _, item := range items {
		docID := newDocFeedHistoryID(subscriberID, item.ID)
		doc := col.Doc(docID)
		snp, err := doc.Get(ctx)
		if err != nil {
			if status.Code(err) != codes.NotFound {
				return fmt.Errorf("Get is failed : %w", err)
			}
		}
		if snp.Exists() {
			continue
		}
		*filtered = append(*filtered, item)
	}
	return nil
}

func (t *StoreFeedHistory) Put(
	ctx context.Context,
	subscriberID entity.FeedSubscriberID,
	items []*entity.FeedItem,
) error {
	for _, item := range items {
		d := newDocFeedHistory(subscriberID, item.ID)
		doc := t.FirestoreClient.
			Collection(colNameFeedHistories).
			Doc(d.ID)
		if _, err := doc.Set(ctx, d); err != nil {
			return fmt.Errorf("Set is failed : %w", err)
		}
	}
	return nil
}
