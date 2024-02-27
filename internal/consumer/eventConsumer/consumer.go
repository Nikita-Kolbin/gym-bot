package eventConsumer

import (
	"gym-bot/internal/events"
	"log"
	"time"
)

type Consumer struct {
	fetcher   events.Fetcher
	processor events.Processor
	batchSize int
}

func New(fetcher events.Fetcher, processor events.Processor, batchSize int) *Consumer {
	return &Consumer{
		fetcher:   fetcher,
		processor: processor,
		batchSize: batchSize,
	}
}

func (c *Consumer) Start() error {
	for {
		gotEvents, err := c.fetcher.Fetch(c.batchSize)
		if err != nil {
			log.Printf("ERR consumer %s", err)
			continue
		}

		if len(gotEvents) == 0 {
			time.Sleep(100 * time.Millisecond)
			continue
		}

		if err = c.handleEvents(gotEvents); err != nil {
			log.Printf("ERR consumer %s", err)
			continue
		}
	}
}

func (c *Consumer) handleEvents(events []events.Event) error {
	for _, event := range events {
		// TODO: Сделать логирование новых ивентов
		// TODO: Сделать многопоточную обработку

		if err := c.processor.Process(event); err != nil {
			log.Printf("can't handle event: %s", err)
			continue
		}
	}

	return nil
}
