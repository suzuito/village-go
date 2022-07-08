package main

import (
	"context"
	"fmt"
	"os"

	"github.com/bwmarrin/discordgo"
	"github.com/suzuito/village-go/pkg/inject"
)

func main() {
	ctx := context.Background()
	// discordCli, err := discordgo.New("Bot " + setting.E.DiscordBotToken)
	// if err != nil {
	// 	fmt.Printf("%+v\n", err)
	// 	os.Exit(1)
	// }
	// discordCli.Debug = true
	// discordCli.LogLevel = discordgo.LogDebug
	// defer discordCli.Close()
	// if err := mainSendMessage002(discordCli); err != nil {
	// 	fmt.Printf("%+v\n", err)
	// 	os.Exit(1)
	// }
	u, err := inject.NewUsecase(ctx)
	if err != nil {
		fmt.Printf("%+v\n", err)
		os.Exit(1)
	}
	if err := u.InvokeFeeds(ctx); err != nil {
		fmt.Printf("%+v\n", err)
		os.Exit(1)
	}
}

func mainSendMessage001(
	discordCli *discordgo.Session,
) error {
	channelID := "994062861427036210"
	_, err := discordCli.ChannelMessageSend(channelID, "Hello world!")
	if err != nil {
		return err
	}
	return nil
}

func mainSendMessage002(
	discordCli *discordgo.Session,
) error {
	channelID := "994062861427036210"
	_, err := discordCli.ChannelMessageSendComplex(channelID, &discordgo.MessageSend{
		Content: "Hoge Fuga",
		Components: []discordgo.MessageComponent{
			&discordgo.ActionsRow{
				Components: []discordgo.MessageComponent{
					discordgo.Button{
						Emoji: discordgo.ComponentEmoji{
							Name: "📜",
						},
						Label: "Documentation",
						Style: discordgo.LinkButton,
						URL:   "https://discord.com/developers/docs/interactions/message-components#buttons",
					},
				},
			},
		},
		// Embeds:  []*discordgo.MessageEmbed{},
		// Components: []discordgo.MessageComponent{
		// 	discordgo.Button{
		// 		Emoji: discordgo.ComponentEmoji{
		// 			Name: "📜",
		// 		},
		// 		Label: "Documentation",
		// 		Style: discordgo.LinkButton,
		// 		URL:   "https://discord.com/developers/docs/interactions/message-components#buttons",
		// 	},
		// },
	})
	if err != nil {
		return err
	}
	return nil
}
