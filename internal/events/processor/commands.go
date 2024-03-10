package processor

import (
	"fmt"
	"gym-bot/internal/storage"
	"log"
	"strings"
)

const (
	StartCmd = "/start"
	HelpCmd  = "/help"
	TestCmd  = "/test"
)

func (p *Processor) doCmd(text string, chatID int, username string) error {
	state, err := p.storage.CheckState(username)
	if err != nil {
		return fmt.Errorf("can't do command: %w", err)
	}
	if state != storage.Standard {
		return p.sendInvalidInput(chatID)
	}

	text = strings.ToLower(text)

	log.Printf("got new command %s from %s", text, username)

	switch text {
	case HelpCmd:
		return p.sendHelp(chatID)
	case StartCmd:
		if err := p.storage.CreateUser(username); err != nil {
			return err
		}
		return p.sendHello(chatID)
	case TestCmd:
		return p.sendDefaultKeyboard(chatID)
	default:
		return p.sendUnknownCommand(chatID)
	}
}

func (p *Processor) sendHelp(chatID int) error {
	return p.client.SendMessage(chatID, msgHelp)
}

func (p *Processor) sendHello(chatID int) error {
	return p.client.SendMessage(chatID, msgHello)
}

func (p *Processor) sendUnknownCommand(chatID int) error {
	return p.client.SendMessage(chatID, msgUnknownCommands)
}

func (p *Processor) sendInvalidInput(chatID int) error {
	return p.client.SendMessage(chatID, msgInvalidInput)
}
