package processor

import (
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
		return err
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
		return p.sendHello(chatID)
	case TestCmd:
		return p.sendDefaultKeyboard(chatID, msgHello)
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
