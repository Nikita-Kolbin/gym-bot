package main

import (
	"flag"
	"fmt"
	"gym-bot/internal/client"
	"gym-bot/internal/events/fetcher"
	"log"
	"os"
)

func main() {
	// В процессе разработки используется переменная
	// окружения, в релизе будет флаг консоли
	c := client.New(os.Getenv("TG_BOT_TOKEN"))
	_ = c

	// TODO: storage

	f := fetcher.New(c)

	for {
		e, _ := f.Fetch(10)
		if len(e) > 0 {
			fmt.Printf("%+v\n", e)
		}
	}

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
