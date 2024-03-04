package main

import (
	"flag"
	"gym-bot/internal/storage/sqlite"
	"log"
	"os"

	tgClient "gym-bot/internal/client"
	"gym-bot/internal/consumer/eventConsumer"
	eventFetcher "gym-bot/internal/events/fetcher"
	eventProcessor "gym-bot/internal/events/processor"
)

const (
	sqliteStoragePath = "data/sqlite/storage.db"
	batchSize         = 100
)

func main() {
	// В процессе разработки используется переменная
	// окружения, в релизе будет флаг консоли
	client := tgClient.New(os.Getenv("TG_BOT_TOKEN"))

	fetcher := eventFetcher.New(client)

	storage, err := sqlite.New(sqliteStoragePath)
	if err != nil {
		log.Fatal("can't open storage", err)
	}

	processor := eventProcessor.New(client, storage)

	consumer := eventConsumer.New(fetcher, processor, batchSize)

	if err = consumer.Start(); err != nil {
		log.Fatal("service is stopped", err)
	}

	//s, err := sqlite.New(sqliteStoragePath)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//err = s.Init()
	//if err != nil {
	//	log.Fatal(err)
	//}
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
