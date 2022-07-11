package feed

import (
	"context"
	"fmt"

	"cloud.google.com/go/firestore"
	"github.com/suzuito/village-go/pkg/entity"
	"google.golang.org/api/iterator"
)

type StoreFeedSetting struct {
	FirestoreClient *firestore.Client
}

func (t *StoreFeedSetting) GetFeedSettings(
	ctx context.Context,
	settings *[]*entity.FeedSetting,
) error {
	col := t.FirestoreClient.
		Collection(colNameFeedSettings)
	iter := col.Where("Active", "==", true).Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return fmt.Errorf("Next is failed : %w", err)
		}
		setting := entity.FeedSetting{}
		if err := doc.DataTo(&setting); err != nil {
			return fmt.Errorf("DataTo is failed : %w", err)
		}
		*settings = append(*settings, &setting)
	}
	return nil
}

func (t *StoreFeedSetting) PutFeedSetting(
	ctx context.Context,
	setting *entity.FeedSetting,
) error {
	doc := t.FirestoreClient.
		Collection(colNameFeedSettings).
		Doc(string(setting.ID))
	if _, err := doc.Set(ctx, setting); err != nil {
		return fmt.Errorf("Set is failed : %w", err)
	}
	return nil
}

func (t *StoreFeedSetting) GetFeedSubscribers(
	ctx context.Context,
	settingID entity.FeedSettingID,
	subscribers *[]*entity.FeedSubscriber,
) error {
	col := t.FirestoreClient.
		Collection(colNameFeedSubscriptions)
	iter := col.Where("FeedSettingID", "==", string(settingID)).Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return fmt.Errorf("Next is failed : %w", err)
		}
		subscription := docFeedSubscription{}
		if err := doc.DataTo(&subscription); err != nil {
			return fmt.Errorf("DataTo is failed : %w", err)
		}
		snp, err := t.FirestoreClient.
			Collection(colNameFeedSubscribers).
			Doc(string(subscription.FeedSubscriberID)).
			Get(ctx)
		if err != nil {
			return fmt.Errorf("Get is failed : %w", err)
		}
		subscriber := entity.FeedSubscriber{}
		if err := snp.DataTo(&subscriber); err != nil {
			return fmt.Errorf("DataTo is failed : %w", err)
		}
		*subscribers = append(*subscribers, &subscriber)
	}
	return nil
}

func (t *StoreFeedSetting) PutFeedSubscriber(
	ctx context.Context,
	subscriber *entity.FeedSubscriber,
) error {
	if _, err := t.FirestoreClient.
		Collection(colNameFeedSubscribers).
		Doc(string(subscriber.ID)).
		Set(ctx, subscriber); err != nil {
		return fmt.Errorf("Set is failed : %w", err)
	}
	return nil
}

func (t *StoreFeedSetting) Subscribe(
	ctx context.Context,
	subscriberID entity.FeedSubscriberID,
	settingID entity.FeedSettingID,
) error {
	_, err := t.FirestoreClient.
		Collection(colNameFeedSubscribers).
		Doc(string(subscriberID)).
		Get(ctx)
	if err != nil {
		return fmt.Errorf("Get is failed : %w", err)
	}
	_, err = t.FirestoreClient.
		Collection(colNameFeedSettings).
		Doc(string(settingID)).
		Get(ctx)
	if err != nil {
		return fmt.Errorf("Get is failed : %w", err)
	}
	d := docFeedSubscription{
		FeedSettingID:    settingID,
		FeedSubscriberID: subscriberID,
	}
	if _, err := t.FirestoreClient.
		Collection(colNameFeedSubscriptions).
		Doc(d.ID()).
		Set(ctx, d); err != nil {
		return fmt.Errorf("Set is failed : %w", err)
	}
	return nil
}

func (t *StoreFeedSetting) Unsubscribe(
	ctx context.Context,
	subscriberID entity.FeedSubscriberID,
	settingID entity.FeedSettingID,
) error {
	d := docFeedSubscription{
		FeedSettingID:    settingID,
		FeedSubscriberID: subscriberID,
	}
	if _, err := t.FirestoreClient.
		Collection(colNameFeedSubscriptions).
		Doc(d.ID()).
		Delete(ctx); err != nil {
		return fmt.Errorf("Delete is failed : %w", err)
	}
	return nil
}
