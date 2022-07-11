package inject

import (
	"context"
	"fmt"
	"net/http"

	"cloud.google.com/go/firestore"
	"github.com/bwmarrin/discordgo"
	"github.com/suzuito/village-go/internal/feed"
	"github.com/suzuito/village-go/pkg/entity"
	"github.com/suzuito/village-go/pkg/setting"
	"github.com/suzuito/village-go/pkg/usecase"
	"golang.org/x/oauth2/clientcredentials"
)

func NewUsecase(ctx context.Context) (*usecase.Usecase, error) {
	u := usecase.Usecase{}
	cliFirestore, err := firestore.NewClient(
		ctx,
		setting.E.GCPProjectID,
	)
	if err != nil {
		return nil, fmt.Errorf("NewClient is failed : %w", err)
	}
	u.StoreFeedHistory = &feed.StoreFeedHistory{
		FirestoreClient: cliFirestore,
	}
	storeFeedSetting := &feed.StoreFeedSetting{
		FirestoreClient: cliFirestore,
	}
	u.StoreFeedSetting = storeFeedSetting
	discordCli, err := discordgo.New("Bot " + setting.E.DiscordBotToken)
	if err != nil {
		return nil, fmt.Errorf("New is failed : %w", err)
	}
	if setting.E.Env == "dev" {
		discordCli.Debug = true
		discordCli.LogLevel = discordgo.LogDebug
	}
	u.DiscordClient = discordCli
	u.FetchersFeed = usecase.FetchersFeed{
		Fetchers: map[entity.FeedSettingType]usecase.FetcherFeed{
			entity.FeedSettingTypeGoBlog: &usecase.FetcherFeedGolangBlogFeed{
				HTTPClient: http.DefaultClient,
			},
			entity.FeedSettingTypeTwitter: &usecase.FetcherFeedTwitter{
				HTTPClient: (&clientcredentials.Config{
					ClientID:     setting.E.TwitterAPIKey,
					ClientSecret: setting.E.TwitterAPISecret,
					TokenURL:     "https://api.twitter.com/oauth2/token",
				}).Client(ctx),
			},
		},
	}
	return &u, nil
}
