package fetcher

import (
	"fmt"
	"gym-bot/internal/client"
	"gym-bot/internal/events"
)

type Fetcher struct {
	client *client.Client
	offset int
}

func New(client *client.Client) *Fetcher {
	return &Fetcher{
		client: client,
		offset: 0,
	}
}

func (f *Fetcher) Fetch(limit int) ([]events.Event, error) {
	updates, err := f.client.Updates(f.offset, limit)
	if err != nil {
		return nil, fmt.Errorf("can't get events: %w", err)
	}

	res := make([]events.Event, 0, len(updates))

	for _, u := range updates {
		res = append(res, event(u))
	}

	if len(updates) > 0 {
		f.offset = updates[len(updates)-1].ID + 1
	}

	return res, nil
}

func event(update client.Update) events.Event {
	e := events.Event{
		Type: eventType(update),
	}

	switch e.Type {
	case events.Message:
		e.Meta = events.MessageMeta{
			Text:     update.Message.Text,
			ChatID:   update.Message.Chat.ID,
			Username: update.Message.From.Username,
		}
	}

	return e
}

func eventType(update client.Update) events.Type {
	if update.Message == nil {
		return events.Unknown
	}
	return events.Message
}
