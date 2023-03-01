package discord

import (
	"discord-tube/video"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
	"os"
	"os/signal"
	"regexp"
	"strings"
)

var (
	Token string
)

const expression = "https?:\\/\\/(www\\.)?[-a-zA-Z0-9@:%._\\+~#=]{2,256}\\.[a-z]{2,4}\\b([-a-zA-Z0-9@:%_\\+.~#?&//=]*)"

func Run() {
	// Create new Discord Session
	discord, err := discordgo.New("Bot " + Token)
	if err != nil {
		log.Fatal(err)
	}

	discord.AddHandler(newMessage)

	discord.Open()
	defer discord.Close()

	fmt.Println("Bot running...")
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

}

func newMessage(discord *discordgo.Session, message *discordgo.MessageCreate) {
	// Ignore bot message
	if message.Author.ID == discord.State.User.ID {
		return
	}

	if strings.Contains(message.Content, "!find") {
		r, _ := regexp.Compile(expression)
		input := r.FindString(message.Content)
		log.Println("Getting video id from message :", message.Content)
		videoId := video.GetId(input)

		if videoId == "" {
			discord.ChannelMessageSend(message.ChannelID, "Can not find video id")
			return
		}
		log.Println("Getting title from video id:", videoId)
		name, _ := video.GetTitle(videoId)

		log.Println("Getting url from video id :", videoId)
		url, err := video.GetAudioUrl(videoId)

		if err != nil {
			discord.ChannelMessageSend(message.ChannelID, "Something went wrong :(")
			return
		}

		if name == "" {
			discord.ChannelMessageSendEmbed(message.ChannelID, generateEmbeddedMessage("Video Id", url))
		} else {
			discord.ChannelMessageSendEmbed(message.ChannelID, generateEmbeddedMessage(name, url))
		}
	}
}

func generateEmbeddedMessage(title string, url string) *discordgo.MessageEmbed {
	return &discordgo.MessageEmbed{
		Title: title,
		URL:   url,
	}
}
