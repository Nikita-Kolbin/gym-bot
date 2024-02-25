package main

import (
	"flag"
	"gym-bot/internal/client"
	"gym-bot/internal/events/fetcher"
	"log"
	"os"
)

func main() {
	// В процессе разработки используется переменная
	// окружения, в релизе будет флаг консоли
	c := client.New(os.Getenv("TG_BOT_TOKEN"))

	// TODO: storage

	f := fetcher.New(c)
	_ = f

	// TODO: processor

	// TODO: consumer
}

func mustToken() string {
	token := flag.String(
		"token",
		"",
		"token for access to telegram bot",
	)

	flag.Parse()

	if *token == "" {
		log.Fatal("missing token variable")
	}

	return *token
}
