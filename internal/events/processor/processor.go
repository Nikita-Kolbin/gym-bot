package processor

import (
	"errors"
	"fmt"
	"gym-bot/internal/client"
	"gym-bot/internal/events"
	"gym-bot/internal/storage"
	"strings"
)

type Processor struct {
	client  *client.Client
	storage storage.Storage
}

var (
	ErrUnknownEventType = errors.New("unknown event type")
	ErrUnknownMetaType  = errors.New("unknown meta type")
)

func New(client *client.Client, storage storage.Storage) *Processor {
	return &Processor{
		client:  client,
		storage: storage,
	}
}

func (p *Processor) Process(event events.Event) error {
	switch event.Type {
	case events.Message:
		return p.processMessage(event)
	default:
		return fmt.Errorf("can't process message: %w", ErrUnknownEventType)
	}
}

func (p *Processor) processMessage(event events.Event) error {
	meta, err := messageMeta(event)
	if err != nil {
		return fmt.Errorf("can't process message: %w", err)
	}

	chatID := meta.ChatID
	text := strings.TrimSpace(meta.Text)
	username := meta.Username

	exist, err := p.storage.UserIsExists(username)
	if err != nil {
		return fmt.Errorf("can't process message: %w", err)
	}
	if !exist {
		return p.storage.CreateUser(username)
	}

	if text[0] == '/' {
		if err = p.doCmd(text, chatID, username); err != nil {
			return fmt.Errorf("can't do command: %w", err)
		}
	} else {
		if err = p.handleMsg(text, chatID, username); err != nil {
			return fmt.Errorf("can't handle message: %w", err)
		}
	}

	return nil
}

func messageMeta(event events.Event) (*events.MessageMeta, error) {
	res, ok := event.Meta.(events.MessageMeta)
	if !ok {
		return nil, fmt.Errorf("can't get meta: %w", ErrUnknownMetaType)
	}
	return &res, nil
}

func (p *Processor) sendUserDoesNotExists(chatID int) error {
	return p.client.SendMessage(chatID, msgUserDoesNotExist)
}
