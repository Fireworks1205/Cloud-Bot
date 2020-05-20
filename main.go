package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

var (
	Token string
)

func init() {

	flag.StringVar(&Token, "t", "", "Bot Token")
	flag.Parse()
}

func main() {

	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	dg.AddHandler(messageCreate)

	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running. Press Ctrl + C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	if m.Content == "[도움" {
		s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
			Author:      &discordgo.MessageEmbedAuthor{},
			Description: "도움말",
			Fields: []*discordgo.MessageEmbedField{
				&discordgo.MessageEmbedField{
					Name:   "명령어",
					Value:  "`초대` `정보` `업다운`",
					Inline: true,
				},
				&discordgo.MessageEmbedField{
					Name: "게임",
					Value: "`업다운`",
					Inline: true,
				}
			},
		})
	}

	if m.Content == "[초대" {
		s.ChannelMessageSendEmbed(
			m.ChannelID,
		 	&discordgo.MessageEmbed{
			Author: &discordgo.MessageEmbedAuthor{},
			Description: "초대",
			Fields: []*discordgo.MessageEmbedField{
				&discordgo.MessageEmbedField{
					Name:   "구름봇의 초대링크",
					Value:  "https://bit.ly/36mP8UH",
					Inline: true,
				},
			},
		})
	}

	if m.Content == "[정보" {
		s.ChannelMessageSendEmbed(
			m.ChannelID,
			&discordgo.MessageEmbed{
				Author: &discordgo.MessageEmbedAuthor{},
				Description: "정보",
				Fields: []*discordgo.MessageEmbedField{
					&discordgo.MessageEmbedField{
						Name: "구름봇",
						Value: "이 봇은 <@421620869672992768>가 만들었어요!",
						Inline: true,
					},
				},
			},
		)
	}

	if m.Content == "[업다운" {
		s.ChannelMessageSend(m.ChannelID, "아직 준비중이에요!")
	}

}
