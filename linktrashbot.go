package main

import (
	"flag"
	"log"
)

func main() {
	t := mustToken()

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
