package main

import (
	"discord-tube/discord"
	"discord-tube/server"
	"log"
	"os"
)

func main() {
	botToken, ok := os.LookupEnv("BOT_TOKEN")
	if !ok {
		log.Fatal("Must set Discord token as env variable")
	}
	discord.Token = botToken
	go server.StartServer()
	discord.Run()
}
