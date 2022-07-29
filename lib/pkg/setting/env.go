package setting

import (
	"os"

	"github.com/kelseyhightower/envconfig"
	"github.com/rs/zerolog/log"
)

type Env struct {
	Env          string `envconfig:"ENV"`
	GCPProjectID string `envconfig:"GCP_PROJECT_ID"`

	DiscordBotToken  string `envconfig:"DISCORD_BOT_TOKEN"`
	TwitterAPIKey    string `envconfig:"TWITTER_API_KEY"`
	TwitterAPISecret string `envconfig:"TWITTER_API_SECRET"`
}

var E Env

func init() {
	err := envconfig.Process("", &E)
	if err != nil {
		log.Error().Err(err).Send()
		os.Exit(1)
	}
}
