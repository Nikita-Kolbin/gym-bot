package main

import (
	"flag"
	"gym-bot/internal/consumer/eventConsumer"
	"log"
	"os"

	tgClient "gym-bot/internal/client"
	eventFetcher "gym-bot/internal/events/fetcher"
	eventProcessor "gym-bot/internal/events/processor"
	"gym-bot/internal/storage/storageMock"
)

func main() {
	// В процессе разработки используется переменная
	// окружения, в релизе будет флаг консоли
	client := tgClient.New(os.Getenv("TG_BOT_TOKEN"))

	fetcher := eventFetcher.New(client)

	sMock := storageMock.New()

	processor := eventProcessor.New(client, sMock)

	consumer := eventConsumer.New(fetcher, processor, 100)

	if err := consumer.Start(); err != nil {
		log.Fatal("service is stopped", err)
	}
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
