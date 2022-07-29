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
