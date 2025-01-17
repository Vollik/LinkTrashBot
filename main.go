package main

import (
	"flag"
	"log"

	"LinkStrashBot/clients/telegram"
)

const (
	tgBotHost = "api.telegram.org"
)

func main() {
	tgClient := telegram.New(tgBotHost, mustToken())

}

func mustToken() string {
	token := flag.String(
		"token-bot-token",
		"",
		"token for access to tel bot",
	)
	flag.Parse()
	if *token == "" {
		log.Fatal("no token")
	}
	return *token
}
