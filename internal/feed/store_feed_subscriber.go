package feed

import (
	"context"
	"fmt"

	"cloud.google.com/go/firestore"
	"github.com/suzuito/village-go/pkg/entity"
	"google.golang.org/api/iterator"
)

type StoreFeedSubscriber struct {
	FirestoreClient *firestore.Client
}

func (t *StoreFeedSubscriber) GetSubscribers(
	ctx context.Context,
	settingID entity.FeedSettingID,
	subscribers *[]*entity.FeedSubscriber,
) error {
	col := t.FirestoreClient.
		Collection(colNameFeedSubscribers)
	iter := col.Where("FeedSettingID", "==", settingID).Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return fmt.Errorf("Next is failed : %w", err)
		}
		subscriber := entity.FeedSubscriber{}
		if err := doc.DataTo(&subscriber); err != nil {
			return fmt.Errorf("DataTo is failed : %w", err)
		}
		*subscribers = append(*subscribers, &subscriber)
	}
	return nil
}

func (t *StoreFeedSubscriber) PutSubscriber(
	ctx context.Context,
	subscriber *entity.FeedSubscriber,
) error {
	docRef := t.FirestoreClient.
		Collection(colNameFeedSubscribers).
		Doc(string(subscriber.ID))
	if _, err := docRef.Set(ctx, subscriber); err != nil {
		return fmt.Errorf("Set is failed : %w", err)
	}
	return nil
}
