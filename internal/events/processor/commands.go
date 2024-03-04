package processor

import (
	"log"
	"strings"
)

const (
	StartCmd = "/start"
	HelpCmd  = "/help"
)

func (p *Processor) doCmd(text string, chatID int, username string) error {
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
