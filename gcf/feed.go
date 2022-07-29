package gcf

import "context"

func InvokeFeeds(ctx context.Context, _ PubSubMessage) error {
	return errorHandler(u.InvokeFeeds(ctx))
}

func InitStaticFeedSettings(ctx context.Context, _ PubSubMessage) error {
	return errorHandler(u.InitStaticFeedSettings(ctx))
}
