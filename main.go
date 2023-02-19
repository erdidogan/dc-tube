package main

import (
	"discord-tube/discord"
	"log"
	"os"
)

func main() {
	botToken, ok := os.LookupEnv("BOT_TOKEN")
	if !ok {
		log.Fatal("Must set Discord token as env variable")
	}
	discord.Token = botToken
	discord.Run()
}
