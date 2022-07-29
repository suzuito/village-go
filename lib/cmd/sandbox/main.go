package main

import (
	"github.com/bwmarrin/discordgo"
)

func main() {
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
